package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/1c7e42d0e4274c2d8c24ee68613ff1d0")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// get block number
	header, err := conn.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	before, _ := new(big.Int).SetString("1000", 10)
	blockNumber := big.NewInt(0).Sub(header.Number, before)
	block, err := conn.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block time: ", block.Time())
	fmt.Println("Block difficulty: ", block.Difficulty())
	fmt.Println("Block hash: ", block.Hash())
	fmt.Println("Number of TXs: ", block.Transactions().Len())
	fmt.Println("First TX's Hash of block: ", block.Transactions()[0].Hash())
	fmt.Println()

	block, err = conn.BlockByHash(ctx, block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block time: ", block.Time())
	fmt.Println("Block difficulty: ", block.Difficulty())
	fmt.Println("Block hash: ", block.Hash())
	fmt.Println("Number of TXs: ", block.Transactions().Len())
	fmt.Println("First TX's Hash of block: ", block.Transactions()[0].Hash())

}
