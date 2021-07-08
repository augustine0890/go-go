// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) // log a message and stop the program
	}

	input = strings.TrimSpace(input)            // remove whitespace (newlines, tabs, and regular spaces)
	grade, err := strconv.ParseFloat(input, 64) // convert string to float
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}
	fmt.Println("A grade of", grade, "is", status)
}
