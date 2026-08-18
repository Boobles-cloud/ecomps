package database

import (
	"database/sql"
	"os"
	"path"
	"time"

	"ecomps.boobles.cloud/backend/logging"
	_ "github.com/go-sql-driver/mysql"
)

type DbHandler struct {
	DbConnection *sql.DB
	CachedQuerys []QueryJsonStruct
}

// For configuration of our database stuff
func CreateDbHandler() (*DbHandler, bool) {

	dbHandler := DbHandler{}

	db, ok := createDBConn()

	if !ok {
		return nil, false
	}

	// Configure our database
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(2 * time.Minute)

	// Set it
	dbHandler.DbConnection = db

	currentDir, _ := os.Getwd()
	wantedFolder := path.Join(currentDir, "database", "database_sql_statements")

	// Get all querys from the json files
	querys, ok := readJsonFile([]string{path.Join(wantedFolder, "delete_querys.json"), path.Join(wantedFolder, "insert_querys.json"),
		path.Join(wantedFolder, "select_querys.json"), path.Join(wantedFolder, "update_querys.json")})

	if !ok {
		return nil, false
	}

	dbHandler.CachedQuerys = querys

	return &dbHandler, true
}

// Creates a connection to the database
// Returns the connection and a bool to indicate success
func createDBConn() (*sql.DB, bool) {

	db, err := sql.Open("mysql", os.Getenv("Database-Conn"))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil, false
	}

	return db, true
}
