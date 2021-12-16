package service

import (
	"fmt"
	"math/rand"
	"strings"
)

var letters = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "")

// randomKey() generates a random hexatridecimal (base 36) key of length 50 and returns it if the key does not already exist within the db
func RandomKey(length int) string {
	db := OpenDB()
	var key []string

	for i := 0; i < length; i++ {
		key = append(key, letters[rand.Intn(len(letters))])
	}
	rKey := strings.Join(key, "")

	var (
		email    string
		password string
		id       string
	)

	ids, err := db.Query(fmt.Sprintf("SELECT * FROM signin_data WHERE Id IN ('%v')", rKey))
	errorCheck(err)

	for ids.Next() {
		ids.Scan(&email, &password, &id)
		if id == rKey {
			return RandomKey(length)
		}
	}
	return rKey
}
