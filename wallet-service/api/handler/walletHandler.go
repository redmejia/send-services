package handler

import (
	"encoding/json"
	"net/http"
)

// wallet by wallet id return username wallet id and balance
func (a *App) WalletHandler(w http.ResponseWriter, r *http.Request) {

	walletID := r.URL.Query().Get("user_id")

	wallet := a.Db.GetWalletInformation(walletID)

	resp, err := json.Marshal(&wallet)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (a *App) ShareWalletHandler(w http.ResponseWriter, r *http.Request) {

	shareID := r.URL.Query().Get("share_id")

	share := a.Db.ShareWalletInfo(shareID)

	resp, err := json.Marshal(&share)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
