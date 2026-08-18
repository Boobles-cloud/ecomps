package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
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
		return
	}

	var user userstructs.UserStruct
	if err := json.Unmarshal(body, &user); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if !user.UpdateUserInDB(u.Dh) {
		fail(http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
