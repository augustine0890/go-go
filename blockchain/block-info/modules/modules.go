package modules

import (
	"block-info/models"
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlock(client *ethclient.Client) *models.Block {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	ctx := context.Background()
	header, _ := client.HeaderByNumber(ctx, nil)
	blockNumer := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(ctx, blockNumer)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
		}
	}()

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Println(err)
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

func TransferEth(client *ethclient.Client, privKey string, to string, amount int64) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	// Load the privateKey
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	// Public address of the sending account
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Read the nonce that we should use for the account's transaction
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(amount) // in wei
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// Sending to
	toAddress := common.HexToAddress(to)
	var data []byte

	// Transaction payload
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	// Sign the transaction using the sender's private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// Broadcast the transaction to the network
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	// Transaction hash
	return signedTx.Hash().String(), nil
}
