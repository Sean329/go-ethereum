package main

import (
	"fmt"
	"log"
	"os"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mnemonic := os.Getenv("WEB3_PRACTICE_MNEMONIC")
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 
}