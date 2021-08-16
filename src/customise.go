package main

import (
	"database/sql"
	"net/http"
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

// Responds with color scheme of user, not used by client program, only for outside queries
func customise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//findColors()
}
