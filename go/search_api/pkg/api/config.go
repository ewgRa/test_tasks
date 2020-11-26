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
	JwtSecret    string `envconfig:"JWT_SECRET" required:"true"`
	EsTimeout    int    `envconfig:"ES_TIMEOUT" default:"300"`
	EsURL        string `envconfig:"ES_URL" required:"true"`
	EsIndex      string `envconfig:"ES_INDEX" required:"true"`
}

// LoadFromEnv loads Config properties from environment.
func (c *Config) LoadFromEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return fmt.Errorf("can't process environment for config: %w", err)
	}

	return nil
}
