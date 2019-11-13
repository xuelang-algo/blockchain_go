package client

import (
	"fmt"

	"github.com/xuelang-algo/blockchain_go"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := main.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
