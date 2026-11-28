package database

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: implement the embed stuff here!

type DbService struct {
	handler *DbHandler
}

type DbHandler struct {
	dbConn *sql.DB
}

func CreateNewDbService(dh *DbHandler) *DbService {
	return &DbService{
		handler: dh,
	}
}

// For configuration of our database stuff
func CreateDbHandler() (*DbHandler, error) {

	dbHandler := DbHandler{}

	db, err := createDBConn()

	if err != nil {
		return nil, err
	}

	// Configure our database
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(2 * time.Minute)

	// Set it
	dbHandler.dbConn = db

	return &dbHandler, nil
}

// Creates a connection to the database
// Returns the connection and a bool to indicate success
func createDBConn() (*sql.DB, error) {

	db, err := sql.Open("mysql", os.Getenv("database_connection_string"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
