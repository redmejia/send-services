package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"p2p/auth"
)

func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var register auth.Register

	json.NewDecoder(r.Body).Decode(&register)

	var successRegister auth.Success
	userRegister(a, &register, &successRegister)

	respSuccess, _ := json.Marshal(&successRegister)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respSuccess)

}

func userRegister(a *App, register *auth.Register, success *auth.Success) {

	body, err := json.Marshal(register)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	url := "http://users-services/api/v1/register"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		a.InfoLog.Fatal(err)
	}

	err = json.Unmarshal(respBody, success)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

}
