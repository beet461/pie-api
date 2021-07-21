package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// sqlList takes the values provided and transforms them into a string that can be used in the sql statement
// example: values consists of ["1", "2"]
// sqlList transforms it into " '1','2' "
func sqlList(v []string) string {
	var r []string
	for i := 0; i < len(v); i++ {
		toAppend := fmt.Sprintf("'%v',", v[i])
		if i == len(v)-1 {
			toAppend = fmt.Sprintf("'%v'", v[i])
		}
		r = append(r, toAppend)
	}
	return strings.Join(r, "")
}

// DB commands are prepared, with db.Prepare(). ? marks are used as unkown values.
// After creating statement (e.g. stmt), stmt.Exec() will run the statement.
// To fill the unknown values, pass the values in the order the question marks are in stmt to stmt.Exec()

// insert() takes the table and array of values to be inserted and runs the INSERT command
func insert(db *sql.DB, table string, values []string) {
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %v VALUES (%v)", table, sqlList(values)))
	fmt.Printf("INSERT INTO %v VALUES ('%v')\n\n\n\n\n\n\n\n", table, sqlList(values))

	errorCheck(err)
	stmt.Exec()
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "db/pie-db.sqlite3")
	errorCheck(err)

	return db
}
