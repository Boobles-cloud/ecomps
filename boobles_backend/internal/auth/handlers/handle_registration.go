package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// TODO: implement func to check if user is already signed up!!!

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
	}

	var tmpUserStruct userstructs.UserStruct

	// Get all content from the body
	if err := json.Unmarshal(body, &tmpUserStruct); err != nil {
		fail(http.StatusBadRequest, err)
	}

	// Creates the user in the database
	ok, id := tmpUserStruct.CreateUserInDB(ha.Dh)

	if !ok {
		fail(http.StatusInternalServerError, nil)
	}

	// Sets the id for a user
	tmpUserStruct.UserId = id

	// Creates a token for the user
	token, ok := createJWT(tmpUserStruct, ha.Dh)

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
