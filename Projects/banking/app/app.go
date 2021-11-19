package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getCurrentTime).Methods(http.MethodGet)
	// define routes
	router.HandleFunc("/greet", greetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
