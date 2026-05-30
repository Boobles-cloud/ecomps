package database

import (
	"reflect"

	"boobles.cloud/backend/logging"
)

// Executes a query in the database and returns the result.
// Wants:
// - T -> struct
// - queryName -> name defined in the json
// - args -> any data you want to query after :)
func QueryDatabase[T any](queryName string, queryArgs []any) ([]T, bool) {

	var results []T

	// Creates our database connection
	db, ok := CreateDBConn()

	if !ok {
		return results, ok
	}

	// Gets the wanted query
	query, ok := selectWantedSqlStatement(Select, queryName)

	if !ok {
		return results, false
	}

	defer db.Close()

	// Executes our query
	rows, err := db.Query(query, queryArgs...)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return results, false
	}

	// Gets all columns
	columns, _ := rows.Columns()

	// Loops over all data
	for rows.Next() {

		var data T

		v := reflect.ValueOf(&data).Elem()

		fieldPointers := make([]any, len(columns))

		for i := range columns {
			if i < v.NumField() {
				fieldPointers[i] = v.Field(i).Addr().Interface()
			}
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			logging.Log(logging.Error, err.Error())
			break
		}

		results = append(results, data)
	}

	return results, true
}
