package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	account "p2p/settingAccount"
	"p2p/transfer"
)

// transfer money from bank to wallet
func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {

	var txIntent transfer.TransferIntent

	err := json.NewDecoder(r.Body).Decode(&txIntent)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	bankInfo := a.DB.GetInfoBank(txIntent.UserID)

	transfer := checkTransfer(a, &bankInfo, txIntent.Amount)
	if transfer.Proceed && transfer.TxStatusCode == 2 {
		// tranfer is a accepted than make tranfer to wallet
		a.DB.TransferFromBankToWallet(txIntent.UserID, txIntent.Amount)

		transferByte, err := json.Marshal(&transfer)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(transferByte)

	} else {

		transferByte, err := json.Marshal(&transfer)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write(transferByte)
	}

}

func checkTransfer(a *App, bankInfo *account.Bank, amount int) transfer.Transaction {
	// http: //localhost:8090/api/txintent?card=****-****-****-1491&cv=172&amount=53
	// http://localhost:8083/api/txintent?card=1111-2222-3333-1871&cv=127&amount=53
	url := fmt.Sprintf("http://bank-service/api/txintent?card=%s&cv=%s&amount=%d",
		bankInfo.Card, bankInfo.CvNumber, amount)

	// url := fmt.Sprintf("http://bank-service/api/txintent?card=%s&cv=%s&amount=%d",

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		a.ErrorLog.Fatalf("bad status code expect %d but %d was recived insted ",
			http.StatusAccepted, resp.StatusCode)
	}

	var transfer transfer.Transaction
	err = json.NewDecoder(resp.Body).Decode(&transfer)
	if err != nil {
		log.Fatal(err)
	}

	return transfer

}
