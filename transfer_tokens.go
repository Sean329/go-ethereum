package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
	"os"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    // "github.com/ethereum/go-ethereum/crypto/sha3"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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

	toAddress := common.HexToAddress("0xeF145f88397f2E5aa64b56Ec49e67e7Ed2644FC8")
	tokenAddress := common.HexToAddress("0x087aF2906D4cFa2E0c01910854815aA945F9B757")
	value := big.NewInt(0) // in wei (1 eth)
    
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256(transferFnSignature)
	// hash.Write(transferFnSignature)
	methodID := hash[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x000000000000000000000000ef145f88397f2e5aa64b56ec49e67e7ed2644fc8

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) //0x00000000000000000000000000000000000000000000003635c9adc5dea00000


    var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	  })
	if err != nil {
		log.Fatal(err)
	}
	  
	fmt.Println(gasLimit) //21644

    tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s \n", signedTx.Hash().Hex()) //tx sent: 0x7ed721861e8a255596f07d9fd54844815c1dd5a3717a16adac7e2f786efbd617
	
}