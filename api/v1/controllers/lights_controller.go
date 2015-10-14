package controllers

import (
	"flag"
	// "github.com/gorilla/mux"
	// "fmt"
	"github.com/kidoman/embd"
	"net/http"
	"time"
)

type LightsController struct{}

var Lights LightsController

func (l *LightsController) Toggle_led_light(rw http.ResponseWriter, req *http.Request) {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(17, embd.Out)
	embd.DigitalWrite(17, embd.High)
	time.Sleep(1 * time.Second)
	embd.SetDirection(17, embd.Out)
	embd.DigitalWrite(17, embd.Low)
}
