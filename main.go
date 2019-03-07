package main

import (
	"flag"
	"fmt"
	"github.com/christickner/go-blockchain/blockchain"
	"log"
	"os"
)

type CLI struct {
	bc *blockchain.Blockchain
}

func (cli *CLI) Run() {
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Print("yu ")
			log.Fatal(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Print("yohb ")
			log.Fatal(err)
		}
	default:
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	log.Printf("%x Successfully added, new tip\n", cli.bc.Tip)
}

func (cli *CLI) printChain() {
	i := cli.bc.Iterator()

	for h := 0; h < 10; h++ {
		block := i.Next()

		log.Printf("nexting: %d\n", h)

		if block == nil {
			fmt.Println("no next block: ")
			break
		}

		log.Printf("%x: Prev. Hash\n", block.PrevBlockHash)
		log.Printf("Data: %s\n", block.Data)
		log.Printf("Nonce: %d\n", block.Nonce)
		log.Printf("%x: Hash\n", block.Hash)
		fmt.Println()
	}
}

func main() {
	bc := blockchain.NewBlockchain()

	cli := CLI{
		bc: bc,
	}
	cli.Run()
	bc.CloseDb()
}
