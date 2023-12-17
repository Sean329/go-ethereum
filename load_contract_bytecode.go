package main

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEthereumGoerliURL:=os.Getenv("ALCHEMY_GOERLI_URL")

	client, err := ethclient.Dial(alchemyEthereumGoerliURL)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x60193B0F538D02cccA6Ab55F7EbA9d7Cb000C773")
    bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(hex.EncodeToString(bytecode)) // 
}