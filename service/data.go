package service

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// sqlList takes an array of values and transforms them into a string with each element of the array surrounded by ' and seperated by a comma. Example: ["1", "2"] --> " '1','2' " (no double quotes)
func SqlList(v []string) string {
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

// insert() takes the table and array of values to be inserted and runs the INSERT command
func Insert(db *sql.DB, table string, values []string) {
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %v VALUES (%v)", table, SqlList(values)))
	errorCheck(err, "INSERT")

	defer stmt.Close()

	stmt.Exec()
}

// Returns *sql.DB, which is an instance of the pie-db in memory
func OpenDB() *sql.DB {
	// Executable will be built in the previous folder so './' instead of '../'
	db, err := sql.Open("sqlite3", "./db/pie-db.sqlite3")
	errorCheck(err, "OPEN")

	return db
}

func getUserSalt(db *sql.DB, id string) string {
	var salt string
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM salts WHERE Id='%v'", id))
	errorCheck(err, "SALT QUERY")

	for rows.Next() {
		rows.Scan(&salt, &id)
	}
	return salt
}

func hashString(input, salt string) string {
	hashedStr := sha256.Sum256([]byte(input + salt))
	return fmt.Sprintf("%x", hashedStr[:])
}
