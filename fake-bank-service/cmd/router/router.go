package router

import (
	"bank/fake-cards/cmd/api"
	"fmt"
	"net/http"
)

func Router(apiConfig *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/txintent", apiConfig.TxIntentHandler)
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Playing with docker")
	})

	return mux
}
