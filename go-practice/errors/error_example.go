package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func load() (string, error) {
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		return "nil", err
	}
	return string(data), nil
}

func main() {
	configData, err := load()
	if err != nil {
		log.Fatalf("Impossible to load config because: %s", err)
	}
	fmt.Println(configData)

	file, err := os.Open("test.csv")
	defer file.Close()

	if err != nil {
		log.Printf("impossible to open file %s", err)
		return
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
