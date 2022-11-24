package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"p2p/transaction"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tx transaction.Tx

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		fmt.Println("err ", err)
	}

	resp, err := http.Get("http://localhost:8084/api/health")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

}
