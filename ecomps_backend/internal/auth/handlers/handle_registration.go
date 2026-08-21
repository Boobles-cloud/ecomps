package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the registration of a user.
// Sends back an access token.
func (ha *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Auth | HandleRegistration")

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
		fail(http.StatusInternalServerError, errors.New("Failed creating user in database"))
		return
	}

	// Sets the id for a user
	tmpUserStruct.UserId = id

	// Creates a token for the user
	cookie, err := httputils.CreateAuthCookie(tmpUserStruct.UserId, tmpUserStruct.TenantId, ha.Dh)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
