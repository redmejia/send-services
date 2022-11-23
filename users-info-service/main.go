package main

import (
	"log"
	"net/http"
	"os"
	"p2p/db/dbpostgres"
	"time"
	"users/accounts/api/handler"
	"users/accounts/api/router"
)

func main() {
	db, err := dbpostgres.DSNConnection(os.Getenv("DSN"))
	if err != nil {
		log.Fatal("error main", err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

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
		Addr:         ":8081",
		Handler:      router.Router(&app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running at http://localhost:8081/")
	err = srv.ListenAndServe()
	log.Fatal("error server ", err)
}
