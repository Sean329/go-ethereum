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
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

    exchange "github.com/Sean329/go-ethereum/contracts/contract_exchange_events" 
)

type LogFill struct {
    Maker                  common.Address
    Taker                  common.Address
    FeeRecipient           common.Address
    MakerToken             common.Address
    TakerToken             common.Address
    FilledMakerTokenAmount *big.Int
    FilledTakerTokenAmount *big.Int
    PaidMakerFee           *big.Int
    PaidTakerFee           *big.Int
    Tokens                 [32]byte
    OrderHash              [32]byte
}

type LogCancel struct {
    Maker                     common.Address
    FeeRecipient              common.Address
    MakerToken                common.Address
    TakerToken                common.Address
    CancelledMakerTokenAmount *big.Int
    CancelledTakerTokenAmount *big.Int
    Tokens                    [32]byte
    OrderHash                 [32]byte
}

type LogError struct {
    ErrorID   uint8
    OrderHash [32]byte
}

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEthereumMainnetURL:=os.Getenv("ALCHEMY_ETHEREUM_MAINNET_URL") // For log reading, use either websocket or regular URL for RPC, both are fine.

	client, err := ethclient.Dial(alchemyEthereumMainnetURL)
    if err != nil {
        log.Fatal(err)
    }

	
}