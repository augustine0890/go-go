package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var phoneBookSize int
	fmt.Scan(&phoneBookSize)
	phoneBook := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= phoneBookSize; i++ {
		for scanner.Scan() {
			data := scanner.Text()
			textData := strings.Split(data, " ")
			name := textData[0]
			phone, _ := strconv.Atoi(textData[1])
			phoneBook[name] = phone
			break
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		name := scanner.Text()
		_, phoneResult := phoneBook[name]
		if phoneResult {
			fmt.Printf("%v=%v\n", name, phoneBook[name])
		} else {
			fmt.Println("Not found.")
		}
	}
}
