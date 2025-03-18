package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var db *sql.DB
var dbErr error

func InitDB() error {
	_, err := GetDB()
	return err
}

func GetDB() (*sql.DB, error) {
	once.Do(func() {
		db, dbErr = sql.Open("sqlite3", "quotation.db")
		if dbErr != nil {
			log.Fatal(dbErr)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
	})

	return db, dbErr
}
