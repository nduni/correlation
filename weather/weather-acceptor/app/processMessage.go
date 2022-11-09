package app

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
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
	client := resty.New()
	log.Info().Msgf("sending new request to %v", weather_api)
	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"latitude":  "-21.31",
			"longitude": "-157.85",
			"timezone":  "GMT",
		}).
		SetQueryParamsFromValues(url.Values{
			"hourly": []string{"temperature_2m", "showers"},
		}).
		Get(weather_api)
	if err != nil {
		return resp.Body(), fmt.Errorf("error during request to %v: %v", weather_api, err)
	}
	if statusCode := resp.StatusCode(); statusCode != http.StatusOK {
		return resp.Body(), fmt.Errorf("response status code doesn't equal 200: %v", statusCode)
	}
	log.Info().Msgf("request to %v was succesful", resp.Request.URL)
	return resp.Body(), err
}
