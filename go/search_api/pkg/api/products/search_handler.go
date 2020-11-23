package products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sony/gobreaker"
)

// NewSearchHandler create handler that process products search requests.
func NewSearchHandler(esClient *elastic.Client, timeout int, index string) gin.HandlerFunc {
	var circuitBreakerSettings gobreaker.Settings
	circuitBreakerSettings.Name = "Elasticsearch"
	esCircuitBreaker := gobreaker.NewCircuitBreaker(circuitBreakerSettings)

	requestPool := &sync.Pool{
		New: func() interface{} { return newSearchRequest() },
	}

	responsePool := &sync.Pool{
		New: func() interface{} { return newSearchResponse() },
	}

	handler := &searchHandler{
		requestPool:    requestPool,
		responsePool:   responsePool,
		client:         esClient,
		timeout:        timeout,
		index:          index,
		circuitBreaker: esCircuitBreaker,
	}

	return handler.Handle
}

type searchHandler struct {
	requestPool    *sync.Pool
	responsePool   *sync.Pool
	client         *elastic.Client
	timeout        int
	index          string
	circuitBreaker *gobreaker.CircuitBreaker
}

func (ph *searchHandler) Handle(c *gin.Context) {
	productsRequest := ph.requestPool.Get().(*searchRequest)
	defer ph.requestPool.Put(productsRequest)

	if err := c.ShouldBindQuery(productsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	searchResult, err := ph.doSearch(c.Request.Context(), productsRequest)
	if err != nil {
		var e *elastic.Error

		if errors.As(err, &e) {
			log.Ctx(c.Request.Context()).Error().
				Msgf("Search request to elasticsearch failed with status %d and error %v.", e.Status, e.Details)
		} else {
			log.Ctx(c.Request.Context()).Error().Err(err).Msg("Search request to elasticsearch failed")
		}

		c.Status(http.StatusInternalServerError)

		return
	}

	ph.flushResult(c, searchResult)
}

func (ph *searchHandler) doSearch(ctx context.Context, request *searchRequest) (*elastic.SearchResult, error) {
	source := ph.searchSource(request)
	service := ph.searchService().SearchSource(source)

	ctx.Deadline()

	res, err := ph.circuitBreaker.Execute(func() (interface{}, error) {
		return service.Do(ctx)
	})
	if err != nil {
		return nil, fmt.Errorf("error when execute circuit breaker for search: %w", err)
	}

	return res.(*elastic.SearchResult), nil
}

func (ph *searchHandler) searchSource(request *searchRequest) *elastic.SearchSource {
	boolQuery := elastic.NewBoolQuery()

	if request.Query != "" {
		boolQuery.Must(elastic.NewMatchQuery("title", request.Query))
	}

	if request.Brand != "" {
		boolQuery.Filter(elastic.NewTermQuery("brand", request.Brand))
	}

	searchSource := elastic.NewSearchSource()

	searchSource.
		Query(boolQuery).
		Sort(request.OrderBy, request.Sort == "asc").
		From(request.Offset).Size(request.Limit)

	return searchSource
}

func (ph *searchHandler) flushResult(c *gin.Context, result *elastic.SearchResult) {
	productsJSON, err := ph.productsJSON(result)
	if err != nil {
		log.Ctx(c.Request.Context()).Error().Err(err).Msg("Can't marshal search results")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", productsJSON)
}

func (ph *searchHandler) productsJSON(res *elastic.SearchResult) ([]byte, error) {
	resp := ph.responsePool.Get().(*searchResponse)
	defer ph.responsePool.Put(resp)
	resp.reset()

	if res.TotalHits() > 0 {
		for _, hit := range res.Hits.Hits {
			product := resp.product()

			err := json.Unmarshal(hit.Source, product)
			if err != nil {
				return []byte{}, fmt.Errorf("failed to unmarshal products json %s: %w", hit.Source, err)
			}

			resp.Data = append(resp.Data, product)
		}
	}

	return json.Marshal(resp)
}

func (ph *searchHandler) searchService() *elastic.SearchService {
	return ph.client.Search().
		TimeoutInMillis(ph.timeout).
		Index(ph.index)
}
