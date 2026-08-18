package database

import (
	"context"
	"database/sql"

	"boobles.cloud/backend/logging"
)

type Result struct {
	LastId uint
	Ok     bool
}

// Executes a registered SQL statement against the pooled connection.
// NOTE: Only use this for executing DELETE or INSERT commads!!
func (dh *DbHandler) ExecuteSQLStatement(statementName string, args []any) *Result {
	return dh.execStatement(context.Background(), dh.DbConnection, statementName, args)
}

// Same as ExecuteSQLStatement, but runs inside an existing transaction
// instead of the pooled connection. Use this whenever the statement has
// to succeed/fail atomically together with other writes.
func (dh *DbHandler) ExecuteSQLStatementTx(ctx context.Context, tx *sql.Tx, statementName string, args []any) *Result {
	return dh.execStatement(ctx, tx, statementName, args)
}

// sqlExecer is implemented by both *sql.DB and *sql.Tx.
type sqlExecer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func (dh *DbHandler) execStatement(ctx context.Context, execer sqlExecer, statementName string, args []any) *Result {
	wantedQuery, ok := dh.findQuery(statementName)
	if !ok {
		logging.Log(logging.Error, "[Database | execStatement] Failed getting query")
		return &Result{LastId: 0, Ok: false}
	}

	queryResult, err := execer.ExecContext(ctx, wantedQuery.QueryVal, args...)
	if err != nil {
		logging.Log(logging.Error, "[Database | execStatement] "+err.Error())
		return &Result{LastId: 0, Ok: false}
	}

	id, _ := queryResult.LastInsertId()
	return &Result{LastId: uint(id), Ok: true}
}
