package postgres

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	weatherModels "github.com/nduni/correlation/weather/weather-processor/models/weather-internal"
)

func (db PGXRepository) InsertWeather(weather weatherModels.WeatherInternal) (uint64, error) {
	var weatherId uint64

	err := db.DbPool.QueryRow(context.Background(), `select create_weather($1, $2)`, weather.Latitude, weather.Longitude).Scan(&weatherId)
	if err != nil {
		return weatherId, fmt.Errorf("inserting Weather '%v' to db failed: %v", spew.Sdump(weather), err)
	}
	return weatherId, nil
}

func (db PGXRepository) FindWeather(longitude, latitude float64, measurment string) (uint64, error) {
	var weatherId uint64
	return weatherId, nil
}

func (db PGXRepository) GetWeahterById(weatherId uint64) (weatherModels.WeatherInternal, error) {
	weather := weatherModels.WeatherInternal{}
	return weather, nil
}
