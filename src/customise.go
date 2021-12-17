package main

import (
	"net/http"
	service "pie-api/service"
)

func cust(w http.ResponseWriter, r *http.Request) {
	db := service.OpenDB()
	cust := service.Customise{}
	service.Unpack(w, r, &cust)

	scheme := service.FindColourScheme(db, cust.Id)
	service.Respond(w, scheme)
}
