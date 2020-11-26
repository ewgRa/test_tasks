package product

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sony/gobreaker"
)

// NewRepository creates new Repository instance.
func NewRepository(esClient *elastic.Client, timeout int, index string) *Repository {
	var circuitBreakerSettings gobreaker.Settings
	circuitBreakerSettings.Name = "Elasticsearch"
	esCircuitBreaker := gobreaker.NewCircuitBreaker(circuitBreakerSettings)

	repository := &Repository{
		client:         esClient,
		timeout:        timeout,
		index:          index,
		circuitBreaker: esCircuitBreaker,
	}

	return repository
}

// Repository responsible for query and store products in storage.
type Repository struct {
	client         *elastic.Client
	timeout        int
	index          string
	circuitBreaker *gobreaker.CircuitBreaker
}

// Search products in storage.
func (r *Repository) Search(
	ctx context.Context,
	query string,
	brand string,
	orderBy string,
	sort string,
	offset int,
	limit int,
) ([]*Product, error) {
	esQuery := r.searchQuery(query, brand)
	source := r.searchSource(esQuery, orderBy, sort, offset, limit)
	service := r.searchService(source)

	res, err := r.circuitBreaker.Execute(func() (interface{}, error) {
		return service.Do(ctx)
	})
	if err != nil {
		return nil, fmt.Errorf("error when search inside circuit breaker: %w", err)
	}

	return r.resultToProducts(res.(*elastic.SearchResult))
}

func (r *Repository) searchSource(
	query elastic.Query,
	orderBy string,
	sort string,
	offset int,
	limit int,
) *elastic.SearchSource {
	searchSource := elastic.NewSearchSource()

	searchSource.
		Query(query).
		Sort(orderBy, sort == "asc").
		From(offset).Size(limit)

	return searchSource
}

func (r *Repository) searchService(source *elastic.SearchSource) *elastic.SearchService {
	return r.client.Search().
		TimeoutInMillis(r.timeout).
		Index(r.index).
		SearchSource(source)
}

func (r *Repository) searchQuery(query string, brand string) elastic.Query {
	boolQuery := elastic.NewBoolQuery()

	if query != "" {
		boolQuery.Must(elastic.NewMatchQuery("title", query))
	}

	if brand != "" {
		boolQuery.Filter(elastic.NewTermQuery("brand", brand))
	}

	return boolQuery
}

func (r *Repository) resultToProducts(res *elastic.SearchResult) ([]*Product, error) {
	productList := make([]*Product, res.TotalHits())[:0]

	if res.TotalHits() == 0 {
		return productList, nil
	}

	for _, hit := range res.Hits.Hits {
		product := NewProduct()

		err := json.Unmarshal(hit.Source, product)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal products json %s: %w", hit.Source, err)
		}

		productList = append(productList, product)
	}

	return productList, nil
}
