package global

import "errors"

var (
	ERR_TOO_MUCH_RETURN = errors.New("ERROR: func returned too much data")
)
