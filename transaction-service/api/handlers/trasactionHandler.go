package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p2p/transaction"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tx transaction.Tx

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		fmt.Println("err ", err)
	}

	fmt.Println("tx ", tx)
	fmt.Println("this")
}
