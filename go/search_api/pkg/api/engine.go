package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ewgra/go-test-task/pkg/api/middleware"
	"github.com/ewgra/go-test-task/pkg/api/products"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

// Config store environment settings to setup api engine.
type Config struct {
	Listen       string `envconfig:"API_LISTEN" required:"true"`
	AllowOrigins string `envconfig:"ALLOW_ORIGINS" default:"*"`
	JwtSecret    string `envconfig:"JWT_SECRET" required:"true"`
	EsTimeout    int    `envconfig:"ES_TIMEOUT" default:"300"`
	EsURL        string `envconfig:"ES_URL" required:"true"`
	EsIndex      string `envconfig:"ES_INDEX" required:"true"`
}

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(cfg *Config) (*gin.Engine, error) {
	r := gin.New()
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.CorrelationIDMiddleware)
	r.Use(middleware.CorsMiddleware(cfg.AllowOrigins))
	addHealthcheckEndpoint(r)

	authMiddleware, err := middleware.AuthMiddleware(cfg.JwtSecret)
	if err != nil {
		return nil, errors.WithMessage(err, "Can't create auth middleware")
	}

	addAuthEndpoints(r, authMiddleware)

	v1 := r.Group("v1")
	v1.Use(authMiddleware.MiddlewareFunc())

	err = addProductsEndpoint(cfg, v1)

	if err != nil {
		return nil, errors.WithMessage(err, "Can't add products endpoint")
	}

	return r, nil
}

func addAuthEndpoints(engine *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	auth := engine.Group("v1")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.MiddlewareFunc(), authMiddleware.RefreshHandler)
}

func addProductsEndpoint(cfg *Config, group *gin.RouterGroup) error {
	esClient, err := elastic.NewClient(
		elastic.SetURL(cfg.EsURL),
	)
	if err != nil {
		return errors.WithMessage(err, "Can't create elasticsearch client")
	}

	productsHandler := products.NewSearchHandler(esClient, cfg.EsTimeout, cfg.EsIndex)

	group.GET("/products", productsHandler)

	return nil
}

func addHealthcheckEndpoint(engine *gin.Engine) {
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
