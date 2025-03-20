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
	if err != nil {
		return err
	}

	err = createQuotationTable()
	return err
}


// GetDB returns a single instance of the database
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


func createQuotationTable() error {
	sql := `CREATE TABLE IF NOT EXISTS quotation (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL,
		code_in TEXT NOT NULL,
		name TEXT NOT NULL,
		high TEXT NOT NULL,
		low TEXT NOT NULL,
		var_bid TEXT NOT NULL,
		pct_change TEXT NOT NULL,
		bid TEXT NOT NULL,
		ask TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		create_date TEXT NOT NULL
		)`

	
	_, err := db.Exec(sql)
	return err
} 
