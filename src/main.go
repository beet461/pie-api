package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/signin", signin).Methods("POST")

	fmt.Println("Listening and Serving on port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
