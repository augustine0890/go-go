package modules

import (
	"block-info/models"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlock(client *ethclient.Client) *models.Block {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := context.Background()
	header, _ := client.HeaderByNumber(ctx, nil)
	blockNumer := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(ctx, blockNumer)
	if err != nil {
		log.Fatal(err)
	}

	blockRes := &models.Block{
		BlockNumber:       block.Number().Int64(),
		Timestamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: block.Transactions().Len(),
		Transactions:      []models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		blockRes.Transactions = append(blockRes.Transactions, models.Transaction{
			Hash:     tx.Hash().String(),
			Value:    tx.Value().String(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice().Uint64(),
			Nonce:    tx.Nonce(),
			To:       tx.To().String(),
		})
	}
	return blockRes
}
