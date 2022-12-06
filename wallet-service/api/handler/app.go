package handler

import (
	"log"
	"p2p/db/dbpostgres"
)

type App struct {
	Db                dbpostgres.DBPostgres
	InfoLog, ErrorLog *log.Logger
}
