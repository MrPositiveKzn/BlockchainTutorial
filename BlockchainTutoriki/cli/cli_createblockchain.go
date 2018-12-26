package cli

import (
	"fmt"
	"log"
	"../wallets"
	"../core"
	"../transaction"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := transaction.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
