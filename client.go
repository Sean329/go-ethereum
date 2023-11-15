package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"os"
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
    _ = client // we'll use this in the upcoming sections
}