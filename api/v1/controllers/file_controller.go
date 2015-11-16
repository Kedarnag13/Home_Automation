package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	// "github.com/kidoman/embd"
	// _ "github.com/kidoman/embd/host/all"
	"io/ioutil"
	"net/http"
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
	fmt.Println("Entering postFile")
	err = postFile(file.Name, file.Description, file.Source_path, file.Target_path)
	fmt.Println("Exiting postFile")
	if err == nil {
		fmt.Println("File copied successfully!")
	}
	// Shell command to move to that directory where the song is stored + Python code to play and sample the song, will be called by shell commands.
	command := sh.Command("cd", sh.Dir("/home/ubuntu/Desktop")).Command("sudo python py/synchronized_lights.py --file=").SetInput(file.Target_path).Run()
	if command == nil {
		panic(err)
	}
	rw.Header().Set("Connection", "close")
	fmt.Println(req.Close)
}
