package router

import (
	"bank/fake-cards/cmd/api"
	"bank/fake-cards/cmd/cors"
	"encoding/json"
	"log"
	"net/http"
)

// {"echo" : "hello world", "send": true}
type Sounds struct {
	Echo string `json:"echo"`
	Send bool   `json:"send"`
}

func Router(apiConfig *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()
	// http: //localhost:8083/api/txintent?card=1111-2222-3333-1871&cv=127&amount=53
	mux.HandleFunc("/api/txintent", cors.Cors(apiConfig.TxIntentHandler))
	mux.HandleFunc("/api/health", cors.Cors(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			log.Println("recived")

			// var echo Sounds
			// err := json.NewDecoder(r.Body).Decode(&echo)
			// if err != nil {
			// 	log.Println("JSON : ", err)
			// }

			// log.Println("recived ", echo)

			var echoback = struct {
				Msg string `json:"msg"`
			}{
				Msg: "message was reviced",
			}

			b, _ := json.Marshal(&echoback)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(b)

		}
	}))

	return mux
}
