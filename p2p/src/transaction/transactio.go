package transaction

import (
	"fmt"
	"p2p/wallet"
)

type Tx struct {
	Amount      int    `json:"amount"`
	Source      string `json:"source"` // id source or wallet id
	Destination string `json:"dest"`   // id destination or wallet id
	Note        string `json:"note"`
}

type TxStatus struct {
	WalletId     string `json:"wallet_id"`
	StatusNumber int    `json:"status_no"`
}

// check if wallet has suficient found to make the transacttion take writer
// and make new request checking transacttion
func (tx *Tx) CheckTxIntent(w wallet.Wallet) TxStatus {
	var txStatus TxStatus
	if tx.Amount < w.Balance && tx.Source == w.WalletId {
		fmt.Println("proced")
		txStatus.StatusNumber = 2
		txStatus.WalletId = w.WalletId
	} else {
		fmt.Println("unable to proced")
		txStatus.StatusNumber = 4
		txStatus.WalletId = w.WalletId
	}

	return txStatus
}
