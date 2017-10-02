package cli

import (
	"blockchain/block"
	"blockchain/chain"
	"flag"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	Bc *chain.Blockchain
}

func (cli *CLI) Run() {
	// cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		_ = addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		_ = printChainCmd.Parse(os.Args[2:])
	default:
		// cli.printUsage()
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
	cli.Bc.AddBlock(data)
	log.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		blk := bci.Next()

		log.Printf("Prev. hash: %x\n", blk.PrevBlockHash)
		log.Printf("Data: %s\n", blk.Data)
		log.Printf("Hash: %x\n", blk.Hash)
		pow := block.NewProofOfWork(blk)
		log.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		log.Println()

		if len(blk.PrevBlockHash) == 0 {
			break
		}
	}
}
