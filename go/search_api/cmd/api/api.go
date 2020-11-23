package main

import (
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/strategy"
	"github.com/ewgra/go-test-task/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var cfg api.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't process environment variables")

		return
	}

	var engine *gin.Engine

	err = retry.Retry(
		func(attempt uint) error {
			log.Info().Uint("attempt", attempt).Msg("Trying to create a server")
			engine, err = api.CreateAPIEngine(&cfg)

			if err != nil {
				log.Error().Err(err).Msg("Fail to create server")
			}

			return err
		},
		strategy.Backoff(backoff.Linear(1*time.Second)),
	)

	if nil != err {
		log.Fatal().Err(err).Msg("Create server failed")

		return
	}

	log.Info().Msg("Start server")
	log.Fatal().Err(engine.Run(cfg.Listen)).Msg("Fail to listen and serve")
}
