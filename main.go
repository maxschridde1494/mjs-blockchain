package main

import (
	"blockchain/chain"
)

func main() {
	bc := chain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
}
