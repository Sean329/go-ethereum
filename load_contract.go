package main

import (
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

    store "github.com/Sean329/go-ethereum/contracts" 
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

    address := common.HexToAddress("0x60193B0F538D02cccA6Ab55F7EbA9d7Cb000C773")
    instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("contract is loaded")
    
	version, err := instance.Version(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(version) // "1.0"
}