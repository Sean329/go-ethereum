package main

import (
    "fmt"
    "log"
	"os"
	"crypto/ecdsa"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

func main() {

	// First let's obtain the signature
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKeyHex:=os.Getenv("WEB3_PRACTICE_PRIVATE_KEY") //

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        log.Fatal(err)
    }

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

    data := []byte("hello")
    dataHash := crypto.Keccak256Hash(data)
    fmt.Println(dataHash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(dataHash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(hexutil.Encode(signature)) // 0xaf71022d37b1d0b946d8016c240fdd09fa9b0457107129bdacdf5635b976577d1a6b76ff9c4ae109870803ce8e91ea988aa47a2558fc3b5180db89431003d76701


	// Verify the owner of this address: 0x60742b4330562fbA2b0914d901A905Ac793bFdC3
	addressToVerify := common.HexToAddress("0x60742b4330562fbA2b0914d901A905Ac793bFdC3")

	// Method 1
	// ---- Extract the public key bytes from the signature
	sigPublicKey, err := crypto.Ecrecover(dataHash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

	// ---- Generate the address hex and convert it to address type
	hash := sha3.NewLegacyKeccak256()
    hash.Write(sigPublicKey[1:])
    addressFromSig := common.HexToAddress(hexutil.Encode(hash.Sum(nil)[12:]))

	matches := (addressToVerify == addressFromSig)
	fmt.Println(matches) // true

	// Method 2
	// ---- Extract the public key in ECDSA type from the signature
	sigPublicKeyECDSA, err := crypto.SigToPub(dataHash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

	addressFromSig = crypto.PubkeyToAddress(*sigPublicKeyECDSA)
    matches = (addressToVerify == addressFromSig)
	fmt.Println(matches) // true

	// Method 3
	// If you already know the public key bytes to verify, then there is a method to directly compare the public key
    signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id which is the last byte of the signature
    verified := crypto.VerifySignature(publicKeyBytes, dataHash.Bytes(), signatureNoRecoverID)
    fmt.Println(verified) // true

}