package controllers

import (
	// "github.com/ninjasphere/go-samsung-tv"
	"encoding/json"
	"fmt"
	"github.com/chbmuc/lirc"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
)

type AppliancesController struct{}

var Appliances AppliancesController

func (a *AppliancesController) Control_tv(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	var tv models.Remote
	// var event lirc.Event
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &tv)
	if err != nil {
		panic(err)
	}
	ir, err := lirc.Init("/var/run/lirc/lircd")
	if err != nil || ir == nil {
		panic(err)
	}
	fmt.Printf("Code:%v", tv.Key_code)
	fmt.Printf("Name:%v", tv.Key_name)

	reply := ir.Command(`LIST micromax ""`)
	fmt.Println(reply.DataLength, reply.Data)

	err = ir.Send("micromax PROG-SCAN")
	if err != nil {
		fmt.Println(err)
	}
	err = ir.SendLong("micromax VOL-DOWN", time.Duration(time.Second*3))
	if err != nil {
		fmt.Println(err)
	}

	// ir.Handle("", "KEY_POWER", keyPower(tv.Key_code, tv.Key_name))

	// ir.Handle("", "KEY_POWER", func keyPower(tv.Key_code, tv.Key_name))
	// ir.Handle("", "KEY POWER", func(w http.ResponseWriter, r *http.Request) {
	// 	keyPower(w, r, tv.Key_code, tv.Key_name)
	// })
	// go ir.Run()
	// reply := ir.Command(`LIST micromax ""`)
	// fmt.Println(reply.DataLength, reply.Data)
	// err = ir.Send("micromax KEY_POWER")
	// if err != nil {
	// 	log.Println(err)
	// }
	// ir.Handle("/", "KEY_1", key1(tv.Key_code, tv.Key_name))
	// go ir.Run()
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
	// code, name := keyPower(event, tv.Key_code, tv.Key_name, tv.Remote_name)
	// fmt.Println("Code:", code)
	// fmt.Println("Name:", name)
	// go ir.Run()
	// err = ir.Send("micromax KEY_POWER")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func keyPower(event lirc.Event) {
	fmt.Println("Power Key Pressed")
}

func keyTV(event lirc.Event) {
	fmt.Println("TV Key Pressed")
}

func keyAll(event lirc.Event) {
	fmt.Println(event)
}

// func keyPower(event lirc.Event, code string, name string, rname string) (string, string) {
// 	fmt.Println("Event:", event)
// 	fmt.Println("Power Key Pressed")
// 	return code, name
// }

// func keyPower(event lirc.Event, write http.ResponseWriter, read *http.Request, code string, name string) {
// 	fmt.Println(write, read)
// 	fmt.Println("Code:%v", code)
// 	fmt.Println("Name:%v", name)
// 	log.Println("Power Key Pressed")
// }

// func key1(event lirc.Event, code string, name string) {
// 	fmt.Println("Code:%v", code)
// 	fmt.Println("Name:%v", name)
// 	log.Println("Key 1 Pressed")
// }
