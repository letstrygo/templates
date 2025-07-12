package templates

import (
	"database/sql"
	"io"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	file string
	*sql.DB
}

const (
	DatabaseURL string = "database.sqlite"
)

func NewConnectionWithPath(path string) (*Connection, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &Connection{path, conn}, err
}

func NewConnection() (*Connection, error) {
	return NewConnectionWithPath(DatabaseURL)
}

const (
	RemoteURL            string = "https://raw.githubusercontent.com/letstrygo/templates/refs/heads/main/dist/database.sqlite"
	TemporaryDatabaseURL string = "./temp.sqlite"
)

func NewRemoteConnection() (*Connection, error) {
	// Download the remote database file
	resp, err := http.Get(RemoteURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create the local file
	out, err := os.Create(TemporaryDatabaseURL)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	// Open a connection to the downloaded database
	return NewConnectionWithPath(TemporaryDatabaseURL)
}
