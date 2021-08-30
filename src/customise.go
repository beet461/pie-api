package main

import (
	"database/sql"
)

func findColors(db *sql.DB, id string) Customise {
	var scheme string

	rows, err := db.Query("SELECT * FROM customisation WHERE Account IN ('?')", id)
	errorCheck(err, "QUERY")

	for rows.Next() {
		rows.Scan(&id, &scheme)
	}

	cust := Customise{id, scheme}

	return cust
}

func newTheme(db *sql.DB, id string) Customise {
	cust := Customise{id, "default"}
	insert(db, "customisation", []string{cust.Account, cust.Colorscheme})

	return cust
}
