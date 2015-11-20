package controllers

import (
	// "github.com/ninjasphere/go-samsung-tv"
	"github.com/chbmuc/lirc"
	"log"
	"net/http"
	// "time"
)

type AppliancesController struct{}

var Appliances AppliancesController

func (a *AppliancesController) Control_tv(rw http.ResponseWriter, req *http.Request) {
	ir, err := lirc.Init("/etc/lirc/lircd.conf")
	if err != nil {
		panic(err)
	}
	ir.Handle("", "KEY_POWER", keyPower)
	go ir.Run()
	// samsung.EnableLogging = true
	// tv := samsung.TV{
	// 	Host:            "192.168.1.21",
	// 	ApplicationID:   "go-samsung-tv",
	// 	ApplicationName: "Ninja Sphere         ", // XXX: Currently needs padding
	// }

	// // Once-off check if tv is online (timeout after 2 seconds)
	// if tv.Online(time.Second * 2) {
	// 	log.Println("TV is online!")
	// } else {
	// 	log.Println("TV is offline!")
	// }

	// // Continuous updates as TV goes online and offline
	// for online := range tv.OnlineState(time.Second * 5) {

	// 	if online {
	// 		log.Println("TV came online!")

	// 		// Turn the volume up when it comes online
	// 		if err := tv.SendCommand("KEY_VOLUP"); err != nil {
	// 			log.Printf("Failed to send command. Error: %s", err)
	// 		}
	// 	} else {
	// 		log.Println("TV went offline!")
	// 	}

	// }
}

func keyPower(event lirc.Event) {
	log.Println("Power Key Pressed")
}
