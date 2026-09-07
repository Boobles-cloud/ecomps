package handlers

import (
	"context"
	"net/http"
	"time"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the user change stuff
func (hu *UserHandler) HandleUserChange(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleUserChange")

	user, err := jsonutils.JsonDeserilizeHttpRequestBody[userstructs.UserStruct](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := hu.userService.UpdateUser(ctx, user, "UserId"); err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusOK)
}
