package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dsn string) error {
	var err error

	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}

func CloseDB() {
	db.Close()
}
