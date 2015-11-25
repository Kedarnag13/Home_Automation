package controllers

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/d2r2/go-dht"
	"github.com/jasonwinn/geocoder"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	forecast "github.com/mlbright/forecast/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type WeatherController struct{}

var Weather WeatherController

func (w *WeatherController) Get_information(rw http.ResponseWriter, req *http.Request) {
	// my_ip := get_my_ip()
	// query := my_ip
	city_name := "Mysore"
	lat, lng, err := geocoder.Geocode(city_name)
	if err != nil {
		panic(err)
	}
	latitude := strconv.FormatFloat(lat, 'f', 6, 64)
	longitude := strconv.FormatFloat(lng, 'f', 6, 64)

	keybytes, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		log.Fatal(err)
	}
	key := string(keybytes)
	key = strings.TrimSpace(key)

	f, err := forecast.Get(key, latitude, longitude, "now", forecast.CA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature: %v*C\n", f.Currently.Temperature)
	for i := 0; i <= 6; i++ {
		fmt.Printf("Windspeed: %v km/h \n", f.Daily.Data[i].WindSpeed)
		fmt.Printf("Climate: %v\n", f.Daily.Data[i].Icon)
		fmt.Printf("Humidity: %v%% \n", f.Daily.Data[i].Humidity*100)
		fmt.Printf("Minimum Temperature: %v*C \n", f.Daily.Data[i].TemperatureMin)
		fmt.Printf("Maximum Temperature: %v*C \n", f.Daily.Data[i].TemperatureMax)
	}
	float_lat, err := strconv.ParseFloat(latitude, 32)
	if err != nil {
		log.Fatal(err)
	}
	float_long, err := strconv.ParseFloat(longitude, 32)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(models.GeoLocation{
		Latitude:    float_lat,
		Longitude:   float_long,
		Temperature: f.Currently.Temperature,
		Humidity:    f.Currently.Humidity,
		Windspeed:   f.Currently.WindSpeed,
		Climate:     f.Daily.Icon,
		City:        city_name,
		Success:     "True",
		Message:     "Windspeed updated.",
	})
	if err != nil {
		log.Fatal(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}

// type WindController struct{}

// var Wind WindController

// func (w *WindController) Monitor_wind_velocity(rw http.ResponseWriter, req *http.Request) {

// 	// var geo models.GeoLocation
// 	vars := mux.Vars(req)
// 	latitude := vars["latitude"]
// 	lat := string(latitude)
// 	longitude := vars["longitude"]
// 	long := string(longitude)

// 	keybytes, err := ioutil.ReadFile("api_key.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	key := string(keybytes)
// 	key = strings.TrimSpace(key)

// 	f, err := forecast.Get(key, lat, long, "now", forecast.CA)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Wind Velocity is:", f.Currently.WindSpeed)
// 	float_lat, err := strconv.ParseFloat(lat, 32)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	float_long, err := strconv.ParseFloat(long, 32)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	b, err := json.Marshal(models.GeoLocation{
// 		Latitude:  float_lat,
// 		Longitude: float_long,
// 		Success:   "True",
// 		Message:   "Windspeed updated.",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.Write(b)
// }

type TemperatureController struct{}

var Temp TemperatureController

func (t *TemperatureController) Monitor_temperature_humidity(rw http.ResponseWriter, req *http.Request) {

	flag.Parse()
	embd.InitGPIO()
	// var lig models.Light

	sensorType := dht.DHT11
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensorType, 4, true, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	b, err := json.Marshal(models.WeatherMessage{
		Success:     "True",
		Message:     "Temperature and Humidity updated",
		Temperature: temperature,
		Humidity:    humidity,
	})
	if err != nil {
		log.Fatal(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
	// if temperature <= 22 {
	// 	embd.SetDirection(lig.Pin_number, embd.Out)
	// 	embd.DigitalWrite(lig.Pin_number, embd.High)
	// 	b, err := json.Marshal(models.WeatherLEDMessage{
	// 		Success: "True",
	// 		Message: "It's getting cooler!!",
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Write(b)
	// } else if temperature > 22 {
	// 	embd.SetDirection(lig.Pin_number, embd.Out)
	// 	embd.DigitalWrite(lig.Pin_number, embd.High)
	// 	b, err := json.Marshal(models.WeatherLEDMessage{
	// 		Success: "True",
	// 		Message: "It's getting hotter!!",
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Write(b)
	// }
}
