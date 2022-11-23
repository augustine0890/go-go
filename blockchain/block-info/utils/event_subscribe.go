package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogMint struct {
	To     common.Address
	Tokens *big.Int
}

func EventSubscribe() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/9c7e53d0c2984c2d8c24ee45613ff1k0")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xCC8048eF226eb2383B08949F752Cf31932d487cc")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			// var mintEvent LogMint

			// err := contractAbi.Unpack(&mintEvent, "Mint", vLog.Data)
			// if err != nil {
			// log.Fatal(err)
			// }
		}
	}
}
