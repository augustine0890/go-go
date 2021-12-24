package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sizes := make(chan int)
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}

func responseSize(url string, channel chan int) {
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
	channel <- len(body) // The size of the slice of bytes

}
