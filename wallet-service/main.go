package main

import (
	"log"
	"net/http"
	"os"
	"p2p/db/dbpostgres"
	"time"
	"wallet/api/handler"
	"wallet/api/router"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	db, err := dbpostgres.DSNConnection(os.Getenv("DSN"))
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := handler.App{
		Db: dbpostgres.DBPostgres{
			Db:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.Route(&app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Println("Wallet service is running at http://localhost:80/api/v1")
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
