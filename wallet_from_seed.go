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
	account0, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account0.Address.Hex()) // 0x60742b4330562fbA2b0914d901A905Ac793bFdC3

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account1, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account1.Address.Hex()) // 0xebee8a7d056600b4eAE4717e74Aaa9890e6692c7
}