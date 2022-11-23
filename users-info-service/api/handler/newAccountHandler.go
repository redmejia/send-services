package handler

import (
	"encoding/json"
	"log"
	"net/http"
	settingaccount "p2p/settingAccount"
	"users/accounts/internal/utils/security"
)

func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	var register settingaccount.Register

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Println("this json error ", err)
	}

	err = security.HashPassword(&register.Password)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	security.LastFour(&register.Card, `\d{12}`, "****-****-****-")

	ok := a.Db.InsertNewUser(&register)
	if ok {
		a.InfoLog.Println("New user was created")
	} else {
		a.ErrorLog.Println("Ooops something went wrong")
	}

	regByte, err := json.Marshal(&register)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(regByte)

}
