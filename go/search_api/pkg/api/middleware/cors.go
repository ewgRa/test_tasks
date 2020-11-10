package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware create gin middleware for control CORS headers.
// We add "Authorization" as allowed header for proper SwaggerUI work.
func CorsMiddleware(allowOrigins string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowOrigins = []string{allowOrigins}

	return cors.New(corsConfig)
}
