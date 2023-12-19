package main

import (
    "fmt"
    "log"
	"os"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/joho/godotenv"

    token "github.com/Sean329/go-ethereum/contracts/contract_ierc20" 
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

    // Evie (EVE) Address
    tokenAddress := common.HexToAddress("0x087aF2906D4cFa2E0c01910854815aA945F9B757")
    instance, err := token.NewToken(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    
	address := common.HexToAddress("0x60742b4330562fbA2b0914d901A905Ac793bFdC3")
	bal, err := instance.BalanceOf(nil, address) // Can use &bind.CallOpts{} to replace nil
	if err != nil {
  		log.Fatal(err)
	}

	fmt.Printf("wei: %s\n", bal) // "wei: 988898999999999999990000"

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
  		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
  		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
  		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Evie"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: EVE"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f\n", value) // "balance: 988899.000000"

}