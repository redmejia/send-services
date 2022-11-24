package main

import (
	"bank/fake-cards/cmd/api"
	"bank/fake-cards/cmd/router"
	"bank/fake-cards/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	// init new nonedb or new connection to postgres db
	var nonedb = database.NewNoneDb()

	// Accepted
	_ = nonedb.GenerateFakeCards("111122223333", 100000, 2, true)
	// Decline
	_ = nonedb.GenerateFakeCards("222233334444", 0, 4, false)

	dbByte, err := json.MarshalIndent(&nonedb.Db, "", "	")
	if err != nil {
		log.Println(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	infoLog.Println("===== TESTING CARDS ======")

	infoLog.Println(string(dbByte))

	app := api.ApiConfig{
		Port:    ":8090",
		InfoLog: infoLog,
		ErrLog:  errLog,
		DB: &database.NoneDB{
			Db:       nonedb.Db,
			InfoLog:  infoLog,
			ErrorLog: errLog,
		},
	}

	srv := &http.Server{
		Addr:    app.Port,
		Handler: router.Router(&app),
	}
	log.Println()
	log.Println("Server is running at http://localhost:8090/api/")
	log.Fatal(srv.ListenAndServe())

}
