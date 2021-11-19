package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func greetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Customer!!")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Augustine", "Seoul", "103456"},
		{"David", "Hanoi", "1063566"},
		{"Juth", "London", "143436"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	// define routes
	http.HandleFunc("/customer", greetCustomer)
	http.HandleFunc("/customers", getAllCustomer)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
