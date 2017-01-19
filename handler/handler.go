package handler

import (
	"KuKaHome/Models"
	"KuKaHome/ORM"
	"KuKaHome/helper"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}

func Productlist(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {
	helper.SetGlobalHeader(w)
	r.ParseForm()
	username := ""
	password := ""
	if r.Method == "POST" {
		parms := r.PostForm
		username = parms.Get("username")
		password = parms.Get("password")
	} else if r.Method == "GET" {
		parms := r.URL.Query()
		username = parms.Get("username")
		password = parms.Get("password")
	}

	user := Models.User{
		Name:     username,
		Password: password,
	}
	result := ORM.CheckUserValid(&user)
	var message string
	var code int
	if result {
		message = "登录成功"
		code = 200
	} else {
		message = "登录失败"
		code = -1
	}
	helper.WriteResponse(w, http.StatusOK, code, map[string]string{
		"username": username,
		"message":  message,
	})
}
