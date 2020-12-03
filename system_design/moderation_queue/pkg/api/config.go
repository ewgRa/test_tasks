// Package api provides gin engine (router) and other common files.
package api

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// NewConfig creates new config instance to provide configuration to engine.
func NewConfig() *Config {
	return &Config{}
}

// Config is configuration storage for application.
type Config struct {
	Listen       string `envconfig:"API_LISTEN" default:":8080"`
	AllowOrigins string `envconfig:"ALLOW_ORIGINS" default:"*"`
	RabbitMqURL  string `envconfig:"RABBIT_MQ_URL" default:"amqp://guest:guest@rabbitmq:5672/"`
}

// LoadFromEnv loads Config properties from environment.
func (c *Config) LoadFromEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return fmt.Errorf("can't process environment for config: %w", err)
	}

	return nil
}
