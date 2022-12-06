package router

import (
	"net/http"
	"wallet/api/handler"
)

func Route(a *handler.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/wallet", a.WalletHandler)
	mux.HandleFunc("/api/v1/share", a.ShareWalletHandler)

	return mux

}