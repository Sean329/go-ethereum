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

    alchemyEthereumMainnetURL:=os.Getenv("ALCHEMY_ETHEREUM_MAINNET_URL")

	client, err := ethclient.Dial(alchemyEthereumMainnetURL)
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
            fmt.Println(block.Number().Uint64())   // 
            fmt.Println(block.Time().Uint64())     // 
            fmt.Println(block.Nonce())             // 
            fmt.Println(len(block.Transactions())) // 
        }
    }
}