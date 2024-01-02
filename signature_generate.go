package main

import (
    "fmt"
    "log"
	"os"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKeyHex:=os.Getenv("WEB3_PRACTICE_PRIVATE_KEY") //

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        log.Fatal(err)
    }

    data := []byte("hello")
    hash := crypto.Keccak256Hash(data)
    fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

    signature, err := crypto.Sign(hash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(hexutil.Encode(signature)) // 0xaf71022d37b1d0b946d8016c240fdd09fa9b0457107129bdacdf5635b976577d1a6b76ff9c4ae109870803ce8e91ea988aa47a2558fc3b5180db89431003d76701
}