package pielib

import (
	"database/sql"
	"math/rand"
	"net/mail"
	"regexp"
	"strings"
)

var letters = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "")

// randomKey() generates a random hexatridecimal (base 36) key of length 50 and returns it if the key does not already exist within the db
func RandomKey(db *sql.DB) string {
	var key []string

	for i := 0; i < 51; i++ {
		key = append(key, letters[rand.Intn(len(letters))])
	}
	rKey := strings.Join(key, "")

	var (
		email    string
		password string
		id       string
	)

	ids, err := db.Query("SELECT * FROM signin_data WHERE Id IN ('?')", rKey)
	errorCheck(err)

	for ids.Next() {
		ids.Scan(&email, &password, &id)
		if id == rKey {
			return RandomKey(db)
		}
	}
	return rKey
}

// Results of checking, success status and error message
type Results struct {
	Success bool
	Error   string
}

// All results in one struct, each field is of type Results{}
type Checks struct {
	Email     Results
	Password  Results
	Firstname Results
	Lastname  Results
}

// General checks such as empty field, spaces or special characters
func generalChecks(v string) Results {
	if v == "" || len(strings.Split(v, "")) > 1 {
		return Results{false, "Error: Empty field or spaces present"}
	} else if match, _ := regexp.MatchString(`[!#~$%^&*()+|\x60¬¦\]"£<>':;\=,]`, v); match {
		return Results{false, "Error: Special characters present. Allowed special characters are - and  . and _"}
	}
	return Results{true, ""}
}

// Specific checks for email, such as syntax
func email(email string) Results {
	if r := generalChecks(email); !r.Success {
		return Results{false, r.Error}
	} else if _, err := mail.ParseAddress(email); err != nil {
		return Results{false, "Error: Incorrect syntax"}
	}

	return Results{true, ""}
}

// Specific checks for password such as length
func password(pwd string) Results {
	if r := generalChecks(pwd); !r.Success {
		return Results{false, r.Error}
	} else if len(pwd) > 20 {
		return Results{false, "Error: Password is over 20 characters"}
	} else if len(pwd) < 5 {
		return Results{false, "Error: Password length must be higher than 5 characters"}
	}
	return Results{true, ""}
}

// Checks for namessuch as length
func name(name string) Results {
	if r := generalChecks(name); !r.Success {
		return Results{false, r.Error}
	} else if len(name) > 50 {
		return Results{false, "Error: Name is over 50 characters, please check your name is spelt correctly"}
	}
	return Results{true, ""}
}

// The verify function returns all the results of the checks
func Verify(em, pa, firstnm, lastnm string) Checks {
	res := Checks{
		Email:     email(em),
		Password:  password(pa),
		Firstname: name(firstnm),
		Lastname:  name(lastnm),
	}

	return res
}
