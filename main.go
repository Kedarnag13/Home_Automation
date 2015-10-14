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
	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://192.168.2.112:3000 ...")
	log.Fatal(http.ListenAndServe("192.168.2.112:3000", nil))
}
