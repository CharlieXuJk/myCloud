package handler

import (
	"fmt"
	"io/ioutil"
	"myCloud/db"
	"myCloud/util"
	"net/http"
	"time"
)

const(
	pwd_salt="*#890"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		data,err := ioutil.ReadFile("./static/signup.html")
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

func SignInHandler(w http.ResponseWriter, r*http.Request){

	if r.Method == http.MethodGet{
		data,err := ioutil.ReadFile("./static/signin.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	//check name and password
	r.ParseForm()
	username:=r.Form.Get("username")
	password:=r.Form.Get("password")

	encPasswd:=util.Sha1([]byte(password+pwd_salt))

	pwdChecked:=db.UserSignIn(username,encPasswd)
	if !pwdChecked{
		w.Write([]byte("FAILED"))
		return
	}
	//create token, and save the token in db
	token:=GenToken(username)
	uploadTokenResult:=db.UploadToken(username, token)
	if !uploadTokenResult {
		w.Write([]byte("FAILED"))
		return
	}
	//redirect to the homepage
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

func GenToken(username string)string{
	//md5(username+timestamp+token_salt)+timestamp[:8]  40bits
	ts:=fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix:=util.MD5([]byte(username+ts+"_tokensalt"))
	return tokenPrefix+ts[:8]
}

//look up user's info
func UserInfoHandler(w http.ResponseWriter, r*http.Request){
	r.ParseForm()
	username := r.Form.Get("username")

	//isValidToken := isTokenValid(token)
	//if !isValidToken{
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	user,err := db.GetUserInfo(username)
	if err!=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}

	resp:=util.RespMsg{
		Code:0,
		Msg:"OK",
		Data:user,
	}
	w.Write(resp.JSONBytes())
}

func isTokenValid(token string) bool{
	//to check whether token lasts too long
	//to check whether token is the same as data in db
	if len(token) != 40 {
		return false
	}
// TODO: 判断token的时效性，是否过期
// TODO: 从数据库表tbl_user_token查询username对应的token信息
// TODO: 对比两个token是否一致
	return true
}