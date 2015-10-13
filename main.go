package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Home_Automation/api/v1/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lights/{pin_no:[0-9]+}/{status:[0-9]+}", controllers.Lights.Toggle_led_light).Methods("GET")
	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
