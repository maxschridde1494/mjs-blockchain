package chain

import (
	"blockchain/block"
	"log"
	"strconv"
)

type Blockchain struct {
	Blocks []*block.Block
}

var BC *Blockchain

func Setup() {
	if BC != nil {
		log.Println("Blockchain is already setup")
		return
	}
	BC = NewBlockchain()
	log.Println("Blockchain is initialized")
	BC.AddBlock("Max Schridde")
	BC.AddBlock("Miles Schridde")
	for _, b := range BC.Blocks {
		pow := block.NewProofOfWork(b)
		log.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	}
}

func (bc *Blockchain) AddBlock(data string) {
	log.Println("Adding new block with data: %s", data)
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	log.Printf("Adding Block")
}
func NewGenesisBlock() *block.Block {
	log.Println("Creating Genesis Block")
	return block.NewBlock("Genesis Block", []byte{})
}
func NewBlockchain() *Blockchain {
	log.Println("Initializing Blockchain")
	return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}
