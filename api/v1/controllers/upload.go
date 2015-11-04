package controllers

import (
	// "bytes"
	// "fmt"
	"io"
	// "io/ioutil"
	// "mime/multipart"
	"net/http"
	"os"
)

func postFile(name string, description string, path string, target_path string, w http.ResponseWriter, r *http.Request) error {
	// resp, err := http.Get("0.0.0.0:3000/upload")
	// if resp == nil {
	// 	panic(err)
	// }
	// // defer resp.Body.Close()
	// out, err := os.Create(target_path)
	// if err != nil {
	// 	panic(err)
	// }
	// defer out.Close()
	// io.Copy(out, resp.Body)
	// return nil
	// bodyBuf := &bytes.Buffer{}
	// bodyWriter := multipart.NewWriter(bodyBuf)

	// // this step is very important
	// fileWriter, err := bodyWriter.CreateFormFile("uploadfile", path)
	// if err != nil {
	// 	fmt.Println("error writing to buffer")
	// 	return err
	// }

	// // open file handle
	// fh, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("error opening file")
	// 	return err
	// }

	// //iocopy
	// _, err = io.Copy(fileWriter, fh)
	// if err != nil {
	// 	return err
	// }

	// contentType := bodyWriter.FormDataContentType()
	// bodyWriter.Close()

	// resp, err := http.Post(target_path, contentType, bodyBuf)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Am here")
	// // defer resp.Body.Close()
	// resp_body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))
	switch r.Method {
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			panic(err)
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		//for each fileheader, get a handle to the actual file
		files := m.File["myfiles"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				panic(err)
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create(target_path)
			defer dst.Close()
			if err != nil {
				panic(err)
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				panic(err)
			}
		}
	}
	return nil
}
