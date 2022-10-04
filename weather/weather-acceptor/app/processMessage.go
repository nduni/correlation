package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	mapper "github.com/nduni/correlation/weather/weather-acceptor/mappers"
)

const weather_api = "https://api.open-meteo.com/v1/forecast"

func ProcessWeather() {
	weatherMessage, err := GetWeather()
	if err != nil {
		fmt.Println(err)
	}
	internalWeather, err := mapper.MapToInternalWeather(weatherMessage)
	if err != nil {
		fmt.Println(err)
	}
	internalMessage, err := json.Marshal(internalWeather)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(internalMessage))

	// sending on topic
}

func GetWeather() ([]byte, error) {
	client := resty.New()
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
	return resp.Body(), err
}
