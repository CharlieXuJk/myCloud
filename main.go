package main

import (
	"fmt"
	"myCloud/handler"
	"net/http"
)

func main(){
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.SucceedHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)

	http.HandleFunc("/file/download", handler.DownloadHandler)

	http.HandleFunc("/user/signup", handler.SignUpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err!=nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
