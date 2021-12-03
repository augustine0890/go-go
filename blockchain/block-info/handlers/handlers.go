package handlers

import (
	"block-info/modules"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

// ClientHandler blockchain network client instance
type ClientHandler struct {
	*ethclient.Client
}

func (h ClientHandler) getLatestBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	block := modules.GetLatestBlock(h.Client)
	json.NewEncoder(w).Encode(block)
}
