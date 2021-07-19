package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println("Status: ", resp.Status)
	for k, v := range resp.Header {
		fmt.Println(k, ": ", v)
	}
	os.Exit(0)
}
