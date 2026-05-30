package database

import (
	"encoding/json"
	"os"
	"path"

	"boobles.cloud/backend/logging"
)

const (
	Insert = iota
	Delete
	Select
	Update
)

type queryJsonStruct struct {
	QueryName string `json:"QueryName"`
	QueryVal  string `json:"QueryVal"`
	QueryType uint   `json:"QueryType"`
}

// Gets the wanted sql statement
func selectWantedSqlStatement(statementType int, statementName string) (string, bool) {
	currDir, _ := os.Getwd()

	wantedFolder := path.Join(currDir, "database_sql_statements")

	switch statementType {
	case Insert:

		jsonStruct, ok := readJsonFile(path.Join(wantedFolder, "insert_querys.json"), statementName)

		if !ok {
			return "", false
		}

		return jsonStruct.QueryVal, true
	case Delete:

		jsonStruct, ok := readJsonFile(path.Join(wantedFolder, "delete_querys.json"), statementName)

		if !ok {
			return "", false
		}

		return jsonStruct.QueryVal, true
	case Select:
		jsonStruct, ok := readJsonFile(path.Join(wantedFolder, "delete_querys.json"), statementName)

		if !ok {
			return "", false
		}

		return jsonStruct.QueryVal, true
	default:
		return "", false
	}
}

// Gets the query from the json file
func readJsonFile(filePath, statementName string) (*queryJsonStruct, bool) {

	byteFileData, err := os.ReadFile(filePath)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return &queryJsonStruct{}, false
	}

	var querys = make([]queryJsonStruct, 20)

	if err := json.Unmarshal(byteFileData, &querys); err != nil {
		logging.Log(logging.Error, err.Error())
		return &queryJsonStruct{}, false
	}

	for i := range querys {
		if querys[i].QueryName == statementName {
			return &querys[i], true
		}
	}
	return &queryJsonStruct{}, false
}
