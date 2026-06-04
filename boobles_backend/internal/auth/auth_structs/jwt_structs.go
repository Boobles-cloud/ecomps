package authstructs

import (
	"time"

	"boobles.cloud/backend/database"
	"github.com/golang-jwt/jwt/v4"
)

// The claims of every JWT
type JWTClaimsStruct struct {
	UserId   uint
	TenantId uint
	jwt.RegisteredClaims
}

// Checks if the token is expired.
// If so it returns true.
func (j *JWTClaimsStruct) IsExpired() bool {
	currDate := time.Now()

	if currDate.Equal(j.IssuedAt.Time) || j.IssuedAt.Time.After(currDate) {
		return true
	}

	return false
}

// For getting the JWT from the database
type JWTDatabaseStruct struct {
	UserAccessId uint
	TokenVal     string
	TokenExpire  time.Time
	UserId       uint
}

// Checks if the token is expired.
// If so it returns true.
func (j *JWTDatabaseStruct) IsExpired() bool {
	currDate := time.Now()

	if currDate.Equal(j.TokenExpire) || j.TokenExpire.After(currDate) {
		return true
	}

	return false
}

// Inserts the given jwt into the database
func (j *JWTDatabaseStruct) InsertIntoDB() bool {

	result := database.ExecuteSQLStatement("InsertUserAccessToken", database.Insert, []any{j.TokenVal, j.TokenExpire, j.UserId})
	return result.Ok
}
