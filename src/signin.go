package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	pielib "pie-api/pie-lib"
)

func login(w http.ResponseWriter, signin SignIn) {
	db := openDB()
	var (
		email     string
		password  string
		firstname string
		lastname  string
		id        string
	)

	// This statement selects all the rows where the email matches the given email
	rows, err := db.Query("SELECT * FROM signin_data WHERE Emails IN ('?')", signin.Email)
	errorCheck(err, "QUERY")
	defer rows.Close()

	// If the account exists and the passwords match ael (account exists login), all the needed data, and customisation values are responded with
	for rows.Next() {
		rows.Scan(&email, &password, &firstname, &lastname, &id)
		if password == signin.Password {
			respond(w, "ael", SignIn{email, password, firstname, lastname, id}, findColors(db, id))
			return
		}
		respond(w, "piw", password, signin.Password)
		return
	}

	// If there is no match return ade (account doesn't exist)
	respond(w, "ade")
}

func register(w http.ResponseWriter, signin SignIn) {
	db := openDB()

	// This sql statement selects rows from the db where the Email column has a value matching the email entered by the client
	rows, err := db.Query("SELECT * FROM signin_data WHERE Emails IN ('?')", signin.Email)
	errorCheck(err, "QUERY")
	defer rows.Close()

	// If there are matching rows aae (account already exists) is set as the response and current process is ended
	for rows.Next() {
		respond(w, "aae")
		return
	}

	// If there is no match, the email, password and id are inserted into the db
	// A new instance of the default colours is also inserted into the db
	// After this srp (succeful registration proceed) is responded along with account details and customisation details
	signin.Id = pielib.RandomKey(db)
	insert(db, "signin_data", []string{signin.Email, signin.Password, signin.Firstname, signin.Lastname, signin.Id})
	cust := Customise{signin.Id, "default"}
	// Password and id are set to null, so it isn't freely available on client's device
	signin.Password = ""
	signin.Id = ""
	respond(w, "srp", signin, cust)
}

// Signin function is for either logging in or registering an account
func signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signin SignIn

	body, err := ioutil.ReadAll(r.Body)
	errorCheck(err, "BODY READING")

	jerr := json.Unmarshal([]byte(body), &signin)
	errorCheck(jerr, "UNMARSHAL")

	// The raw query is parsed giving a map of parameters
	m, _ := url.ParseQuery(r.URL.RawQuery)
	// The type parameter's value is what type of request this is (login or register)
	typ := m["type"][0]

	if typ == "login" {
		login(w, signin)
	} else if typ == "register" {
		register(w, signin)
	} else {
		// If parameter of url is incorrect, the program responds with ipv (incorrect parameter value)
		respond(w, "ipv")
	}
}
