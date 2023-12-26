package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"
	"os"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

    store "github.com/Sean329/go-ethereum/contracts/contract_store" 
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

    contractAddress := common.HexToAddress("0x60193B0F538D02cccA6Ab55F7EbA9d7Cb000C773") 
    
}