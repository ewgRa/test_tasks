// Package search provides functionality to handle search product endpoint.
package search

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/product"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// NewHandler creates new Handler instance.
func NewHandler(repository *product.Repository) *Handler {
	requestPool := &sync.Pool{
		New: func() interface{} { return newRequest() },
	}

	return &Handler{
		requestPool: requestPool,
		repository:  repository,
	}
}

// Handler handles product search requests.
type Handler struct {
	repository  *product.Repository
	requestPool *sync.Pool
}

// Handle reads request, perform search in storage and response.
func (h *Handler) Handle(c *gin.Context) {
	request := h.requestPool.Get().(*request)
	defer h.requestPool.Put(request)

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	products, err := h.repository.Search(
		c.Request.Context(),
		request.Query,
		request.Brand,
		request.OrderBy,
		request.Sort,
		request.Offset,
		request.Limit,
	)
	if err != nil {
		log.Ctx(c.Request.Context()).Error().Caller().Err(err).Msg("Search products failed")
		c.Status(http.StatusInternalServerError)

		return
	}

	response := newResponse()
	response.Products = products

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Ctx(c.Request.Context()).Error().Caller().Err(err).Msg("Error json marshal products")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonResponse)
}
