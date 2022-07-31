package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"github.com/ethereum/go-ethereum/core/types"
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
	ethClient, err := ethclient.Dial(NodeEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	header, err := ethClient.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest Block number: ", header.Number.String())

	blockNumber := big.NewInt(15250706)
	
	block, err := ethClient.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	readBlockTransactions(block)
	readFromAddress(ethClient, block)
	readTransactionReceipt(ethClient, block)
	readBlockFromHash(ethClient, "0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	readSingleTransaction(ethClient, "0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
}

func readSingleTransaction(ethClient *ethclient.Client, hash string) {
	txHash := common.HexToHash(hash)
	tx, isPending, err := ethClient.TransactionByHash(context.Background(), txHash)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}

func readBlockFromHash(ethClient *ethclient.Client, hash string) {
	blockHash := common.HexToHash(hash)
	count, err := ethClient.TransactionCount(context.Background(), blockHash)

	if err != nil {
		log.Fatal(err)
	}
	for  idx := 0; uint(idx) < count; idx++ {
		tx, err := ethClient.TransactionInBlock(context.Background(), blockHash, uint(idx))
	
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(tx.Hash().Hex())
	}

}


func readTransactionReceipt(ethClient *ethclient.Client, block *types.Block) {
	for _, tx := range block.Transactions() {
		receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
	}
}

func readFromAddress(ethClient *ethclient.Client, block *types.Block) {
	chainId, err := ethClient.NetworkID(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), big.NewInt(0)); err != nil {
			fmt.Println(msg.From().Hex())
		}
 	}
}

func readBlockTransactions(block *types.Block) {
	fmt.Println("Info about block number: ", block.Number().Uint64())
	fmt.Println("Block timestamp: ", block.Time())
	fmt.Println("Block difficulty: ", block.Difficulty().Uint64())
	fmt.Println("Block Hash: ", block.Hash().Hex())
	fmt.Println("Number of transactions in block: ", len(block.Transactions()))

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        
		fmt.Println(tx.Value().String())    
		fmt.Println(tx.Gas())               
		fmt.Println(tx.GasPrice().Uint64()) 
		fmt.Println(tx.Nonce())             
		fmt.Println(tx.Data())              
		fmt.Println(tx.To().Hex())          
	}
}