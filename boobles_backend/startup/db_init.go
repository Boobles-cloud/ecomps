package startup

import (
	"fmt"
	"os"
	"strings"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

// TODO: implement database updating -> future update

// Reads the included .sql file and excecute it, to setup all our tables
func SetupTabels() bool {

	// Checks if there was a first init before
	if f := os.Getenv("first-init"); f == "false" {
		return true
	}

	data, err := os.ReadFile(os.Getenv("tables_path"))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	// Split all querys
	splited := strings.Split(string(data), ";")

	conn, ok := database.CreateDBConn()

	if !ok {
		fmt.Println("Failed to connect to db! See logs for details")
		return false
	}

	defer conn.Close()

	for i := range splited {

		if _, err := conn.Exec(splited[i]); err != nil {
			logging.Log(logging.Error, err.Error())
			return false
		}
	}

	return true
}
