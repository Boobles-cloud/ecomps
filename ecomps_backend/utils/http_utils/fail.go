package httputils

import (
	"net/http"

	"ecomps.boobles.cloud/backend/utils/logging"
)

// Returns a new fail handler
// This handler logs and
func NewFailHandler(w http.ResponseWriter, funcName string) func(int, error) {
	return func(status int, err error) {
		logging.Log(logging.Error, "["+funcName+"] "+err.Error())
	}
}
