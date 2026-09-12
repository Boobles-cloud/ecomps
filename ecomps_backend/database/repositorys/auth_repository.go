package repositorys

import (
	"context"
	"errors"

	"ecomps.boobles.cloud/backend/database"
	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
)

type AuthRepository struct {
	*DatabaseRepository[authstructs.JWTDatabaseStruct]
}

func NewAuthRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[authstructs.JWTDatabaseStruct]) *AuthRepository {
	return &AuthRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (a *AuthRepository) CreateAccessTokenInDatabase(ctx context.Context, item authstructs.JWTDatabaseStruct) error {

	if result := a.db.ExecuteSQLStatement(ctx, "InsertUserAccessToken", a.toArgs(item)); !result.Ok {
		return errors.New("Failed to insert access token")
	}
	return nil
}

func (a *AuthRepository) DeleteAccessTokenByValue(ctx context.Context, cookieVal string) error {

	if result := a.db.ExecuteSQLStatement(ctx, "DeleteUserAccestokenByValue", []any{cookieVal}); !result.Ok {
		return errors.New("Failed to delete Access token")
	}

	return nil
}
