package auth

import (
	"net/http"
	"os"
	"strings"

	authstructs "boobles.cloud/backend/auth/auth_structs"
	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
	"github.com/golang-jwt/jwt/v4"
)

// Our middleware to handle authentication
func AuthMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			logging.Log(logging.Information, "Request, to authorize user, from: "+r.RemoteAddr+" failed!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var tokenWithoutBaerer string

		if strings.Contains(token, "bearer") {
			tokenWithoutBaerer = strings.ReplaceAll(token, "bearer ", "")
		} else {
			tokenWithoutBaerer = strings.ReplaceAll(token, "Bearer ", "")
		}

		if !tokenValid(tokenWithoutBaerer) && !tokenInDB(tokenWithoutBaerer) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}

// checks if the token is valid
func tokenValid(token string) bool {

	claims := authstructs.JWTClaimsStruct{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT-Secret")), nil
	})

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	return parsedToken.Valid
}

// Checks if the token is in the database
func tokenInDB(token string) bool {

	if _, ok := database.QueryDatabase[authstructs.JWTDatabaseStruct]("SELECT * FROM UserAccesstokens WHERE TokenVal = ?;", []any{token}); !ok {
		return false
	}

	return true
}
