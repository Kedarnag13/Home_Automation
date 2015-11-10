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
	"strings"
)

type WeatherController struct{}

var Weather WeatherController

func (w *WeatherController) Monitor(rw http.ResponseWriter, req *http.Request) {

	var geo models.GeoLocation
	vars := mux.Vars(req)
	latitude := vars["latitude"]
	lat := string(latitude)
	longitude := vars["longitude"]
	long := string(longitude)
	var lig models.Light

	keybytes, err := ioutil.ReadFile("67fb1806008cd6d8610a12f9531c4a15")
	if err != nil {
		log.Fatal(err)
	}
	key := string(keybytes)
	key = strings.TrimSpace(key)

	f, err := forecast.Get(key, lat, long, "now", forecast.CA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wind speed: %.2f\n", f.Currently.WindSpeed)

	flag.Parse()
	embd.InitGPIO()

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
