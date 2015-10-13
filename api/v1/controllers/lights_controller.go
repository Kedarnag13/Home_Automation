package controllers

import (
	// "flag"
	// "github.com/gorilla/mux"
	// "github.com/kidoman/embd"
	"fmt"
	"net/http"
)

type LightsController struct{}

var Lights LightsController

func (l *LightsController) Check_light_status(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Am here")
}
