package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

// RecoverMiddleware handle gin panics. Gin instead of return errors usually panic.
// We respond to such cases with "500 Internal server error" response code and log such error.
func RecoverMiddleware(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			log.Ctx(c.Request.Context()).Error().Err(rec.(error)).Msg("Internal server error")
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}(c)

	c.Next()
}
