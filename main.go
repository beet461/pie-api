package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var letters = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "")

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var test Test

	body, err := ioutil.ReadAll(r.Body)
	errorCheck(err)

	json.Unmarshal([]byte(body), &test)
	values := []string{test.TestRespond}

	db := openDB()
	insert(db, "test", values)
}

func randomKey() string {
	var key []string

	for i := 0; i < 51; i++ {
		key = append(key, letters[rand.Intn(len(letters))])
	}

	return strings.Join(key, "")
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println("ERROR CHECK:", err)
	}
}

func respond(w http.ResponseWriter, v ...interface{}) {
	enc, err := json.Marshal(v)
	errorCheck(err)
	w.Write([]byte(enc))
}

func customise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	router.HandleFunc("/test", test).Methods("POST")
	router.HandleFunc("/customise", customise).Methods("POST")
	router.HandleFunc("/signin", signin).Methods("POST")

	fmt.Println("Listening and Serving on port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
