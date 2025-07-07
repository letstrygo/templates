package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DatabaseURL string = "database.sqlite"
)

var (
	connection *sql.DB
)

func GetConnection() (*sql.DB, error) {
	if connection == nil {
		conn, err := sql.Open("sqlite3", DatabaseURL)
		if err != nil {
			return nil, err
		}

		connection = conn
	}

	return connection, nil
}
