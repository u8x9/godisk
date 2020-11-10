package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/u8x9/godisk/db"
	"github.com/u8x9/godisk/token"
	"github.com/u8x9/godisk/util"
)

const passwordSalt = "@!39Xdird"

type H map[string]interface{}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		buf, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			showErr(w, err)
			return
		}
		w.Write(buf)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 || len(password) < 5 {
			showErr(w, errors.New("invalid parameter"))
			return
		}
		password = util.Sha1([]byte(password + passwordSalt))
		id, err := db.UserSignUp(username, password)
		if err != nil {
			showErr(w, err)
			return
		}
		buf, err := json.Marshal(H{"id": id})
		if err != nil {
			showErr(w, err)
			return
		}
		w.Write(buf)
		return
	}
}
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		buf, err := ioutil.ReadFile("./static/view/login.html")
		if err != nil {
			showErr(w, err)
			return
		}
		w.Write(buf)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 || len(password) < 5 {
			showErr(w, errors.New("invalid parameter"))
			return
		}
		password = util.Sha1([]byte(password + passwordSalt))
		ok, err := db.UserLogin(username, password)
		if err != nil {
			showErr(w, err)
			return
		}
		if !ok {
			w.Write([]byte("用户名或密码错误"))
			return
		}
		userToken := token.GenToken(username)
		_, err = db.SetUserToken(username, userToken)
		if err != nil {
			showErr(w, err)
			return
		}
		data, _ := json.Marshal(H{"token": userToken, "msg": "登录成功"})
		w.Write(data)
	}
}
