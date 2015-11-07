package controllers

import (
	// "flag"
	// "github.com/gorilla/mux"
	"fmt"
	// "github.com/kidoman/embd"
	// _ "github.com/kidoman/embd/host/all"
	// "bytes"
	"io/ioutil"
	"net/http"
	//"time"
	"encoding/json"
	"github.com/codeskyblue/go-sh"
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
	err = postFile(file.Name, file.Description, file.Source_path, file.Target_path)
	if err == nil {
		fmt.Println("File copied successfully!")
	}
	// Shell command to move to that directory where the song is stored
	command := sh.Command("cd", sh.Dir("/Users/kedarnag/Downloads")).Command("sudo python py/synchronized_lights.py --file=").Run()
	// Python code to play and sample the song, will be called by shell commands.
	if command == nil {
		panic(err)
	}
}
