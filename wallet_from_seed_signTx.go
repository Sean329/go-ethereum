package main

import (
	"log"
	"math/big"
	"fmt"
	"os"
	"context"

	"github.com/joho/godotenv"

	// "github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mnemonic := os.Getenv("WEB3_PRACTICE_MNEMONIC")
	alchemyEthereumGoerliURL:=os.Getenv("ALCHEMY_GOERLI_URL")

	client, err := ethclient.Dial(alchemyEthereumGoerliURL)
    if err != nil {
        log.Fatal(err)
    }

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account0, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account0.Address.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), account0.Address)
    if err != nil {
        log.Fatal(err)
    }
	value := big.NewInt(0)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
	chainID, err := client.NetworkID(context.Background())
		if err != nil {
  			log.Fatal(err)
		}
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account0, tx, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// spew.Dump(signedTx)

	err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}