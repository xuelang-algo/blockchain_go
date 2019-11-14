package client

import (
	"fmt"

	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := services.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
