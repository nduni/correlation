package main

import (
	"github.com/nduni/correlation/common/logger"
	app "github.com/nduni/correlation/weather/weather-processor/app"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger = logger.NewPackageLogger("main")

func main() {
	err := app.LoadConfig()
	if err != nil {
		log.Panic().Msgf(err.Error())
	}
	err = app.StartSubscription()
	if err != nil {
		log.Panic().Msgf(err.Error())
	}

	select {}
}
