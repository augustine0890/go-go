package main

import (
	"block-info/handlers"
	"block-info/middleware"
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

	var clientHandler *handlers.ClientHandler
	clientHandler = handlers.NewClientHandler(conn)

	s := router.PathPrefix("/api/v1/polygon").Subrouter()
	// Retrieves the latest block
	s.HandleFunc("/latest", clientHandler.GetLatestBlock).Methods("GET")
	// Retrieve imformation a given transaction hash
	s.HandleFunc("/get-tx", clientHandler.GetTxByHash).Methods("GET")
	logger := middleware.Logging(router)

	log.Fatal(http.ListenAndServe(":8080", logger))
}
