package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)             // remove whitespace (newlines, tabs, and regular spaces)
	number, err := strconv.ParseFloat(input, 64) // parse the string as a float64
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Printf("The student's grade is %.2f and the student's status is %s.\n", grade, status)
}
