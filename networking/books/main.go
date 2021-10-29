package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Models
type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

}

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
