package database

import (
	"os"
	"path"

	"boobles.cloud/backend/logging"
)

const (
	Insert = iota
	Delete
)

const (
	// TODO: Update new tables here!!
	TenantContext          = "tenant_statements"
	UserContext            = "user_statements"
	UserPermissionContext  = "user_permission_statements"
	UserAccessTokenContext = "user_permission_statements"
)

type Result struct {
	LastId uint
	Ok     bool
}

// This func executes an sql query
// NOTE:
// Only use this for executing DELETE or INSERT commads!!
// And only use the constants as statement contexts and types!
// It returns an struct with the last Id and an bool to indicate success.
func ExecuteSQLStatement(statementContext string, statementType int, args []any) *Result {
	// TODO: what do we do if we have multiple sql statements inside one file?

	query, ok := readSqlContent(statementType, statementContext)

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

// Gets the wanted query from a file
func readSqlContent(statementType int, statementContext string) (string, bool) {
	currDir, _ := os.Getwd()

	wantedFolder := path.Join(currDir, "sql_statements", statementContext)

	switch statementType {
	case Insert:
		// Reads the file
		byteFileData, err := os.ReadFile(path.Join(wantedFolder, "insert.sql"))

		if err != nil {
			logging.Log(logging.Error, err.Error())
			return "", false
		}

		return string(byteFileData), false
	case Delete:
		// Reads the file
		byteFileData, err := os.ReadFile(path.Join(wantedFolder, "delete.sql"))

		if err != nil {
			logging.Log(logging.Error, err.Error())
			return "", false
		}

		return string(byteFileData), false
	default:
		return "", false
	}
}
