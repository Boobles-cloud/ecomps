package database

import (
	"reflect"
	"strings"

	"boobles.cloud/backend/logging"
)

// For updating stuff in the database.
// NOTE:
// The names of the database tables must match the names of the struct fields!
func UpdateDatabaseEntry[T any](dh *DbHandler, queryName, filterValueName string, queryData T) bool {

	// Gets our unfinished query from the json file
	unfinishedQuery, ok := dh.findQuery(queryName)

	if !ok {
		return false
	}

	valOfT := reflect.ValueOf(queryData)

	// Check if there is any pointer
	if valOfT.Kind() == reflect.Ptr {
		valOfT = valOfT.Elem()
	}

	// Checks if its a struct
	if valOfT.Kind() != reflect.Struct {
		logging.Log(logging.Error, "[Database | UpdateDatabaseEntry] queryData must be a struct or a pointer to a struct")
		return false
	}

	// Gets the struct type
	typOfT := valOfT.Type()

	// this map has the name of the given type as key and the value it has
	var tmpMap = make(map[string]any, valOfT.NumField())

	// This loops over our struct
	for i := 0; i < valOfT.NumField(); i++ {
		// Gets the name of the loop
		fieldName := typOfT.Field(i).Name
		// Gets the field value
		fieldValue := valOfT.Field(i)

		// Checks if its zero or if its valid
		if fieldValue.IsValid() && !fieldValue.IsZero() {
			// Adds to our map
			tmpMap[fieldName] = fieldValue.Interface()
		}
	}

	// Builds our query
	query, args, ok := buildQueryForUpdate(unfinishedQuery.QueryVal, filterValueName, tmpMap)

	if !ok {
		return false
	}

	if _, err := dh.DbConnection.Exec(query, args...); err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	return true
}

// Builds the query and all the arguments to it
func buildQueryForUpdate(unfinishedQuery, filterValueName string, queryArgs map[string]any) (string, []any, bool) {

	var allArgumentsInOrder []any
	// This only includes the key we want to filter after
	var filterValue = make(map[string]any, 1)

	var tmpCounter1 = 0
	for key := range queryArgs {

		if strings.EqualFold(key, filterValueName) {
			filterValue[key] = queryArgs[key]
		} else {
			// So we dont have a comma if its the end
			if tmpCounter1 == len(queryArgs) {
				unfinishedQuery += key + " = ? "
				allArgumentsInOrder = append(allArgumentsInOrder, queryArgs[key])
			} else {
				unfinishedQuery += key + " = ?, "
				allArgumentsInOrder = append(allArgumentsInOrder, queryArgs[key])
			}
		}
		tmpCounter1++
	}

	if len(filterValue) == 0 {
		logging.Log(logging.Error, "[Database | buildQueryForUpdate] Filter value is less then one!")
		return "", []any{}, false
	}

	// loops over the filter values
	var tmpCounter2 = 0
	for key := range filterValue {

		if tmpCounter2 == 0 {
			unfinishedQuery += "WHERE "
		}

		// Just check if its one and then return
		if len(filterValue) == 1 {
			unfinishedQuery += key + " = ?;"
			allArgumentsInOrder = append(allArgumentsInOrder, filterValue[key])
			break
		} else {

			if tmpCounter2 == len(filterValue) {
				// Adds the last item to the query
				unfinishedQuery += key + " = ?;"
			} else {
				// Adds our key to the query
				unfinishedQuery += key + " = ? AND"
			}
			// Add all our arguments
			allArgumentsInOrder = append(allArgumentsInOrder, filterValue[key])
		}
		tmpCounter2++
	}

	// returns our query -> its finished and all arguments
	return unfinishedQuery, allArgumentsInOrder, true
}
