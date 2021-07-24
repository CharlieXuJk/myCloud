package handler

import (
	"io"
	"io/ioutil"
	"net/http"
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
	}

}