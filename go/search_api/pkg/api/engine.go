// Package api provides gin engine (router) and other common files.
package api

import (
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/middleware"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/product"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/product/search"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(cfg *Config) (*gin.Engine, error) {
	engine := gin.New()

	engine.Use(
		middleware.RecoverMiddleware,
		middleware.CorrelationIDMiddleware,
		middleware.CorsMiddleware(cfg.AllowOrigins),
	)

	authMiddleware, err := middleware.AuthMiddleware(cfg.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("can't create auth middleware: %w", err)
	}

	addAuthEndpoints(engine, authMiddleware)

	err = addProductEndpoints(engine, cfg, authMiddleware.MiddlewareFunc())
	if err != nil {
		return nil, fmt.Errorf("can't add product endpoints: %w", err)
	}

	err = addHealthEndpoints(engine, cfg)
	if err != nil {
		return nil, fmt.Errorf("can't create health endpoints: %w", err)
	}

	return engine, nil
}

func addAuthEndpoints(engine *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	routerGroup := engine.Group("v1")
	routerGroup.POST("/login", authMiddleware.LoginHandler)
	routerGroup.GET("/refresh_token", authMiddleware.MiddlewareFunc(), authMiddleware.RefreshHandler)
}

func addProductEndpoints(engine *gin.Engine, cfg *Config, authHandler gin.HandlerFunc) error {
	routerGroup := engine.Group("v1")
	routerGroup.Use(authHandler)

	esClient, err := elastic.NewClient(
		elastic.SetURL(cfg.EsURL),
	)
	if err != nil {
		return fmt.Errorf("can't create elasticsearch client: %w", err)
	}

	productRepository := product.NewRepository(esClient, cfg.EsTimeout, cfg.EsIndex)
	productHandler := search.NewHandler(productRepository)
	routerGroup.GET("/products", productHandler.Handle)

	return nil
}

func addHealthEndpoints(engine *gin.Engine, cfg *Config) error {
	esClient, err := elastic.NewClient(
		elastic.SetURL(cfg.EsURL),
	)
	if err != nil {
		return fmt.Errorf("can't create elasticsearch client: %w", err)
	}

	health := NewHealth(esClient)
	engine.GET("/health/liveness", health.liveness)
	engine.GET("/health/readiness", health.readiness)

	return nil
}
