package cli

import (
	"fmt"
	"log"
	"../core"
	"../transaction"
	"../utils"
	"../wallets"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain(nodeID)
	UTXOSet := transaction.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
