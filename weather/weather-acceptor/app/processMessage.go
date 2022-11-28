package app

import (
	"context"

	"github.com/nduni/correlation/common/messaging"
	mapper "github.com/nduni/correlation/weather/weather-acceptor/mappers"
)

const weather_api = "https://api.open-meteo.com/v1/forecast"

func ProcessWeather(ctx context.Context) error {
	weatherMessage, err := GetWeather()
	if err != nil {
		return err
	}
	internalWeather, err := mapper.MapToInternalWeather(weatherMessage)
	if err != nil {
		return err
	}

	err = messaging.SendToBroker(ctx, Senders, TOPIC_WEATHER, internalWeather)
	return err
}

func GetWeather() ([]byte, error) {
	return HTTPClient.SendRequest()
}
