package api

import (
	"bank/fake-cards/internal/database"
	"log"
)

type ApiConfig struct {
	Port            string
	InfoLog, ErrLog *log.Logger
	DB              database.DBFaker
}
