package modules

import (
	"block-info/models"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

func GetTxByHash(client *ethclient.Client, hash common.Hash) *models.Transaction {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}

	return &models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		Nonce:    tx.Nonce(),
		To:       tx.To().String(),
		Pending:  pending,
	}
}

func GetAddressBalance(client *ethclient.Client, address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "0", err
	}
	return balance.String(), nil
}
