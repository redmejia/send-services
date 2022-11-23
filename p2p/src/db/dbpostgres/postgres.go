package dbpostgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DBPostgres struct {
	Db                *sql.DB
	DNS               string
	InfoLog, ErrorLog *log.Logger
}

const (
	OpenConns = 10
	IdleConns = 3
	LifeTime  = 60 * time.Second
)

func Connection() (*sql.DB, error) {
	// port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	// connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	os.Getenv("HOSTNAME"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB"), os.Getenv("SSLMODE"),
	// )
	con := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("HOSTNAME"),
		os.Getenv("DBPORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("pgx", con)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(OpenConns)
	db.SetMaxIdleConns(IdleConns)
	db.SetConnMaxLifetime(LifeTime)

	return db, err
}

func DSNConnection(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(OpenConns)
	db.SetMaxIdleConns(IdleConns)
	db.SetConnMaxLifetime(LifeTime)

	return db, err

}
