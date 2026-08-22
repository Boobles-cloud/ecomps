package startup

import (
	"os"
	"strings"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/utils/logging"
)

// TODO: implement database updating -> future update

// Reads the included .sql file and excecute it, to setup all our tables
func SetupTabels(dh *database.DbHandler) bool {

	data, err := os.ReadFile(os.Getenv("tables_path"))

	if err != nil {
		logging.Log(logging.Error, "[Startup | SetupTabels] "+err.Error())
		return false
	}

	// Split all querys
	splited := strings.Split(string(data), ";")

	for i := range splited {

		query := strings.TrimSpace(splited[i])

		if query == "" {
			logging.Log(logging.Error, "[Startup | SetupTabels] Query Empty")
			continue
		}

		if _, err := dh.DbConnection.Exec(query); err != nil {
			logging.Log(logging.Error, "[Startup | SetupTabels] "+err.Error())
			return false
		}
	}

	return true
}
