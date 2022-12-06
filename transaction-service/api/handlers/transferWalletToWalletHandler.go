package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p2p/transfer"
)

// This tranfer from wallet to walllet
func (a *App) TransferWalletToWalletHandler(w http.ResponseWriter, r *http.Request) {
	var tx transfer.TransactionIntent

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		fmt.Println("err ", err)
	}

	var txStatus transfer.Transaction
	ok := a.DB.TransferWalletToWallet(&tx.Sender, &tx.Reciver)
	if ok {
		_ = a.DB.InsertTrxsRecordWalletToWallet(&tx.Sender, &tx.Reciver)

		txStatus.Proceed = true
		txStatus.TxMessage = "Accepted"
		txStatus.TxStatusCode = 2

	} else {
		txStatus.Proceed = false
		txStatus.TxMessage = "Declined"
		txStatus.TxStatusCode = 4
	}

	resp, err := json.Marshal(&txStatus)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)

}
