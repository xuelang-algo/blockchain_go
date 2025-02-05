package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := services.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
