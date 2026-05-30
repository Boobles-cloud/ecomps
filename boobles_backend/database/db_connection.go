package database

import (
	"database/sql"
	"os"

	"boobles.cloud/backend/logging"
	_ "github.com/go-sql-driver/mysql"
)

// Creates a connection to the database
// Returns the connection and a bool to indicate success
// Needs to be public, because other services use this :)
func CreateDBConn() (*sql.DB, bool) {

	db, err := sql.Open("mysql", os.Getenv("Database-Conn"))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil, false
	}

	return db, true
}
