package controllers

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
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

type WindController struct{}

var Wind WindController

func (w *WindController) Monitor_wind_velocity(rw http.ResponseWriter, req *http.Request) {

	// var geo models.GeoLocation
	vars := mux.Vars(req)
	latitude := vars["latitude"]
	lat := string(latitude)
	longitude := vars["longitude"]
	long := string(longitude)

	keybytes, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		log.Fatal(err)
	}
	key := string(keybytes)
	key = strings.TrimSpace(key)

	f, err := forecast.Get(key, lat, long, "now", forecast.CA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wind Velocity is:", f.Currently.WindSpeed)
	float_lat, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		log.Fatal(err)
	}
	float_long, err := strconv.ParseFloat(long, 32)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(models.GeoLocation{
		Latitude:  float_lat,
		Longitude: float_long,
		Success:   "True",
		Message:   "Windspeed updated.",
	})
	if err != nil {
		log.Fatal(err)
	}
}

type TemperatureController struct{}

var Temp TemperatureController

func (w *WindController) Monitor_temmperature_humidity(rw http.ResponseWriter, req *http.Request) {

	flag.Parse()
	embd.InitGPIO()
	var lig models.Light

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
	if temperature <= 22 {
		embd.SetDirection(lig.Pin_number, embd.Out)
		embd.DigitalWrite(lig.Pin_number, embd.High)
		b, err := json.Marshal(models.WeatherLEDMessage{
			Success: "True",
			Message: "It's getting cooler!!",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	} else if temperature > 22 {
		embd.SetDirection(lig.Pin_number, embd.Out)
		embd.DigitalWrite(lig.Pin_number, embd.High)
		b, err := json.Marshal(models.WeatherLEDMessage{
			Success: "True",
			Message: "It's getting hotter!!",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}
