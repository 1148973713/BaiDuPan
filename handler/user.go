package handler

import (
	"BaiDuPan/db"
	"BaiDuPan/util"
	"io/ioutil"
	"net/http"
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

	if len(username) > 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_password := util.Sha1([]byte(password + pwd_salt))
	suc := db.UserSignup(username, enc_password)
	if suc {
		w.Write([]byte("success"))
	} else {
		w.Write([]byte("warning"))
	}
}
