package controllers

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	"github.com/kidoman/embd"
	// _ "github.com/kidoman/embd/host/all"
	"io/ioutil"
	"log"
	"net/http"
)

type LightsController struct{}

var Lights LightsController

func (l *LightsController) Toggle_led_light(rw http.ResponseWriter, req *http.Request) {
	flag.Parse()
	embd.InitGPIO()

	body, err := ioutil.ReadAll(req.Body)

	var lig models.Light

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &lig)
	if err != nil {
		panic(err)
	}
	fmt.Println(lig.Pin_number)

	if lig.Status == true {
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
		embd.SetDirection(lig.Pin_number, embd.Out)
		embd.DigitalWrite(lig.Pin_number, embd.Low)
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
