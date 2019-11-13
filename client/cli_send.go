package client

import (
	"fmt"
	"log"

	"github.com/xuelang-algo/blockchain_go"
	"github.com/xuelang-algo/blockchain_go/protos"
	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !main.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !main.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := services.NewBlockchain(nodeID)
	UTXOSet := services.UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := main.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := protos.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := protos.NewCoinbaseTX(from, "")
		txs := []*protos.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		services.sendTx(services.knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
