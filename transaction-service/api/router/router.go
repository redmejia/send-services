package router

import (
	"net/http"
	"transaction/api/handlers"
)

func Router() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", handlers.TransactionHandler)
	return mux

}
