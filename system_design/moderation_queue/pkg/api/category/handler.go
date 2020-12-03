// Package category provides functionality to handle category endpoint.
package category

import (
	"net/http"

	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// NewHandler creates new Handler instance.
func NewHandler(rabbitMq *rabbitmq.RabbitMq) *Handler {
	return &Handler{rabbitMq: rabbitMq}
}

// Handler handles category requests.
type Handler struct {
	rabbitMq *rabbitmq.RabbitMq
}

// Handle reads request, create category and response.
func (h *Handler) Handle(c *gin.Context) {
	request := &struct {
		Topic    string `json:"topic" binding:"required"`
		Category string `json:"Category" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"reason": "bad_request",
			"error":  err.Error(),
		})

		return
	}

	err := h.rabbitMq.CreateCategory(request.Topic, request.Category)
	if err != nil {
		log.Ctx(c.Request.Context()).Error().Caller().Err(err).Msg("Failed to create category")
		c.Status(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
