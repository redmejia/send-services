package main

import (
	"log"
	"net/http"
	"time"
	"transaction/api/router"
)

func main() {

	srv := &http.Server{
		Addr:         ":8082",
		Handler:      router.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running at http://localhost:8082/api/v1")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
