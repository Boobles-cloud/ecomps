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

const (
	UserIdContextKey   = "User.Id.Context"
	TenantIdContextKey = "Tenant.Id.Context"
)

// Our middleware to handle authentication
func AuthMiddleware(dh *database.DbHandler) Middleware {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")

			if token == "" {
				logging.Log(logging.Information, "Request, to authorize user, from: "+r.RemoteAddr+" failed!")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var tokenWithoutBaerer string

			// Replaces the bearer with empty string
			if strings.Contains(token, "bearer") {
				tokenWithoutBaerer = strings.ReplaceAll(token, "bearer ", "")
			} else {
				tokenWithoutBaerer = strings.ReplaceAll(token, "Bearer ", "")
			}

			c := context.Background()

			// Checks if the token is valid and gets all claims
			tokenValid, claims := tokenValid(tokenWithoutBaerer)

			if !tokenValid && !tokenInDB(tokenWithoutBaerer, dh, c) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Adds the user and tenant Id to the context
			ct := context.WithValue(c, UserIdContextKey, strconv.Itoa(int(claims.UserId)))
			ctx := context.WithValue(ct, TenantIdContextKey, strconv.Itoa(int(claims.TenantId)))

			rCtx := r.WithContext(ctx)

			h.ServeHTTP(w, rCtx)
		})
	}
}

// checks if the token is valid
func tokenValid(token string) (bool, authstructs.JWTClaimsStruct) {

	claims := authstructs.JWTClaimsStruct{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT-Secret")), nil
	})

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false, claims
	}

	return parsedToken.Valid, claims
}

// Checks if the token is in the database
func tokenInDB(token string, dh *database.DbHandler, ctx context.Context) bool {

	if _, ok := database.QueryOne[authstructs.JWTDatabaseStruct](ctx, dh, "SelectUserAccessTokenByValue", token); !ok {
		return false
	}

	return true
}
