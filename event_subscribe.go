package main

import (
    "context"
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEthereumGoerliWebSocket:=os.Getenv("ALCHEMY_GOERLI_WEBSOCKET")

	client, err := ethclient.Dial(alchemyEthereumGoerliWebSocket)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0xEe3c6B0346a5aD69137912C6869fCaF88ba03B74")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog) // pointer to event log
        }
    }
}