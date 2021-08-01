package main

import (
	"fmt"
	"strings"
)

func main() {
	var sIn string

	fmt.Printf("Enter string: ")
	fmt.Scan(&sIn)
	sIn = strings.ToLower(sIn)
	fmt.Printf("string is %s.\n", sIn)

	if strings.HasPrefix(sIn, "i") && strings.HasSuffix(sIn, "n") && strings.Contains(sIn, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Foound!")
	}
}
