package api

import (
	"net/http"

	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// NewHealth creates new Health instance.
func NewHealth(rabbitMq *rabbitmq.RabbitMq) *Health {
	return &Health{rabbitMq: rabbitMq}
}

// Health provides liveness and readiness probe responses.
type Health struct {
	rabbitMq *rabbitmq.RabbitMq
}

func (h *Health) liveness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Health) readiness(c *gin.Context) {
	ping, err := h.rabbitMq.Ping()
	if !ping || err != nil {
		log.Ctx(c.Request.Context()).Error().Caller().Err(err).Msg("Ping to RabbitMQ failed")
		c.Status(http.StatusServiceUnavailable)

		return
	}

	c.Status(http.StatusOK)
}
