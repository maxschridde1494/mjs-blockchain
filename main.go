package main

import (
	"blockchain/chain"
	"blockchain/cli"
)

func main() {
	bc := chain.NewBlockchain()
	defer bc.Db.Close()

	commandLine := cli.CLI{bc}
	commandLine.Run()
	// restserver.Setup()
}
