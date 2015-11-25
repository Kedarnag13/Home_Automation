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
	WeatherForWeek	DaysOfWeek
	Latitude    float64
	Longitude   float64
	Temperature float64
	Humidity    float64
	Windspeed   float64
	Climate     string
	City        string
	Success     string
	Message     string
}

type DailyWeather struct {
	Windspeed float64
	Climate string
	Humidity float64
	Minimum_temp float64
	Maximum_temp float64
}

type DaysOfWeek struct {
	Sunday	DailyWeather
	Monday	DailyWeather
	Tuesday	DailyWeather
	Wednesday	DailyWeather
	Thursday	DailyWeather
	Friday		DailyWeather
	Saturday	DailyWeather
}