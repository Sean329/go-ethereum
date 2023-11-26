package main

import (
    "context"
    "fmt"
    "log"
    "regexp"
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
	alchemyEthereumMainnetURL:=os.Getenv("ALCHEMY_ETHEREUM_MAINNET_URL")

	client, err := ethclient.Dial(alchemyEthereumMainnetURL)
    if err != nil {
        log.Fatal(err)
    }

    re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

    fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
    fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false


    // 0x Protocol Token (ZRX) smart contract address
    address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
    bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    isContract := len(bytecode) > 0

    fmt.Printf("is contract: %v\n", isContract) // is contract: true

    // a random user account address
    address = common.HexToAddress("0x60742b4330562fbA2b0914d901A905Ac793bFdC3")
    bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
    if err != nil {
        log.Fatal(err)
    }

    isContract = len(bytecode) > 0

    fmt.Printf("is contract: %v\n", isContract) // is contract: false
}