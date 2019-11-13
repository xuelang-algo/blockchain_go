package client

import (
	"fmt"

	"github.com/xuelang-algo/blockchain_go/services"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := services.NewBlockchain(nodeID)
	UTXOSet := services.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
