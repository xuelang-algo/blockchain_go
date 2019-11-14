package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !services.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := services.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := services.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
