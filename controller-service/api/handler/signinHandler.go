package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"p2p/auth"
)

func (a *App) SigninHandler(w http.ResponseWriter, r *http.Request) {

	var signin auth.Signin

	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	var successSignin auth.Success

	usersSignin(a, &signin, &successSignin)

	respSucces, _ := json.Marshal(&successSignin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respSucces)

}

func usersSignin(a *App, signin *auth.Signin, success *auth.Success) {

	body, err := json.Marshal(signin)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	url := "http://users-services/api/v1/signin"

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
