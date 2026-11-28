package global

import "errors"

var (
	ERR_KEY_NOT_FOUND  = errors.New("ERROR: key not found!")
	ERR_STRING_EMPTY   = errors.New("ERROR: string is empty!")
	ERR_INVALID_TYPE   = errors.New("ERROR: invalid type")
	ERR_INVALID_LENGTH = errors.New("ERROR: invalid length")
)
