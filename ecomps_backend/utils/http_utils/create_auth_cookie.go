package httputils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"ecomps.boobles.cloud/backend/database"
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthTokenCookieName = "EcompsAuthToken"
)

// Creates an auth cookie
func CreateAuthCookie(userId, tenantId uint, dh *database.DbHandler) (http.Cookie, error) {

	token, ok := createJWT(userId, tenantId, dh)

	if !ok {
		return http.Cookie{}, errors.New("Failed to creat jwt")
	}

	return http.Cookie{
		Name:     AuthTokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 3),
	}, nil
}

// Creates the jwt
func createJWT(userId, tenantId uint, dh *database.DbHandler) (string, bool) {

	claims := authstructs.JWTClaimsStruct{
		UserId:   userId,
		TenantId: tenantId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Ecomps_backend_server",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(os.Getenv("JWT-Secret")))

	if err != nil {
		logging.Log(logging.Error, "Httputils | CreateJwt"+err.Error())
		return "", false
	}

	// Insert our token into the database
	tokenDB := authstructs.JWTDatabaseStruct{
		UserAccessId: 0,
		TokenVal:     tokenSigned,
		TokenExpire:  time.Now().AddDate(0, 0, 3),
		UserId:       userId,
	}

	// Checking if ok
	if !tokenDB.InsertIntoDB(dh) {
		return "", false
	}

	return tokenSigned, true
}
