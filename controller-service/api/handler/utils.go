package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"p2p/transfer"
)

// transfer from bank to wallet
func (a *App) tranferToWallet(w http.ResponseWriter, transferFounds *transfer.TranferFounds) {

	transferFound := transfer.TranferFounds{
		UserID:   transferFounds.UserID,
		WalletId: transferFounds.WalletId,
		Username: transferFounds.Username,
		Amount:   transferFounds.Amount,
	}

	newTransfer, err := json.Marshal(&transferFound)
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

func (a *App) walletToWallet(w http.ResponseWriter, sender *transfer.Sender, reciver *transfer.Reciver) {

	trxIntentWalletToWallet := transfer.TransactionIntent{Sender: *sender, Reciver: *reciver}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (a *App) GetWalletInfoById(w http.ResponseWriter, walletID string) {

	url := "http://wallet-service/api/v1/wallet?user_id=" + walletID

	resp, err := http.Get(url)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (a *App) ShareWalletToSender(w http.ResponseWriter, shareID string) {

	url := "http://wallet-service/api/v1/share?share_id=" + shareID

	resp, err := http.Get(url)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}
