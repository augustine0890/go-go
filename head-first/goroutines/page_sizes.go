package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	time.Sleep(3 * time.Second)
}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
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
	// fmt.Println(string(body))
	// Page URLs and page sizes in bytes
	fmt.Println(len(body)) // The size of the slice of bytes

}
