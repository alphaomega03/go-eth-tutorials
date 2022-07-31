package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var (
	NodeEndpoint = ""
)

func init() {
	NodeEndpoint = os.Getenv("NODE_ENDPOINT")
}

func main() {

	ethClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Hello client", ethClient)
}
