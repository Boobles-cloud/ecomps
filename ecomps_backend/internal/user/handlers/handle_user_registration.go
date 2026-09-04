package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the registration of a user.
// Sends back an access token.
func (hu *UserHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleRegistration")

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

	id, err := hu.userService.CreateUser(ctx, tmpUserStruct)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	// Sets the id for a user
	tmpUserStruct.UserId = id

	// Creates a token for the user
	cookie, err := httputils.CreateAuthCookie(tmpUserStruct.UserId, tmpUserStruct.TenantId, hu.Dh)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
