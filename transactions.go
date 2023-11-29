package main

import (
    "context"
    "fmt"
    "log"
    // "math/big"
	"os"

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
	alchemyEthereumMainnetURL:=os.Getenv("ALCHEMY_ETHEREUM_MAINNET_URL")

	client, err := ethclient.Dial(alchemyEthereumMainnetURL)
    if err != nil {
        log.Fatal(err)
    }

	header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    blockNumber := header.Number
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    for _, tx := range block.Transactions() {
        fmt.Println(tx.Hash().Hex())  
		
		// Too much printing as it is printing the entire blcok -- commenting out some
        // fmt.Println(tx.Value().String())    
        // fmt.Println(tx.Gas())               
        // fmt.Println(tx.GasPrice().Uint64()) 
        // fmt.Println(tx.Nonce())             
        // fmt.Println(tx.Data())              
        // fmt.Println(tx.To().Hex())          

        chainID, err := client.NetworkID(context.Background())
        if err != nil {
            log.Fatal(err)
        }

        // if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
        //     fmt.Println(msg.From().Hex()) // Method depreciated
        // }
		if from, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println(from.Hex())
		}

        receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(receipt.Status) // 
    }

    blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
    count, err := client.TransactionCount(context.Background(), blockHash)
    if err != nil {
        log.Fatal(err)
    }

    for idx := uint(0); idx < count; idx++ {
        tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
    }

    txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
    tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
    fmt.Println(isPending)       // false
}