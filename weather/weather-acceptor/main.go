package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	app "github.com/nduni/correlation/weather/weather-acceptor/app"
)

func main() {
	err := app.LoadConfig()
	if err != nil {
		panic(err)
	}
	err = app.InitBrokerSenders()
	if err != nil {
		panic(err)
	}
	runCronJobs()
}

func runCronJobs() {
	cron := gocron.NewScheduler(time.UTC)
	fmt.Println("Next cron job is starting at", time.Now().Round(1*time.Hour).Add(1*time.Hour).String())
	cron.Every(1).Hour().StartAt(time.Now().UTC().Round(1 * time.Hour)).Do(func() {
		fmt.Println(time.Now().String())
		app.ProcessWeather(context.Background())
	})
	cron.StartBlocking()
}
