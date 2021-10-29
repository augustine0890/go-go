package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Models

func main() {
	// Router
	r := mux.NewRouter()

	// Router Handler (Endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
