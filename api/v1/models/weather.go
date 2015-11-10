package models

type WeatherMessage struct {
	Success     string
	Message     string
	Temperature float32
	Humidity    float32
}
