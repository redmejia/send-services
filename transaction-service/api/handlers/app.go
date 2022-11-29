package handlers

import (
	"log"
	"p2p/db/dbpostgres"
)

type App struct {
	DB       dbpostgres.DBPostgres
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
