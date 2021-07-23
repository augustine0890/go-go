package main

import (
	"fmt"
	"net/http"
	"time"
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
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:         config.Address,
		Handler:      mux,
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}
	server.ListenAndServe()
}
