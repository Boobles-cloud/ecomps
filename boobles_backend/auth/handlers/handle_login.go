package handlers

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"time"

	authstructs "boobles.cloud/backend/auth/auth_structs"
	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
	userstructs "boobles.cloud/backend/user_structs"
	"github.com/golang-jwt/jwt/v4"
)

// Creates a new JWT for the given user.
// The pw from the user, is encrypted via the frontend.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	basicAuth := r.Header.Get("Authorization")

	if basicAuth == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authPWUser := strings.ReplaceAll(basicAuth, "Basic ", "")

	encodedAuthPWUser, err := base64.StdEncoding.DecodeString(authPWUser)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authSplitet := strings.Split(string(encodedAuthPWUser), ":")

	userFromDB, ok := database.QueryDatabase[userstructs.UserStruct]("SelectUserByUserNameAndPW", []any{authSplitet[0], authSplitet[1]})

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if len(userFromDB) != 1 {
		logging.Log(logging.Error, "Multiple users from database. [login handler]")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, ok := createJWT(userFromDB[0])

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "AuthTokenBoobles",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 3),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

// Creates the jwt
func createJWT(user userstructs.UserStruct) (string, bool) {

	claims := authstructs.JWTClaimsStruct{
		UserId:   user.UserId,
		TenantId: user.TenantId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Boobles_backend_server",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(os.Getenv("JWT-Secret")))

	// Insert our token into the database
	tokenDB := authstructs.JWTDatabaseStruct{
		UserAccessId: 0,
		TokenVal:     tokenSigned,
		TokenExpire:  time.Now().AddDate(0, 0, 3),
		UserId:       user.UserId,
	}

	// Checking if ok
	if !tokenDB.InsertIntoDB() {
		return "", false
	}

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return "", false
	}

	return tokenSigned, true
}
