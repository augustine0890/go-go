package handlers

import (
	"block-info/modules"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

// ClientHandler blockchain network client instance
type ClientHandler struct {
	conn *ethclient.Client
}

func NewClientHandler(conn *ethclient.Client) *ClientHandler {
	return &ClientHandler{
		conn: conn,
	}
}

func (h *ClientHandler) GetLatestBlock(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	block := modules.GetLatestBlock(h.conn)
	json.NewEncoder(w).Encode(block)
}
