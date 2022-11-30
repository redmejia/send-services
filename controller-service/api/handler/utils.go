package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"p2p/transfer"
)

func (a *App) tranferToWallet(w http.ResponseWriter, userUID string, amount int) {

	transferIntent := transfer.TransferIntent{
		UserID: userUID,
		Amount: amount,
	}

	newTransfer, err := json.Marshal(&transferIntent)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	url := "http://trasanction-service/api/v1/transfer"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(newTransfer))
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	defer request.Body.Close()

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	if resp.StatusCode != http.StatusAccepted {
		a.ErrorLog.Fatal("no expect status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	a.InfoLog.Println("Recived ", string(body))

	var transaction transfer.Transaction

	err = json.Unmarshal(body, &transaction)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	trx, err := json.Marshal(&transaction)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(trx)

}

func (a *App) walletToWallet(w http.ResponseWriter, userUID string, toWallet string, amount int) {

	trxIntentWalletToWallet := transfer.TransferIntent{
		UserID:            userUID, // user uid is the same for wallet_id
		DestinationWallet: toWallet,
		Amount:            amount,
	}

	trxPaylod, err := json.Marshal(&trxIntentWalletToWallet)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	url := "http://trasanction-service/api/v1/tx/wallet"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(trxPaylod))
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	defer request.Body.Close()

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	// status code created?
	if resp.StatusCode != http.StatusAccepted {
		a.ErrorLog.Fatal("no expected status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	a.InfoLog.Println("RESPONSE :  ", string(body))

}
