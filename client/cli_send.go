package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) send(from, to string, amount int32, nodeID string, mineNow bool) {
	if !services.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !services.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := services.NewBlockchain(nodeID)
	UTXOSet := services.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := services.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := services.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := services.NewCoinbaseTX(from, "")
		txs := []*services.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		services.SendTx(services.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
