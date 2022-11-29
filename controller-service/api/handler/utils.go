package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"p2p/transfer"
)

func (a *App) tranfer(w http.ResponseWriter, userUID string, amount int) {

	transferIntent := transfer.TransferIntent{
		UserID: userUID,
		Amount: amount,
	}

	newTransfer, err := json.Marshal(&transferIntent)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://trasanction-service/api/v1/transfer", bytes.NewBuffer(newTransfer))
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

	body, err := ioutil.ReadAll(resp.Body)
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

func (a *App) checkTransfer(userID string, amount int) transfer.Transaction {

	bank := a.DB.GetInfoBank(userID)
	log.Println("this is bank info ", bank)

	url := fmt.Sprintf("http://bank-service/api/txintent?card=%s&cv=%s&amount=%d",
		bank.Card, bank.CvNumber, amount)

	resp, err := http.Get(url)
	if err != nil {
		a.InfoLog.Println("this ", err)
	}

	defer resp.Body.Close()
	// request.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(request)
	// if err != nil {
	// 	a.InfoLog.Println(err)
	// }

	// if resp.StatusCode != http.StatusAccepted {
	// 	a.ErrorLog.Fatal("no exact response")
	// }
	// a.InfoLog.Println("this ", resp.Body)

	// var transfer transfer.Transaction

	rebytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		a.InfoLog.Println("this ERROR outuil : ======== ", err)
	}

	a.InfoLog.Println("RESPONSE : ", string(rebytes))

	return transfer.Transaction{}
}

func (a *App) SendData(data string) {

	resp, err := http.Get("http://bank-service/api/health")
	if err != nil {
		a.ErrorLog.Println("error GEt ", err)
	}

	defer resp.Body.Close()

	// ansByte, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	a.ErrorLog.Println("Read all error ", err)
	// }
	a.InfoLog.Println("Done....")
	// a.InfoLog.Panicln(resp)

	// request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8083/api/health", bytes.NewBuffer([]byte(data)))
	// if err != nil {
	// 	a.ErrorLog.Println("Error send data : ", err)
	// }

	// // defer request.Body.Close()

	// request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// client := &http.Client{}
	// resp, err := client.Do(request)
	// if err == nil {
	// 	a.ErrorLog.Print("Error Response : ", err)
	// }

	// // defer resp.Body.Close()

	// respBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	a.ErrorLog.Println("Error reading body ", err)
	// }

	// a.InfoLog.Println("this is resp ", string(respBytes))

}
