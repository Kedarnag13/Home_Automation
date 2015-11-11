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

type GeoLocation struct {
	Latitude    float64
	Longitude   float64
	Temperature float64
	Humidity    float64
	Windspeed   float64
	Success     string
	Message     string
}
