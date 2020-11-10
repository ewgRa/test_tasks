package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const (
	// CorrelationIDHeader header key, that used to store correlation id in response headers
	CorrelationIDHeader = "X-Correlation-Id"
	// CorrelationIDCtxValue context key, that used to store correlation id in context
	CorrelationIDCtxValue = "correlationId"
)

// CorrelationIDMiddleware adds a correlation id to the response. If id is not provided in request - it is generated.
// Correlation id is used for tracing and debugging,
// more information can be found here https://www.scalyr.com/blog/microservices-logging-best-practices.
// A logger that automatically adds correlation id to log records will be stored to the request context.
// You should obtain logger in handlers from context: "log.Ctx(c.Request.Context())"
func CorrelationIDMiddleware(c *gin.Context) {
	id := c.GetHeader(CorrelationIDHeader)

	if id == "" {
		id = uuid.New().String()
	}

	c.Header(CorrelationIDHeader, id)
	c.Set(CorrelationIDCtxValue, id)

	logger := log.With().Str("correlation_id", id).Logger()
	c.Request = c.Request.WithContext(logger.WithContext(c.Request.Context()))
	c.Next()
}
