package main

import (
	app "github.com/nduni/correlation/weather/weather-processor/app"
)

func main() {
	err := app.LoadConfig()
	if err != nil {
		panic(err)
	}
	err = app.StartSubscription()
	if err != nil {
		panic(err)
	}

	select {}
}
