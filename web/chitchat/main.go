package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("ChitChat %s started at %s\n", version(), config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// define in route_auth.go
	mux.HandleFunc("/login", login)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
