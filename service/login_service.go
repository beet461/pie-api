package service

import (
	"database/sql"
	"fmt"
	"net/http"
)

func checkAccountDetails(eml string, db *sql.DB) (bool, UserData) {
	var (
		email     string
		password  string
		firstname string
		lastname  string
		id        string
	)

	// This statement selects all the rows where the email matches the given email
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM signin_data WHERE Emails='%v'", eml))
	errorCheck(err, "QUERY")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&email, &password, &firstname, &lastname, &id)
		return true, UserData{email, password, firstname, lastname, id}
	}

	return false, UserData{}
}

func LoginAccount(w http.ResponseWriter, acc UserData) (int, Account) {

	db := OpenDB()

	exists, det := checkAccountDetails(acc.Email, db)
	if !exists {
		return 145, Account{}
	}

	hpwd := hashString(acc.Password, getUserSalt(db, det.Id))
	acc.Id = det.Id

	if det.Password == hpwd {
		return 199, Account{acc, findColourScheme(db, det.Id)}
	} else {
		return 163, Account{}
	}
}
