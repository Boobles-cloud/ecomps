package hostedservices

import (
	"context"
	"time"

	"ecomps.boobles.cloud/backend/database"
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
)

// Deletes all expired jwt.
// This runs once a day
func DeleteExpiredJWT(ctx context.Context, dh *database.DbHandler) {

	timer := time.NewTicker(time.Hour * 24)

	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			allTokens, ok := database.QueryMany[authstructs.JWTDatabaseStruct](ctx, dh, "SelectAllToken", []any{})

			if !ok {
				timer.Reset(time.Hour * 24)
			}

			for i := range allTokens {
				if allTokens[i].IsExpired() {
					dh.ExecuteSQLStatement("DeleteUserAccestoken", []any{allTokens[i].UserAccessId})
				}
			}

			timer.Reset(time.Hour * 24)
		}
	}
}
