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

	// Mock USDC for Project Seahawk token address
    contractAddress := common.HexToAddress("0xEe3c6B0346a5aD69137912C6869fCaF88ba03B74")
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(10199440), //The block where the contract was deployed
        ToBlock:   nil,
        Addresses: []common.Address{
            contractAddress,
        },
    }

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	  
    contractAbi, err := abi.JSON(strings.NewReader(string(tokenevents.TokeneventsABI)))
	if err != nil {
  		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)

	for _, vLog := range logs {
        fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
        fmt.Printf("Log Index: %d\n", vLog.Index)

        switch vLog.Topics[0].Hex() {
        case logTransferSigHash.Hex():
            fmt.Printf("Log Name: Transfer\n")

            var transferEvent LogTransfer

            err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
            transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

            fmt.Printf("From: %s\n", transferEvent.From.Hex())
            fmt.Printf("To: %s\n", transferEvent.To.Hex())
            fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

        case logApprovalSigHash.Hex():
            fmt.Printf("Log Name: Approval\n")

            var approvalEvent LogApproval

            err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
            approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

            fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
            fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
            fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
        }

        fmt.Printf("\n\n")
    }
}