/*
 * Weather internal models
 */
package models

type WeatherInternal struct {
	Latitude   *float64    `json:"latitude"`
	Longitude  *float64    `json:"longitude"`
	Measurment *Measurment `json:"hourly"`
}
