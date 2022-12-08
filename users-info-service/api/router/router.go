package router

import (
	"net/http"
	"users/accounts/api/handler"
)

func Router(a *handler.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/register", a.NewAccountHandler)
	mux.HandleFunc("/api/v1/signin", a.SigninHandler)

	return mux
}
