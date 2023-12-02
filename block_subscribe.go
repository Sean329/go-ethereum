package main

import (
    "context"
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    alchemyEthereumMainnetWebsocket:=os.Getenv("ALCHEMY_ETHEREUM_MAINNET_WEBSOCKET")

	client, err := ethclient.Dial(alchemyEthereumMainnetWebsocket) // Need to use websocket RPC here
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("we have a connection")

    headers := make(chan *types.Header)
    sub, err := client.SubscribeNewHead(context.Background(), headers)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case header := <-headers:
            fmt.Println(header.Hash().Hex()) // 

            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println(block.Hash().Hex())        // 
            fmt.Println(block.NumberU64())   // 
            fmt.Println(block.Time())     // 
            fmt.Println(block.Nonce())             // 
            fmt.Println(len(block.Transactions())) // 
        }
    }
}