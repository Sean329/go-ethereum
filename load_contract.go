package main

import (
    "fmt"
    "log"
	"os"
    "context"
    "math/big"
    "crypto/ecdsa"
    "time"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
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

    // Read the contract
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

    // Write to the contract
    pk := os.Getenv("WEB3_PRACTICE_PRIVATE_KEY")
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
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    key := [32]byte{}
    value := [32]byte{}
    copy(key[:], []byte("test15"))
    copy(value[:], []byte("burrrrr"))

    tx, err := instance.SetItem(auth, key, value)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) 

    time.Sleep(35 * time.Second) // time is needed for the contract to update onchain, but no need to reload the instance

    result, err := instance.Items(nil, key)
    if err != nil {
        log.Fatal(err)
    }
    

    fmt.Printf("The result is: %s\n",string(result[:])) // burrrrr
}