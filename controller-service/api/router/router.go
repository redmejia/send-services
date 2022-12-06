package router

import (
	"controller/api/cors"
	"controller/api/handler"
	"net/http"
)

func Router(a *handler.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1", cors.Cors(a.ControllerTrxHandler))

	return mux
}
