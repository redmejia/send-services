package main

import (
	"fmt"
	"p2p/wallet"
)

func main() {
	var walletPool = []wallet.Wallet{
		{WalletId: "5353-35", Balance: 10000},
		{WalletId: "0000-53", Balance: 0},
	}
	fmt.Println("w", walletPool)
}
