package database

import (
	"context"
	"reflect"

	"ecomps.boobles.cloud/mango/global"
)

// Use this to execute sql Statements
func (d *DbService) ExecSqlStatment(ctx context.Context, query string, args []any) (int64, error) {

	result, err := d.handler.dbConn.ExecContext(ctx, query, args...)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// Gets one item from the database
func (d *DbService) QueryOne[T any](ctx context.Context, query string, args []any) (T, error) {

	// Just so we can return an empty struct or something like that
	var tmpVal T

	data, err := d.QueryMany[T](ctx, query, args)

	if err != nil {
		return tmpVal, err
	}

	if len(data) != 1 {
		return tmpVal, global.ERR_TOO_MUCH_RETURN
	}

	return data[0], nil
}

// Gets many items from the database
// TODO: implement this a different way!
func (d *DbService) QueryMany[T any](ctx context.Context, query string, args []any) ([]T, error) {

	var tmp T
	if reflect.TypeOf(tmp).Kind() != reflect.Struct {
		return nil, global.ERR_INVALID_TYPE
	}

	rows, err := d.handler.dbConn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	var results []T

	for rows.Next() {

		var data T
		v := reflect.ValueOf(&data).Elem()

		fieldPointer := make([]any, len(columns))

		for i := range columns {
			fieldPointer[i] = v.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fieldPointer...); err != nil {
			return results, err
		}

		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
