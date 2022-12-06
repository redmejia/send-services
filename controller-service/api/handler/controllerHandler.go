package handler

import (
	"encoding/json"
	"net/http"
	"p2p/transfer"
	"p2p/wallet"
)

type Payload struct {
	Acction           string `json:"acction"`
	transfer.Transfer `json:"transfer"`
	wallet.Wallet     `json:"wallet"`
}

// transaction or transefer controller
func (a *App) ControllerTrxHandler(w http.ResponseWriter, r *http.Request) {

	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.ErrorLog.Println("json controller ", err)
	}

	payloadreq, _ := json.MarshalIndent(&payload, "", "		")
	a.InfoLog.Println(string(payloadreq))

	switch payload.Acction {
	case "transfer_to_wallet":
		a.tranferToWallet(w, &payload.TranferFounds)
	case "wallet_to_wallet":
		a.walletToWallet(w, &payload.Sender, &payload.Reciver)
	case "wallet_info":
		a.GetWalletInfoById(w, payload.Wallet.UserID)
	case "share_wallet_info":
		a.ShareWalletToSender(w, payload.Wallet.ShareID)
	}

}
