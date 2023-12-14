package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
	"os"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

    store "github.com/Sean329/go-ethereum/contracts" 
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pk := os.Getenv("WEB3_PRACTICE_PRIVATE_KEY")
	alchemyEthereumGoerliURL:=os.Getenv("ALCHEMY_GOERLI_URL")

	client, err := ethclient.Dial(alchemyEthereumGoerliURL)
    if err != nil {
        log.Fatal(err)
    }

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
 		log.Fatal(err)
	}

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(600000) // in units
    auth.GasPrice = gasPrice

    input := "1.0"
    address, tx, instance, err := store.DeployStore(auth, client, input)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(address.Hex())   // 
    fmt.Println(tx.Hash().Hex()) // 

    _ = instance
}