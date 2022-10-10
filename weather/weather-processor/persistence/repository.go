package persistence

import weatherModels "github.com/nduni/correlation/weather/weather-processor/models/weather-internal"

type Repository interface {
	WeatherSaver
	WeatherFinder
	WeatherGetter
}

type WeatherSaver interface {
	InsertWeather(weather weatherModels.WeatherInternal)
}

type WeatherFinder interface {
	FindWeather(longitude, latitude float64, measurment string) (uint64, error)
}

type WeatherGetter interface {
	GetWeahterById(weatherId uint64) (weatherModels.WeatherInternal, error)
}
