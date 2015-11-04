package controllers

import (
	// "flag"
	// "github.com/gorilla/mux"
	// "fmt"
	// "github.com/kidoman/embd"
	// _ "github.com/kidoman/embd/host/all"
	// "bytes"
	"io/ioutil"
	"net/http"
	//"time"
	"encoding/json"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	// "log"
)

type FileController struct{}

var File FileController

func (f *FileController) Upload(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	var file models.Upload
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &file)
	if err != nil {
		panic(err)
	}
	postFile(file.Name, file.Description, file.Path, file.Target_path, rw, req)
}
