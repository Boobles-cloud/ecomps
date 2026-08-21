package middleware

import (
	"context"
	"net/http"
	"os"

	"ecomps.boobles.cloud/backend/database"
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	"ecomps.boobles.cloud/backend/utils/logging"
	"github.com/golang-jwt/jwt/v4"
)

const (
	UserIdContextKey   = "UserIdContext"
	TenantIdContextKey = "TenantIdContext"
)

// Our middleware to handle authentication
func AuthMiddleware(dh *database.DbHandler) Middleware {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodOptions {
				h.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie(httputils.AuthTokenCookieName)

			logging.Log(logging.Information, cookie.Value)

			if err != nil {
				logging.Log(logging.Error, "[Middleware | AuthMiddleware] "+err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			reqCtx := r.Context()

			// Checks if the token is valid and gets all claims
			tokenValid, claims := tokenValid(cookie.Value)

			if !tokenValid && !tokenInDB(cookie.Value, dh, reqCtx) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Adds the user and tenant Id to the context
			ctx := context.WithValue(reqCtx, UserIdContextKey, int(claims.UserId))
			ctx = context.WithValue(ctx, TenantIdContextKey, int(claims.TenantId))

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// checks if the token is valid
func tokenValid(token string) (bool, authstructs.JWTClaimsStruct) {

	claims := authstructs.JWTClaimsStruct{}

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT-Secret")), nil
	})

	if err != nil {
		logging.Log(logging.Error, "[AuthMiddleware | tokenValid] "+err.Error())
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
