package database

import (
	"boobles.cloud/backend/logging"
)

type Result struct {
	LastId uint
	Ok     bool
}

// This func executes an sql query
// NOTE:
// Only use this for executing DELETE or INSERT commads!!
// And only use the constants as statement types!
// It returns an struct with the last Id and an bool to indicate success.
func ExecuteSQLStatement(statementName string, statementType int, args []any) *Result {

	query, ok := getWantedSqlStatement(statementType, statementName)

	if !ok {
		return &Result{
			LastId: 0,
			Ok:     ok,
		}
	}

	db, ok := CreateDBConn()

	if !ok {
		return &Result{
			LastId: 0,
			Ok:     ok,
		}
	}

	defer db.Close()

	// Excecute the given command with all arguments
	queryResult, err := db.Exec(query, args...)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return &Result{
			LastId: 0,
			Ok:     false,
		}
	}

	// Gets the last Id
	Id, _ := queryResult.LastInsertId()

	return &Result{
		LastId: uint(Id),
		Ok:     true,
	}
}
