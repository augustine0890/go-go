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

	fmt.Print("Enter racer name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)

	fmt.Print("Enter racer rank: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rank, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	var medal string
	if rank == 1 {
		medal = "gold"
	} else if rank == 2 {
		medal = "silver"
	} else if rank == 3 {
		medal = "bronze"
	} else {
		medal = "participant"
	}

	fmt.Println(name, "gets a", medal, "medal!")
}
