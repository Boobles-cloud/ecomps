package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// Handels the registration of a user.
// Sends back an access token.
func (ha *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Auth | HandleRegistration] "+err.Error())
		}

		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var tmpUserStruct userstructs.UserStruct

	// Get all content from the body
	if err := json.Unmarshal(body, &tmpUserStruct); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if _, ok := database.QueryOne[userstructs.UserStruct](r.Context(), ha.Dh, "SelectUserByEmail", tmpUserStruct.UserMail); ok {
		fail(http.StatusConflict, errors.New("User already exists"))
		return
	}

	// Creates the user in the database
	ok, id := tmpUserStruct.CreateUserInDB(ha.Dh)

	if !ok {
		fail(http.StatusInternalServerError, nil)
		return
	}

	// Sets the id for a user
	tmpUserStruct.UserId = id

	// Creates a token for the user
	token, ok := createJWT(tmpUserStruct, ha.Dh)

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
