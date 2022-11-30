package handler

import (
	"encoding/json"
	"net/http"
	"p2p/transfer"
)

type Payload struct {
	Acction  string                  `json:"acction"`
	Transfer transfer.TransferIntent `json:"transfer"`
}

func (a *App) ControllerHandler(w http.ResponseWriter, r *http.Request) {

	var requestPayload Payload
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		a.ErrorLog.Println("json controller ", err)
	}

	switch requestPayload.Acction {
	case "transfer_to_wallet":
		a.tranferToWallet(w, requestPayload.Transfer.UserID, requestPayload.Transfer.Amount)
	case "wallet_to_wallet":
		a.walletToWallet(w, requestPayload.Transfer.UserID, requestPayload.Transfer.DestinationWallet, requestPayload.Transfer.Amount)
	}

}
