package router

import (
	"net/http"
	"transaction/api/handlers"
)

func Router(a *handlers.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/tx/wallet", a.TransferWalletToWalletHandler)
	mux.HandleFunc("/api/v1/transfer", a.TransferHandler)

	return mux

}
