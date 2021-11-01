package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Models
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Init books
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameters
	for _, b := range books {
		if b.ID == params["id"] {
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{
		ID:     "1",
		ISBN:   "3475983",
		Title:  "Book 1",
		Author: &Author{FirstName: "John", LastName: "Smith"},
	})

	books = append(books, Book{
		ID:     "2",
		ISBN:   "3434583",
		Title:  "Book 2",
		Author: &Author{FirstName: "Roberson", LastName: "Kimberly"},
	})

	books = append(books, Book{
		ID:     "3",
		ISBN:   "345383",
		Title:  "Book 3",
		Author: &Author{FirstName: "Judith", LastName: "Davis"},
	})

	// Router Handler (Endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
