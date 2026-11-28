package httputils

import (
	"encoding/json"
	"net/http"

	"ecomps.boobles.cloud/mango/utils/logging"
)

// Decodes a json http reqest body to a struct
func JsonDeserilizeHttpRequestBody[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

// Decodes a byte slice to a given struct
func JsonDeserilizeBytes[T any](d []byte) (T, error) {
	var v T
	if err := json.Unmarshal(d, &v); err != nil {
		return v, err
	}
	return v, nil
}

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
