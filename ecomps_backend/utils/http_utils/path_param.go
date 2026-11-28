package httputils

import (
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/mango/global"
	"ecomps.boobles.cloud/mango/utils/logging"
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

// Returns the string value of a path param
func StringPathParam(r *http.Request, paramName string) (string, error) {

	p := r.PathValue(paramName)

	if p == "" {
		return "", global.ERR_STRING_EMPTY
	}

	return p, nil
}
