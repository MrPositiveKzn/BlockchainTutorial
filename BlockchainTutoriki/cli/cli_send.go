package cli

import (
	"fmt"
	"log"
	"../wallets"
	"../core"
	"../transaction"
	"../server"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallets.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallets.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := core.NewBlockchain(nodeID)
	UTXOSet := transaction.UTXOSet{bc}
	defer bc.Db.Close()

	wallets, err := wallets.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := transaction.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := transaction.NewCoinbaseTX(from, "")
		txs := []*transaction.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		server.SendTx(server.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
