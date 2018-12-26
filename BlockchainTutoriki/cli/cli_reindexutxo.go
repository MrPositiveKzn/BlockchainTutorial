package cli

import (
		"fmt"
		"../core"
		"../transaction"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := core.NewBlockchain(nodeID)
	UTXOSet := transaction.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
