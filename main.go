package main

import (
	"fmt"
	"github.com/christickner/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 Coin to Chris")
	bc.AddBlock("Send 3 Coins to Chris")
	bc.AddBlock("Send 2 Coins From Chris to John")

	for _, block := range bc.Blocks() {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
