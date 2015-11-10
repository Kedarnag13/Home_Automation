package models

type WeatherMessage struct {
	Success     string
	Message     string
	Temperature float64
	Humidity    float64
}
