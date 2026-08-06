package database

import (
	"encoding/json"
	"os"

	"boobles.cloud/backend/logging"
)

const (
	Insert = iota
	Delete
	Select
	Update
)

type QueryJsonStruct struct {
	QueryName string `json:"QueryName"`
	QueryVal  string `json:"QueryVal"`
	QueryType uint   `json:"QueryType"`
}

// Gets a wanted query by name from cache
func (dh *DbHandler) findQuery(statementName string) (QueryJsonStruct, bool) {

	for i := range dh.CachedQuerys {

		if dh.CachedQuerys[i].QueryName == statementName {
			return dh.CachedQuerys[i], true
		}
	}

	return QueryJsonStruct{}, false
}

// Gets all querys from the jsons
func readJsonFile(filePaths []string) ([]QueryJsonStruct, bool) {

	var querys = make([]QueryJsonStruct, 150)

	for i := range filePaths {

		tmpQueryRange := make([]QueryJsonStruct, 20)

		fileByteData, err := os.ReadFile(filePaths[i])

		if err != nil {
			logging.Log(logging.Error, "[Database | readingJsonFile] "+err.Error())
			return querys, false
		}

		if err := json.Unmarshal(fileByteData, &tmpQueryRange); err != nil {
			logging.Log(logging.Error, "[Database | readingJsonFile] "+err.Error())
			return querys, false
		}

		querys = append(querys, tmpQueryRange...)
	}

	return querys, len(querys) != 0
}
