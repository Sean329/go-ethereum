package main

import (
    "context"
    "fmt"
    "log"
	"os"

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

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(header.Number.String()) 

    blockNumber := header.Number
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(block.NumberU64())     
    fmt.Println(block.Time())       
    fmt.Println(block.Difficulty().Uint64()) 
    fmt.Println(block.Hash().Hex())          
    fmt.Println(len(block.Transactions()))   

    count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(count) 
}