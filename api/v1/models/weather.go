package models

type WeatherMessage struct {
	Success     string
	Message     string
	Temperature float32
	Humidity    float32
}

type WeatherLEDMessage struct {
	Success string
	Message string
}
