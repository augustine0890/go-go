package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var arraySize int
	fmt.Scan(&arraySize)

	array := bufio.NewScanner(os.Stdin)
	array.Scan()
	data := array.Text()
	textStr := strings.Split(data, " ")

	i, j := 0, len(textStr)-1
	for i < j {
		textStr[i], textStr[j] = textStr[j], textStr[i]
		i++
		j--
	}

	for _, e := range textStr {
		fmt.Printf("%v ", e)
	}
	fmt.Printf("\n")
	// fmt.Println(textStr)
}
