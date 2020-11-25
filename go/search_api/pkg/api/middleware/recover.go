package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// RecoverMiddleware handle gin panics. Gin instead of return errors usually panic.
// We respond to such cases with "500 Internal server error" response code and log such error.
func RecoverMiddleware(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			errEvent := log.Ctx(c.Request.Context()).Error()

			if err, ok := rec.(error); ok {
				errEvent.Err(err).Msg("Internal server error")
			} else {
				errEvent.Msgf("Internal server error: %v", rec)
			}

			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}(c)

	c.Next()
}
