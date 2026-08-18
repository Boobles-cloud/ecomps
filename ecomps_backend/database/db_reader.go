package database

import (
	"context"
	"reflect"

	"ecomps.boobles.cloud/backend/utils/logging"
)

// Executes a query in the database and returns the result.
// Wants:
// - T -> struct
// - queryName -> name defined in the json
// - args -> any data you want to query after :)
func QueryOne[T any](ctx context.Context, dh *DbHandler, queryName string, args ...any) (T, bool) {

	var tmpVal T
	results, ok := QueryMany[T](ctx, dh, queryName, args...)

	if !ok || len(results) == 0 {
		logging.Log(logging.Error, "[Database | QueryOne] Failed getting items! "+queryName)
		return tmpVal, false
	}

	return results[0], len(results) != 0
}

// Executes a query in the database and returns the result.
// Wants:
// - T -> struct
// - queryName -> name defined in the json
// - args -> any data you want to query after :)
func QueryMany[T any](ctx context.Context, dh *DbHandler, queryName string, args ...any) ([]T, bool) {

	queryStr, ok := dh.findQuery(queryName)

	if !ok {
		logging.Log(logging.Error, "[Database | QueryMany] Failed to load query")
		return nil, false
	}

	rows, err := dh.DbConnection.QueryContext(ctx, queryStr.QueryVal, args...)

	if err != nil {
		logging.Log(logging.Error, "[Database | QueryMany] "+err.Error())
		return nil, false
	}

	defer rows.Close()

	// Gets all colum
	columns, err := rows.Columns()

	if err != nil {
		logging.Log(logging.Error, "[Database | QueryMany] "+err.Error())
		return nil, false
	}

	var results []T

	var sample T
	v := reflect.ValueOf(&sample).Elem()
	numFields := v.NumField()

	if len(columns) > numFields {
		logging.Log(logging.Error, "[Database | QueryMany] Db columns exceeded struct fields!")
		return nil, false
	}

	for rows.Next() {
		var data T
		v := reflect.ValueOf(&data).Elem()

		fieldPointers := make([]any, len(columns))

		for i := range columns {
			fieldPointers[i] = v.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			logging.Log(logging.Error, "[Database | QueryMany] "+err.Error())
			return nil, false
		}

		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		logging.Log(logging.Error, "[Database | QueryMany] "+err.Error())
		return nil, false
	}

	return results, true
}
