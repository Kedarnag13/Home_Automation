package controllers

import (
	"github.com/inando/go-lirc"
	"log"
	"net/http"
	"time"
)

type AppliancesController struct{}

var Appliances AppliancesController

func (a *AppliancesController) Control(rw http.ResponseWriter, req *http.Request) {
	client, err := lirc.New()
	if err != nil {
		return err
	}
	return client.Send("%s %s %s", "SEND_ONCE", "denon", "vol-up")
}
