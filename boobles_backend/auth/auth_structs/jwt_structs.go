package authstructs

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// The claims of every JWT
type JWTClaimsStruct struct {
	UserId   uint
	TenantId uint
	jwt.RegisteredClaims
}

// For getting the JWT from the database
type JWTDatabaseStruct struct {
	UserAccessId uint
	TokenVal     string
	TokenExpire  time.Time
	UserId       uint
}
