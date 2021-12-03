package main

import (
	"block-info/handlers"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	val := os.Getenv(key)
	return val
}

func main() {
	polygon := GetEnv("POLYGON")

	conn, err := ethclient.Dial(polygon)
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()

	h := handlers.ClientHandler{conn}

	router.HandleFunc("/api/v1/polygon/latest", h.getLatestBlock()).Methods("GET")

	http.ListenAndServe(":8080", router)
}
