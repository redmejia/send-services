package main

import (
	"log"
	"net/http"
	"os"
	"p2p/db/dbpostgres"
	"time"
	"transaction/api/handlers"
	"transaction/api/router"
)

func main() {
	infoLog := log.New(os.Stdout, "TX-INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "TX-ERROR\t", log.Ldate|log.Ltime)

	infoLog.Println("=== CONNECTING TO DB ===")
	db, err := dbpostgres.DSNConnection(os.Getenv("DSN"))
	if err != nil {
		errorLog.Fatal("this is the error", err)
	}

	app := handlers.App{
		DB: dbpostgres.DBPostgres{
			Db:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.Router(&app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Println("==== TRANSACTION SERVICE ====")
	infoLog.Println("Service is running http://localhost:80/api/v1")
	errorLog.Fatal(srv.ListenAndServe())
}
