package main

import (
	"log"
	"net/http"
	"time"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	designA := []byte("<html><head><title>The Golang Hotel</title></head><body><p>The Golang Hotel is a relaxing place !</p><p>We offer 20% discount if you call this number : <strong>1234567891</strong></p></body></html>")

	designB := []byte("<html><head><title>The Golang Hotel</title></head><body><h2>The Golang Hotel is a relaxing place !</h2><h5>We offer 20% discount if you call this number : <strong>1234567892</strong></h5></body></html>")

	minutes := time.Now().Minute()

	if minutes%2 == 0 {
		log.Println("Serving B")
		_, err := w.Write(designB)
		if err != nil {
			log.Print("Error while serving designB: ", err)
		}
	} else {
		log.Println("Serving A")
		_, err := w.Write(designA)
		if err != nil {
			log.Print("Error while serving designA", err)
		}
	}
}

func main() {
	myServer := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &myHandler{},
	}

	// launch the server
	log.Fatal(myServer.ListenAndServe())
}
