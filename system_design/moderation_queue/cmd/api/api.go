package main

import (
	"github.com/ewgRa/test_tasks/system_design/moderation_queue/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	gin.DefaultWriter = log.Logger

	cfg := api.NewConfig()

	err := cfg.LoadFromEnv()
	if nil != err {
		log.Fatal().Caller().Err(err).Msg("Fail to create config")

		return
	}

	engine, err := api.CreateAPIEngine(cfg)
	if nil != err {
		log.Fatal().Caller().Err(err).Msg("Fail to create engine")

		return
	}

	log.Info().Caller().Msg("Start server")
	log.Fatal().Caller().Err(engine.Run(cfg.Listen)).Msg("Fail to listen and serve")
}
