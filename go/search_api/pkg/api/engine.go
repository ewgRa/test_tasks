// Package api provide gin engine (router) and other common files.
package api

import (
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/middleware"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/products/search"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// NewConfig create new config instance to provide configuration to engine.
func NewConfig() *Config {
	return &Config{}
}

// Config store environment settings to setup api engine.
type Config struct {
	Listen       string `envconfig:"API_LISTEN" default:":8080"`
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
		return nil, fmt.Errorf("can't create auth middleware: %w", err)
	}

	addAuthEndpoints(r, authMiddleware)

	v1 := r.Group("v1")
	v1.Use(authMiddleware.MiddlewareFunc())

	err = addProductsEndpoint(cfg, v1)

	if err != nil {
		return nil, fmt.Errorf("can't add products endpoint: %w", err)
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
		return fmt.Errorf("can't create elasticsearch client: %w", err)
	}

	productsHandler := search.NewSearchHandler(esClient, cfg.EsTimeout, cfg.EsIndex)

	group.GET("/products", productsHandler)

	return nil
}

func addHealthcheckEndpoint(engine *gin.Engine) {
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
