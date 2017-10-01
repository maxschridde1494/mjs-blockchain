package main

import (
	"blockchain/chain"
	"blockchain/restserver"
)

func main() {
	chain.Setup()
	restserver.Setup()
}
