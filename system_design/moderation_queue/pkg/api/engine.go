// Package api provides gin engine (router) and other common files.
package api

import (
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/category"
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/message"
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/middleware"
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/rabbitmq"
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api/topic"
	"github.com/gin-gonic/gin"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(cfg *Config) (*gin.Engine, error) {
	engine := gin.New()

	engine.Use(
		middleware.RecoverMiddleware,
		middleware.CorsMiddleware(cfg.AllowOrigins),
	)

	rabbitMq := rabbitmq.NewRabbitMq(cfg.RabbitMqURL)
	topicHandler := topic.NewHandler(rabbitMq)
	engine.POST("/topic", topicHandler.Handle)

	categoryHandler := category.NewHandler(rabbitMq)
	engine.POST("/category", categoryHandler.Handle)

	messageHandler := message.NewHandler(rabbitMq)
	engine.POST("/message", messageHandler.Handle)

	addHealthEndpoints(engine, rabbitMq)

	return engine, nil
}

func addHealthEndpoints(engine *gin.Engine, rabbitMq *rabbitmq.RabbitMq) {
	health := NewHealth(rabbitMq)
	engine.GET("/health/liveness", health.liveness)
	engine.GET("/health/readiness", health.readiness)
}
