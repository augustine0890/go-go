package main

import (
	"fifa/data"
	"fifa/handlers"
	"net/http"
)

func main() {
	data.PrintUsage()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/winners", handlers.ListWinners)
	http.ListenAndServe(":8000", nil)
}
