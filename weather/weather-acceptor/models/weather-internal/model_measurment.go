/*
 * Weather internal models
 */
package models

type Measurment struct {
	Time             string   `json:"time"`
	Temperature2m    float64  `json:"temperature_2m"`
	Showers          float64  `json:"showers"`
	SurfacePressure  *float64 `json:"surface_pressure,omitempty"`
	Windspeed10m     *float64 `json:"windspeed_10m,omitempty"`
	Winddirection10m *float64 `json:"winddirection_10m,omitempty"`
}
