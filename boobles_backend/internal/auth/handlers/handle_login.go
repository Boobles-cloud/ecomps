package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"boobles.cloud/backend/database"
	authstructs "boobles.cloud/backend/internal/auth/auth_structs"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
	"github.com/golang-jwt/jwt/v4"
)

// Creates a new JWT for the given user.
// The pw from the user, is encrypted via the frontend.
func HandleLogin(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Auth | HandleLogin] "+err.Error())
		}

		w.WriteHeader(status)
	}

	basicAuth := r.Header.Get("Authorization")

	if basicAuth == "" {
		fail(http.StatusBadRequest, nil)
	}

	authPWUser := strings.ReplaceAll(basicAuth, "Basic ", "")

	encodedAuthPWUser, err := base64.StdEncoding.DecodeString(authPWUser)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	authSplitet := strings.Split(string(encodedAuthPWUser), ":")

	userFromDB, ok := database.QueryDatabase[userstructs.UserStruct]("SelectUserByUserNameAndPW", []any{authSplitet[0], authSplitet[1]})

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if len(userFromDB) != 1 {
		fail(http.StatusInternalServerError, errors.New("More than one user!"))
	}

	token, ok := createJWT(userFromDB[0])

	if !ok {
		fail(http.StatusInternalServerError, nil)
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
