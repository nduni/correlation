package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/nduni/correlation/common/configuration"
	"github.com/nduni/correlation/common/messaging/rabbitmq"
	app "github.com/nduni/correlation/weather/weather-acceptor/app"
	"github.com/spf13/viper"
)

var Config configuration.Configuration

func main() {
	readConfig()
	rabbitmq.StartRabbitmqSenders(Config.BrokerConnections)
	runCronJobs()
}

func readConfig() {
	viper.SetConfigName("config_local") // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./app/resources/config")
	viper.ReadInConfig()
	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(err.Error())
	}
}

func runCronJobs() {
	cron := gocron.NewScheduler(time.UTC)
	fmt.Println("Next cron job is starting at", time.Now().Round(1*time.Hour).Add(1*time.Hour).String())
	cron.Every(1).Hours().StartAt(time.Now().UTC().Round(1 * time.Hour)).Do(func() {
		fmt.Println(time.Now().String())
		app.ProcessWeather()
	})
	cron.StartBlocking()
}
