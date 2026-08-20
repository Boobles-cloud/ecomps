package jsonutils

import (
	"encoding/json"
	"net/http"

	"ecomps.boobles.cloud/backend/utils/logging"
)

// Response with json
func RespondWithJson(w http.ResponseWriter, status int, data any) bool {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logging.Log(logging.Error, "[JsonUtils | RespondWithJson] "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
	return true
}
