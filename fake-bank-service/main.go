package main

import (
	"bank/fake-cards/cmd/api"
	"bank/fake-cards/cmd/router"
	"bank/fake-cards/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// init new nonedb or new connection to postgres db
	var nonedb = database.NewNoneDb()

	// Accepted
	_ = nonedb.GenerateFakeCards("1111-2222-3333-", 100000, 2, true)
	// Decline
	_ = nonedb.GenerateFakeCards("2222-3333-4444-", 0, 4, false)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	dbByte, err := json.MarshalIndent(&nonedb.Db, "", "	")
	if err != nil {
		errLog.Fatal(err)
	}
	infoLog.Println("===== TESTING CARDS ======")

	infoLog.Println(string(dbByte))

	app := api.ApiConfig{
		Port:    ":80",
		InfoLog: infoLog,
		ErrLog:  errLog,
		DB: &database.NoneDB{
			Db:       nonedb.Db,
			InfoLog:  infoLog,
			ErrorLog: errLog,
		},
	}

	srv := &http.Server{
		Addr:         app.Port,
		Handler:      router.Router(&app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	infoLog.Println("Server is running at http://localhost:80/api/")
	err = srv.ListenAndServe()
	if err != nil {
		errLog.Fatal(err)
	}

}
