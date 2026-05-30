package auth

import (
	"context"
	"time"

	authstructs "boobles.cloud/backend/auth/auth_structs"
	"boobles.cloud/backend/database"
)

// Deletes all expired jwt.
// This runs once a day
func DeleteExpiredJWT(ctx context.Context) {

	timer := time.NewTicker(time.Hour * 24)

	defer timer.Stop()

deleteLoop:
	for {
		select {
		case <-ctx.Done():
			break deleteLoop
		case <-timer.C:
			allTokens, ok := database.QueryDatabase[authstructs.JWTDatabaseStruct]("SelectAllToken", []any{})

			if !ok {
				timer.Reset(time.Hour * 24)
			}

			for i := range allTokens {
				if allTokens[i].IsExpired() {
					database.ExecuteSQLStatement("DeleteUserAccestoken", database.Delete, []any{allTokens[i].UserAccessId})
				}
			}

			timer.Reset(time.Hour * 24)
		}
	}
}
