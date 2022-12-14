package router

import (
	"net/http"
	"wallet/api/handler"

	"wallet/api/cors"
)

func Route(a *handler.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/wallet", cors.Cors(a.WalletHandler))
	mux.HandleFunc("/api/v1/share", cors.Cors(a.ShareWalletHandler))

	return mux

}
