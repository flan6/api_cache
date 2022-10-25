package sql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBConnection struct {
	DB *sql.DB
}

func NewDBConnection(dbPath string) *DBConnection {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	return &DBConnection{
		DB: db,
	}
}
