package main

import (
	"fmt"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/strategy"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	cfg, err := createConfig()
	if nil != err {
		log.Fatal().Caller().Err(err).Msg("Fail to create config")

		return
	}

	engine, err := createEngine(cfg)
	if nil != err {
		log.Fatal().Caller().Err(err).Msg("Fail to create engine")

		return
	}

	log.Info().Caller().Msg("Start server")
	log.Fatal().Caller().Err(engine.Run(cfg.Listen)).Msg("Fail to listen and serve")
}

func createEngine(cfg *api.Config) (*gin.Engine, error) {
	var engine *gin.Engine

	retryErr := retry.Retry(
		func(attempt uint) error {
			log.Info().Caller().Uint("attempt", attempt).Msg("Trying to create api engine")

			var err error

			engine, err = api.CreateAPIEngine(cfg)
			if err != nil {
				log.Error().Caller().Uint("attempt", attempt).Err(err).Msg("Fail to create api engine")

				return fmt.Errorf("fail to create api engine at %v attempt: %w", attempt, err)
			}

			return nil
		},
		strategy.Backoff(backoff.Linear(1*time.Second)),
	)

	if retryErr != nil {
		return nil, fmt.Errorf("fail to create engine: %w", retryErr)
	}

	return engine, nil
}

func createConfig() (*api.Config, error) {
	cfg := api.NewConfig()

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("can't process environment variables: %w", err)
	}

	return cfg, nil
}
