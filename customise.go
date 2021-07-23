package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createNew(db *sql.DB) {

}

func customise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customise Customise

	body, err := ioutil.ReadAll(r.Body)
	errorCheck(err)

	err = json.Unmarshal([]byte(body), &customise)
	errorCheck(err)

	db := openDB()

	//var customiseRes Customise

	rowsC, err := db.Query(fmt.Sprintf("SELECT * FROM colors WHERE Account IN ('%v')", customise.Account))
	errorCheck(err)

	for rowsC.Next() {

	}
}
