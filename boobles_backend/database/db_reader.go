package database

import (
	"reflect"

	"boobles.cloud/backend/logging"
)

func QueryDatabase[T any](query string, queryArgs []any) ([]T, bool) {
	var results []T

	// Creates our database connection
	db, ok := CreateDBConn()

	if !ok {
		return results, ok
	}

	defer db.Close()

	rows, err := db.Query(query, queryArgs...)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return results, false
	}

	// Gets all columns
	columns, _ := rows.Columns()

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
