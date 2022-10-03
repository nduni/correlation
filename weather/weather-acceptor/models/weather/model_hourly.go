/*
 * Weather Forecast API
 *
 * Weather Forecast
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type Hourly struct {
	Time             []string  `json:"time"`
	Temperature2m    []float64 `json:"temperature_2m"`
	Showers          []float64 `json:"showers"`
	SurfacePressure  []float64 `json:"surface_pressure,omitempty"`
	Windspeed10m     []float64 `json:"windspeed_10m,omitempty"`
	Winddirection10m []float64 `json:"winddirection_10m,omitempty"`
}
