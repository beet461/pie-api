package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Prints out error along with optional prefix
func errorCheck(err error, v interface{}) {
	if err != nil {
		fmt.Println(v, " ERROR CHECK:", err)
	}
}

func respond(w http.ResponseWriter, v ...interface{}) {
	enc, err := json.Marshal(v)
	errorCheck(err, "RESPONSE")
	w.Write([]byte(enc))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/signin", signin).Methods("POST")

	fmt.Println("Listening and Serving on port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
