package handler

import (
	"io/ioutil"
	"myCloud/db"
	"myCloud/util"
	"net/http"
)

const(
	pwd_salt="*#890"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		data,err := ioutil.ReadFile("./static/view/sighup.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username:=r.Form.Get("username")
	passwd := r.Form.Get("password")

	enc_passwd := util.Sha1([]byte(passwd+pwd_salt))
	suc:=db.UserSignUp(username,enc_passwd)
	if suc{
		w.Write([]byte("SUCESS"))
	}else{
		w.Write([]byte("FAILED"))
	}
}