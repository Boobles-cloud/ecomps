package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// Handels the adding of a new permission for the user
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it
func HandleAddingNewUserPermission(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleAddingNewUserPermission] "+err.Error())
		}

		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	givenUser := new(userstructs.UserStruct)

	if err := json.Unmarshal(body, &givenUser); err != nil {
		fail(http.StatusInternalServerError, err)
	}

	// First lets get all permissions and check if the user has them allready
	user, ok := database.QueryDatabase[userstructs.UserStruct]("SelectUserById", []any{givenUser.UserId})

	if !ok && len(user) != 1 {
		fail(http.StatusInternalServerError, errors.New("Got more then one user from db..."))
	}

}

// Handels the removing of a user permission
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it.
func HandleRemovingUserPermission(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleRemovingNewUserPermission] "+err.Error())
		}

		w.WriteHeader(status)
	}

	fail(http.StatusBadRequest, errors.New("NOT IMPLEMENTED"))
}
