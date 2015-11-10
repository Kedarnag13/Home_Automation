package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Home_Automation/api/v1/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lights", controllers.Lights.Toggle_led_light).Methods("POST")
	r.HandleFunc("/upload", controllers.File.Upload).Methods("POST")
	r.HandleFunc("/monitor_temp_humidity", controllers.Temp.Monitor_temmperature_humidity).Methods("GET")
	r.HandleFunc("/monitor_wind_speed/{latitude:[0-9./0-9]+}/{longitude:[0-9./0-9]+}", controllers.Wind.Monitor_wind_velocity).Methods("GET")
	// filename := "/Users/kedarnag/Desktop/ROR - Training Material.docx"
	// postFile(filename, "0.0.0.0:3000")
	http.Handle("/", r)
	// target_url := "http://localhost:9090/upload"

	// HTTP Listening Port Raspberry Pi
	log.Println("main : Started : Listening on: http://192.168.2.112:3000 ...")
	log.Fatal(http.ListenAndServe("192.168.2.112:3000", nil))
	// Localhost
	// log.Println("main : Started : Listening on: http://0.0.0.0:3000 ...")
	// log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
