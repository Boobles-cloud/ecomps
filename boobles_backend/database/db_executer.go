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
func (dh *DbHandler) ExecuteSQLStatement(statementName string, args []any) *Result {

	wantedQuery, ok := dh.findQuery(statementName)

	if !ok {
		logging.Log(logging.Error, "[Database | ExecuteSqlStatement] Failed getting query")
		return &Result{
			LastId: 0,
			Ok:     false,
		}
	}

	// Excecute the given command with all arguments
	queryResult, err := dh.DbConnection.Exec(wantedQuery.QueryVal, args...)

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
