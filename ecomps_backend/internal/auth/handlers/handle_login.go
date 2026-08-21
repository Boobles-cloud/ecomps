package handlers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Creates a new JWT for the given user.
// The pw from the user, is encrypted via the frontend.
func (ha *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Auth | HandleLogin")

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

	cookie, err := httputils.CreateAuthCookie(userFromDB.UserId, userFromDB.TenantId, ha.Dh)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
