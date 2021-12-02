package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Transaction data structure
type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
}

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/1c7e42d0e4274c2d8c24ee68613ff1d0")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash("0x980dec9845173ea0f7385623230feb50fa82a61af3930088bd405897286ed53e"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Pending:  pending,
		Nonce:    tx.Nonce(),
	})
}
