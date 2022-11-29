package router

import (
	"net/http"
	"transaction/api/handlers"
)

func Router(a *handlers.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1", a.TransactionHandler)
	mux.HandleFunc("/api/v1/transfer", a.TransferHandler)

	return mux

}
