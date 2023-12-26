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
    // "github.com/ethereum/go-ethereum/crypto"
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
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10181213),
		ToBlock:   nil, //nil means latest block
		Addresses: []common.Address{
		  contractAddress,
		},
	  }
	
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	  
    contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
  		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) 
		fmt.Println(vLog.BlockNumber)     
		fmt.Println(vLog.TxHash.Hex())

		event := struct {
		  Key   [32]byte
		  Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data) // Use UnpackIntoInterface instead of Unpack
		if err != nil {
		  log.Fatal(err)
		}
	  
		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar
	  }

}