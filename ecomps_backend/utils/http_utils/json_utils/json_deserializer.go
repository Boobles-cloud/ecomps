package jsonutils

import (
	"encoding/json"
	"net/http"
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
