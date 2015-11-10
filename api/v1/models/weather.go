package models

type WeatherMessage struct {
	Success     string
	Message     string
	Temperature int
	Humidity    int
}
