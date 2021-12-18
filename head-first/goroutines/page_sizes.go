package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	// Release the network connection once the 'main' function exists
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Conversion from a slice of `byte` values to a `string`
	fmt.Println(string(body))
}
