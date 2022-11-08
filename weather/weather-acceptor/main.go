package main

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/nduni/correlation/common/logger"
	app "github.com/nduni/correlation/weather/weather-acceptor/app"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger = logger.NewPackageLogger("main")

func main() {
	err := app.LoadConfig()
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	err = app.InitBrokerSenders()
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	runCronJobs()

	select {}
}

func runCronJobs() {
	cron := gocron.NewScheduler(time.Local)
	now := time.Now().Local()
	log.Info().Msgf("Next cron job is starting at %s", now.Round(1*time.Hour).Add(1*time.Hour).String())
	cron.Every(1).Hour().StartAt(now.Round(1 * time.Hour)).Do(func() {
		log.Debug().Msgf("starting new cron job at: %s", time.Now().Local().String())
		app.ProcessWeather(context.Background())
	})
	cron.StartAsync()
}
