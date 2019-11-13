package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := main.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
