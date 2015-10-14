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

	body, err := ioutil.ReadAll(req.Body)
	flag := 1
	var lig models.Light

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &lig)
	if err != nil {
		panic(err)
	}
	fmt.Println(lig.Pin_number)
	if lig.Pin_number == 17 {
		embd.SetDirection(lig.Pin_number, embd.Out)
		embd.DigitalWrite(lig.Pin_number, embd.High)
		b, err := json.Marshal(models.LightMessage{
			Success: "True",
			Message: "Light blinked",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	} else {
		b, err := json.Marshal(models.LightMessage{
			Success: "false",
			Message: "Light did not blink",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}
