package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// If the account exists "ael" is sent back (account exists login)
// If it does not "ade" is sent back (account doesn't exist)
func login(w http.ResponseWriter, r *http.Request, signin SignIn) {
	db := openDB()

	var (
		email    string
		password string
		id       string
	)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM signin_data WHERE Emails IN ('%v')", signin.Email))
	errorCheck(err)

	// The db is scanned seperately because there are specific variables needed for each scan
	// The SELECT and WHERE IN statement return all the rows where the email column matches the given email
	// There should only be one unless the register part does not function properly
	// Then there is a loop that loops through the rows (rows.Next())
	// And each time it scans the row and assigns the value to variables that correspond to the columns
	// If the password matches then it is success
	// The match variable depends on if there is a match (true) or not (false) in the db
	match := false
	for rows.Next() {
		rows.Scan(&email, &password, &id)
		if password == signin.Password {
			respond(w, SignIn{
				Email:    email,
				Password: password,
				Id:       id,
			}, "ael")
		}
		match = true
	}

	// If there is no match return ade (account doesn't exist)
	if !match {
		respond(w, "ade")
	}
}

func register(w http.ResponseWriter, r *http.Request, signin SignIn) {
	values := []string{signin.Email, signin.Password, signin.Id}
	db := openDB()

	var (
		email    string
		password string
		id       string
	)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM signin_data WHERE Emails IN ('%v')", signin.Email))
	if err != nil {
		fmt.Println("err on rows:", err)
	}

	// The db is scanned seperately because there are specific variables needed for each scan
	// The SELECT and WHERE IN statement return all the rows where the email column matches the given email
	// There should only be one unless the register part does not function properly
	// Then there is a loop that loops through the rows (rows.Next())
	// And each time it scans the row and assigns the value to variables that correspond to the columns
	// If the email matches then it is a fail as there is an existing account
	// The match variable depends on if there is a match (true) or not (false) in the db
	match := false
	for rows.Next() {
		rows.Scan(&email, &password, &id)
		if email == signin.Email {
			respond(w, "aae")
		}
		match = true
	}

	// If there is no match, the email, password and id are inserted into the db
	if !match {
		values[2] = randomKey()
		insert(db, "signin_data", values)
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signin SignIn

	body, err := ioutil.ReadAll(r.Body)
	errorCheck(err)

	jerr := json.Unmarshal([]byte(body), &signin)
	errorCheck(jerr)

	if signin.Type == "login" {
		login(w, r, signin)
	} else if signin.Type == "register" {
		register(w, r, signin)
	}
}
