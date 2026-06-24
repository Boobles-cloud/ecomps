package middleware

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"boobles.cloud/backend/database"
	authstructs "boobles.cloud/backend/internal/auth/auth_structs"
	"boobles.cloud/backend/logging"
	"github.com/golang-jwt/jwt/v4"
)

const UserIdContextKey = "User.Id.Context"

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

		tokenValid, userId := tokenValid(tokenWithoutBaerer)

		if !tokenValid && !tokenInDB(tokenWithoutBaerer) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		c := context.Background()
		ctx := context.WithValue(c, UserIdContextKey, strconv.Itoa(int(userId)))

		rCtx := r.WithContext(ctx)

		nextHandler.ServeHTTP(w, rCtx)
	})
}

// checks if the token is valid
func tokenValid(token string) (bool, uint) {

	claims := authstructs.JWTClaimsStruct{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT-Secret")), nil
	})

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false, 0
	}

	return parsedToken.Valid, claims.UserId
}

// Checks if the token is in the database
func tokenInDB(token string) bool {

	if _, ok := database.QueryDatabase[authstructs.JWTDatabaseStruct]("SelectUserAccessTokenByValue", []any{token}); !ok {
		return false
	}

	return true
}
