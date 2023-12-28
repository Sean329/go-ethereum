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

    tokenevents "github.com/Sean329/go-ethereum/contracts/contract_erc20_events" 
)

// LogTransfer ..
type LogTransfer struct {
    From   common.Address
    To     common.Address
    Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
    TokenOwner common.Address
    Spender    common.Address
    Tokens     *big.Int
}

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEthereumGoerliWebSocket:=os.Getenv("ALCHEMY_GOERLI_WEBSOCKET") // For log reading, use either websocket or regular URL for RPC, both are fine.

	client, err := ethclient.Dial(alchemyEthereumGoerliWebSocket)
    if err != nil {
        log.Fatal(err)
    }

    
}