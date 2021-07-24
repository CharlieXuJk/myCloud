package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		//return upload page
		data,err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	}else if r.Method == "POST"{
		//receive files and save
		file,head,err := r.FormFile("file")
		if err!=nil{
			fmt.Printf("Failed to get data, error:%s\n", err.Error())
			return
		}
		defer file.Close()

		//create a new file, later will be used for copy
		newFile,err:=os.Create("./tmp/"+head.Filename)
		if err!=nil{
			fmt.Printf("Failed to create file, error:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err!=nil{
			fmt.Printf("Failed to save data to file, err:%s\n", err.Error())
			return
		}
	}



}