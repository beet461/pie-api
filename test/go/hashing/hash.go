package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var in string
	fmt.Print("input:")
	fmt.Scan(&in)
	fmt.Printf("OUTPUT:   %x\n", sha256.Sum256([]byte(in)))
	fmt.Printf("TYPE: %T\n", sha256.Sum256([]byte(in)))
	fmt.Printf("LENGTH: %v\n", len(fmt.Sprintf("%x", sha256.Sum256([]byte(in)))))
}
