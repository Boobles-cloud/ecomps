package handlers

import (
	"net/http"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the user change stuff
func (u *UserHandler) HandleUserChange(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HandleUserChange")

	user, err := jsonutils.JsonDeserilizeHttpRequestBody[userstructs.UserStruct](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if !user.UpdateUserInDB(u.Dh) {
		fail(http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
