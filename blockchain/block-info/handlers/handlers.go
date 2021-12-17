package handlers

import (
	"block-info/models"
	"block-info/modules"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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

func (h *ClientHandler) GetTxByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the query params from url request
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		json.NewEncoder(w).Encode(&models.Error{
			Code:    400,
			Message: "Empty hash",
		})
		return
	}

	txHash := common.HexToHash(hash)
	_tx := modules.GetTxByHash(h.conn, txHash)
	if _tx != nil {
		json.NewEncoder(w).Encode(_tx)
		return
	}

	json.NewEncoder(w).Encode(&models.Error{
		Code:    404,
		Message: "Tx Not Found!",
	})
}

func (h *ClientHandler) GetAddressBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	address := r.URL.Query().Get("address")
	if address == "" {
		json.NewEncoder(w).Encode(&models.Error{
			Code:    400,
			Message: "Empty Address",
		})
		return
	}

	balance, err := modules.GetAddressBalance(h.conn, address)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(&models.Error{
			Code:    500,
			Message: "Internal server error",
		})
		return
	}

	json.NewEncoder(w).Encode(&models.BalanceResponse{
		Address: address,
		Balance: balance,
		Symbol:  "MATIC",
		Units:   "Wei",
	})
}
