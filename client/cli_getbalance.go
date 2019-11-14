package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go/services"
	"github.com/xuelang-algo/blockchain_go/utils"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !services.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := services.NewBlockchain(nodeID)
	UTXOSet := services.UTXOSet{bc}
	defer bc.DB.Close()

	var balance int32
	balance = 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
