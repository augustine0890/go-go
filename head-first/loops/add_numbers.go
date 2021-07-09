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
	reader := bufio.NewReader(os.Stdin)
	var sum float64
	more := "y"

	for more == "y" {
		fmt.Print("Enter a number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number

		fmt.Print("Add more? [Y/n] ")
		more, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		more = strings.ToLower(strings.TrimSpace(more))
		if more == "n" {
			break
		}
	}

	fmt.Println("Sum is", sum)
}
