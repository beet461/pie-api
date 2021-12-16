package service

import (
	"database/sql"
	"fmt"
	"net/http"
)

func checkAccountExistence(eml string, db *sql.DB) bool {
	exists := false

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM signin_data WHERE Emails='%v'", eml))
	errorCheck(err, "QUERY")
	defer rows.Close()

	for rows.Next() {
		exists = true
	}

	return exists
}

func RegisterAccount(w http.ResponseWriter, acc UserData) (int, Account) {
	db := OpenDB()

	if checkAccountExistence(acc.Email, db) {
		return 115, Account{}
	}

	salt := RandomKey(16)
	acc.Id = RandomKey(36)
	acc.Password = hashString(acc.Password, salt)
	cust := Customise{"default", acc.Id}

	Insert(db, "signin_data", []string{acc.Email, acc.Password, acc.Firstname, acc.Lastname, acc.Id})
	Insert(db, "salts", []string{salt, acc.Id})
	Insert(db, "customisation", []string{cust.Id, "default"})

	return 199, Account{acc, cust}
}
