package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// Handels the user change stuff
func (u *UserHandler) HandleUserChange(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[User | HandleUserChange] "+err.Error())
		}
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	var user userstructs.UserStruct
	if err := json.Unmarshal(body, &user); err != nil {
		fail(http.StatusBadRequest, err)
	}

	if !user.UpdateUserInDB(u.Dh) {
		fail(http.StatusInternalServerError, nil)
	}

	w.WriteHeader(http.StatusOK)
}
