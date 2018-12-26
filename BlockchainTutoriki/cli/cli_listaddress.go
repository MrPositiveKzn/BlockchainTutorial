package cli

import (
	"fmt"
	"log"
	"../wallets"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := wallets.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
