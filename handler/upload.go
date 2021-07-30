package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"myCloud/meta"
	"myCloud/store/ceph"
	"myCloud/util"
	"net/http"
	"os"
	"time"
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

		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"./tmp/" + head.Filename,
			Timestamp: time.Now().Format("2021-07-25 15:44:00"),
		}

		//create a new file, later will be used for copy
		newFile,err:=os.Create("./tmp/"+head.Filename)
		if err!=nil{
			fmt.Printf("Failed to create file, error:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err!=nil{
			fmt.Printf("Failed to save data to file, err:%s\n", err.Error())
			return
		}

		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		//write files to ceph
		newFile.Seek(0,0)
		data,_:=ioutil.ReadFile(newFile)
		bucket:=ceph.GetCephBucket("userfile")
		cephPath :="/ceph"/a"47.102.123.183:9080"+

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound )
	}
}

func SucceedHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Upload finished")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}