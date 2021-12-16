package service

import (
	"database/sql"
	"fmt"
)

func findColourScheme(db *sql.DB, id string) Customise {
	var scheme string

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM customisation WHERE Id='%v'", id))
	errorCheck(err, "COLOURSCHEME QUERY")

	for rows.Next() {
		rows.Scan(&id, &scheme)
	}

	cust := Customise{scheme, id}

	return cust
}
