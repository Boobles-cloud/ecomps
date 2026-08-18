package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"ecomps.boobles.cloud/backend/database"
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/logging"
	"github.com/golang-jwt/jwt/v4"
)

// Creates a new JWT for the given user.
// The pw from the user, is encrypted via the frontend.
func (ha *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Auth | HandleLogin] "+err.Error())
		}

		w.WriteHeader(status)
	}

	basicAuth := r.Header.Get("Authorization")

	if basicAuth == "" {
		fail(http.StatusBadRequest, nil)
		return
	}

	authPWUser := strings.ReplaceAll(basicAuth, "Basic ", "")

	encodedAuthPWUser, err := base64.StdEncoding.DecodeString(authPWUser)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	authSplitet := strings.Split(string(encodedAuthPWUser), ":")

	if len(authSplitet) != 2 {
		fail(http.StatusBadRequest, errors.New("Failed spliting basic pw"))
		return
	}

	userFromDB, ok := database.QueryOne[userstructs.UserStruct](r.Context(), ha.Dh, "SelectUserByUserNameAndPW", authSplitet[0], authSplitet[1])

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, ok := createJWT(userFromDB, ha.Dh)

	if !ok {
		fail(http.StatusInternalServerError, nil)
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
func createJWT(user userstructs.UserStruct, dh *database.DbHandler) (string, bool) {

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
	if !tokenDB.InsertIntoDB(dh) {
		return "", false
	}

	if err != nil {
		logging.Log(logging.Error, "[AuthHandler | HandleLogin]"+err.Error())
		return "", false
	}

	return tokenSigned, true
}
