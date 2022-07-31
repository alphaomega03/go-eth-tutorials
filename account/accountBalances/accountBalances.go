package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)
var (
	NodeEndpoint = ""
)

func init() {
 NodeEndpoint = os.Getenv("NODE_ENDPOINT")
}

func main() {

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	ethClient, err := ethclient.Dial(NodeEndpoint)

	if err != nil {
		log.Fatal(err)
	}
	
	balance, err := ethClient.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Balance at current block number %d \n", balance)

	balanceAt := getBalanceAtBlockNumber(account, ethClient)
	getFormattedBalance(balanceAt)
}

func getBalanceAtBlockNumber(account common.Address, ethClient *ethclient.Client) (*big.Int){
	blockNumber := big.NewInt(5532993)

	balance, err := ethClient.BalanceAt(context.Background(), account, blockNumber)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Balance at block number %d \n", balance)

	return balance
}

func getFormattedBalance(balanceAt *big.Int) {
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Formatted balance %f \n", ethValue)
}