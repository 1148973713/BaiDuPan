package handler

import (
	"BaiDuPan/db"
	"BaiDuPan/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const pwd_salt = ".$897"

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_password := util.Sha1([]byte(password + pwd_salt))
	suc := db.UserSignup(username, enc_password)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("WARNING"))
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	encPassword := util.Sha1([]byte(password + pwd_salt))
	//1、用户校验
	db.UserSignIn(username, encPassword)
	//2、生产方位你凭证token
	token := GetToken(username)
	updateRes := db.UpdateToken(username, token)
	if !updateRes {
		w.Write([]byte("FAILED"))
		return
	}
	//3、重定向到首页
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}

func GetToken(username string) string {
	//md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
