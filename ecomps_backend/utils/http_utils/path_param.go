package httputils

import (
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/utils/logging"
)

// Returns the int path value
func IntPathParam(r *http.Request, paramName string) (int, error) {
	p, err := strconv.Atoi(r.PathValue(paramName))

	if err != nil {
		logging.Log(logging.Error, "[HttpUtils | IntPathParam] "+err.Error()+"; "+paramName)
		return 0, err
	}
	return p, nil
}
