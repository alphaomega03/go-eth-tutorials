package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	NodeEndpoint = ""
)

func init() {
 NodeEndpoint = os.Getenv("NODE_ENDPOINT")
}

func main () {
	ethClient, err := ethclient.Dial(NodeEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	header, err := ethClient.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest Block number: ", header.Number.String())

	block, err := ethClient.BlockByNumber(context.Background(), header.Number)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Info about block number: ", block.Number().Uint64())
	fmt.Println("Block timestamp: ", block.Time())
	fmt.Println("Block difficulty: ", block.Difficulty().Uint64())
	fmt.Println("Block Hash: ", block.Hash().Hex())
	fmt.Println("Number of transactions in block: ", len(block.Transactions()))

}