package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// NewHealth creates new Health instance.
func NewHealth(esClient *elastic.Client) *Health {
	return &Health{esClient: esClient}
}

// Health provides liveness and readiness probe responses.
type Health struct {
	esClient *elastic.Client
}

func (h *Health) liveness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Health) readiness(c *gin.Context) {
	esHealth, err := h.esClient.ClusterHealth().Timeout("3s").Do(c.Request.Context())
	if err != nil || esHealth.Status == "red" {
		c.Status(http.StatusServiceUnavailable)

		return
	}

	c.Status(http.StatusOK)
}
