package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2p/auth"
	"users/accounts/internal/utils/security"
)

func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	// This is practice bank information can be add after register information not need on one json payload
	var register auth.Register

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
		// this may not need on register
		successRegister.Fail.IsError = false
		successRegister.Fail.Message = ""
		regByte, err := json.Marshal(&successRegister)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(regByte)

	} else {
		// not need on register

		failRegister := auth.Success{}

		failRegister.Fail.IsError = true
		failRegister.Fail.Message = "unable to resgiter"
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

	var signin auth.Signin

	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	var hashPassword string
	ok := a.Db.GetUserAuthInfo(signin.Email, &hashPassword)

	if ok {

		ok := security.ComparePassword(hashPassword, signin.Password)

		if ok {

			successSignin := a.Db.GetAuthSuccess(signin.Email)
			successSignin.Fail.IsError = false
			successSignin.Fail.Message = ""

			regByte, err := json.Marshal(&successSignin)
			if err != nil {
				a.ErrorLog.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(regByte)

		} else {

			failSignin := auth.Success{}
			failSignin.Fail.IsError = true
			failSignin.Fail.Message = "Unable to signin wrong password or email"

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
