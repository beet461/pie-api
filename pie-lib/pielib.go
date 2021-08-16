package pielib

import (
	"fmt"
)

// Prints out error along with optional prefix
func errorCheck(err error, v ...interface{}) {
	if err != nil {
		fmt.Println(v, " ERROR CHECK:", err)
	}
}
