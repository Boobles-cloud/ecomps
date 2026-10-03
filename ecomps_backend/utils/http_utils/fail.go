package httputils

import (
	"encoding/json"
	"net/http"

	"ecomps.boobles.cloud/backend/utils/logging"
)

type internalError struct {
	Status int    `json:"Status"`
	Err    string `json:"Error"`
}

// Returns a new fail handler
// This handler logs and
func NewFailHandler(w http.ResponseWriter, funcName string) func(int, error) {
	return func(status int, err error) {
		logging.Log(logging.Error, "["+funcName+"] "+err.Error())

		errStruct := internalError{Err: err.Error(), Status: status}

		jsonData, jsonErr := json.Marshal(errStruct)

		if jsonErr != nil {
			logging.Log(logging.Error, "["+funcName+"] "+jsonErr.Error())
			w.WriteHeader(status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(jsonData)
	}
}
