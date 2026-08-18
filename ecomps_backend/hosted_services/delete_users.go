package hostedservices

import (
	"context"
	"time"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Checks every 12h if a user needs a deletion or not
func DeleteUsers(ctx context.Context, dh *database.DbHandler) {

	ticker := time.NewTicker(12 * time.Hour)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			allUserDeletion, ok := database.QueryMany[userstructs.UserDeletionDatabase](ctx, dh, "SelectAllUserDeletions", []any{})

			if !ok {
				logging.Log(logging.Error, "[Hosted Services | DeleteUsers] Failed to get user or there are none")
				ticker.Reset(12 * time.Hour)
			}

			for i := range allUserDeletion {
				if allUserDeletion[i].IsDeletionToday() {
					dh.ExecuteSQLStatement("DeleteUserById", []any{allUserDeletion[i].UserId})
					dh.ExecuteSQLStatement("DeleteUserDeletionById", []any{allUserDeletion[i].DeletionId})
				}
			}

			ticker.Reset(12 * time.Hour)
		}
	}
}
