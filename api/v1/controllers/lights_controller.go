package controllers

import (
	"flag"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

type LightsController struct{}

var Lights LightsController

func (l *LightsController) Check_light_status(rw http.ResponseWriter, req *http.Request) {

}
