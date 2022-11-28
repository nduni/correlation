package mapper

import (
	"encoding/json"
	"errors"

	"github.com/nduni/correlation/common/utils"
	weatherModels "github.com/nduni/correlation/weather/weather-acceptor/models/weather"
	weatherModelsInternal "github.com/nduni/correlation/weather/weather-acceptor/models/weather-internal"
)

func MapToInternalWeather(aggrWeather []byte) ([]weatherModelsInternal.WeatherInternal, error) {
	var weathers weatherModels.Weather

	if len(aggrWeather) == 0 {
		return []weatherModelsInternal.WeatherInternal{}, errors.New("empty response body")
	}
	weathersInternal := []weatherModelsInternal.WeatherInternal{}
	err := json.Unmarshal(aggrWeather, &weathers)
	if err != nil {
		return weathersInternal, err
	}
	hourly := *weathers.Hourly
	for i, time := range hourly.Time {
		timeUTC, err := utils.ParseDateTimeToUTCString(time, weathers.Timezone)
		if err != nil {
			return weathersInternal, err
		}

		weather := weatherModelsInternal.WeatherInternal{
			Latitude:  &weathers.Latitude,
			Longitude: &weathers.Longitude,
			Measurment: &weatherModelsInternal.Measurment{
				Time:             timeUTC,
				Temperature2m:    hourly.Temperature2m[i],
				Showers:          hourly.Showers[i],
				SurfacePressure:  utils.FromSliceToPointer(hourly.SurfacePressure, i),
				Windspeed10m:     utils.FromSliceToPointer(hourly.Windspeed10m, i),
				Winddirection10m: utils.FromSliceToPointer(hourly.Winddirection10m, i),
			},
		}
		weathersInternal = append(weathersInternal, weather)
	}
	return weathersInternal, err
}
