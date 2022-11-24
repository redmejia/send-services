package api

import (
	"bank/fake-cards/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Transaction struct {
	Proceed        bool   `json:"proceed"`
	TxAmountIntent int    `json:"tx_amount_intent"`
	TxStatusCode   int    `json:"tx_status_code"`
	TxMessage      string `json:"tx_message"`
}

func (a *ApiConfig) TxIntentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// valid card 1111222233332369 | 103
		// http://localhost:8080/api/txintent?card=1222&cv=123&amount=1000
		// transaction intent

		txIntent := r.URL.Query()

		card, err := a.DB.GetInfo(txIntent.Get("card"), txIntent.Get("cv"))
		if err != nil {
			if errors.Is(err, database.ErrorNonedDBRowInResultSet) {
				var cardNotFound = struct {
					IsErorr bool   `json:"is_erorr"`
					Message string `json:"message"`
				}{
					IsErorr: true,
					Message: "no record was found",
				}

				b, _ := json.Marshal(&cardNotFound)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write(b)
			}
			return
		}

		txAmount, _ := strconv.Atoi(txIntent.Get("amount"))

		var tx Transaction

		if txAmount < card.Amount {
			tx.Proceed = card.Proceed
			tx.TxAmountIntent = txAmount
			tx.TxStatusCode = card.StatusCode
			tx.TxMessage = "Transanction Accepted"
		} else {
			tx.Proceed = card.Proceed // false
			tx.TxAmountIntent = txAmount
			tx.TxStatusCode = card.StatusCode
			tx.TxMessage = "Transanction Declined"
		}

		txByte, err := json.Marshal(tx)
		if err != nil {
			a.ErrLog.Fatal(err)
		}

		a.InfoLog.Println(string(txByte))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(txByte)
	} else {
		var methodNotImplemented = struct {
			Error   bool   `json:"error"`
			Message string `json:"message"`
		}{
			Error:   true,
			Message: fmt.Sprintf("%s is not implemented", r.Method),
		}

		b, _ := json.Marshal(&methodNotImplemented)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write(b)

	}

}
