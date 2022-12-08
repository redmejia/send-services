package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2p/auth"
	account "p2p/auth"
	"users/accounts/internal/utils/security"
)

func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	var register account.Register

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

		successRegister := a.Db.GetAuthSuccess(register.Email)
		regByte, err := json.Marshal(&successRegister)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(regByte)
	} else {

		failRegister := auth.Fail{IsError: true, Message: "Unable to register"}
		regByte, err := json.Marshal(&failRegister)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(regByte)

	}
}

func (a *App) SigninHandler(w http.ResponseWriter, r *http.Request) {

	var signin account.Signin

	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	var hashPassword string
	ok := a.Db.GetUserAuthInfo(signin.Email, &hashPassword)

	if ok {

		ok, err := security.ComparePassword(hashPassword, signin.Password)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		if ok {

			successSignin := a.Db.GetAuthSuccess(signin.Email)
			regByte, err := json.Marshal(&successSignin)
			if err != nil {
				a.ErrorLog.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(regByte)

		} else {

			failSignin := auth.Fail{IsError: true, Message: "Unable to Signin"}
			regByte, err := json.Marshal(&failSignin)
			if err != nil {
				a.ErrorLog.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(regByte)

		}

	}

}
