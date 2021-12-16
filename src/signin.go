package main

import (
	"net/http"
	"net/url"

	service "pie-api/service"
)

func login(w http.ResponseWriter, udata service.UserData) {
	code, acc := service.LoginAccount(w, udata)
	service.Respond(w, service.Response{Code: code, Account: acc})
}

func register(w http.ResponseWriter, udata service.UserData) {
	code, acc := service.RegisterAccount(w, udata)
	service.Respond(w, service.Response{Code: code, Account: acc})
}

func signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var udata service.UserData

	service.Unpack(w, r, &udata)

	m, _ := url.ParseQuery(r.URL.RawQuery)
	typ := m["type"][0] // type="login" or "register"

	if typ == "login" {
		login(w, udata)
	} else if typ == "register" {
		register(w, udata)
	} else {
		service.Respond(w, 916) // 916 - Incorrect Parameters
	}
}
