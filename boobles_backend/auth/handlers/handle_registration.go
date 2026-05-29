package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"boobles.cloud/backend/logging"
	userstructs "boobles.cloud/backend/user_structs"
)

// Handels the registration of a user.
// Sends back an access token.
func HandleRegistration(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tmpUserStruct userstructs.UserStruct

	// Get all content from the body
	if err := json.Unmarshal(body, &tmpUserStruct); err != nil {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Creates the user in the database
	ok, id := tmpUserStruct.CreateUserInDB()

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Sets the id for a user
	tmpUserStruct.UserId = id

	// Creates a token for the user
	token, ok := createJWT(tmpUserStruct)

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
