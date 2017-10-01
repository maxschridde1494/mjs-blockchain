package chain

import (
	"blockchain/block"
	"log"
)

type Blockchain struct {
	Blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
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
