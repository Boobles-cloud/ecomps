package repositorys

import (
	"context"
	"database/sql"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type UserRepository struct {
	*DatabaseRepository[userstructs.UserStruct]
}

// Creates a new user Repository
func NewUserRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[userstructs.UserStruct]) *UserRepository {
	return &UserRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

// Gets a user by the given email
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (userstructs.UserStruct, error) {

	user, ok := database.QueryOne[userstructs.UserStruct](ctx, r.db, "Select"+r.entityName+"ByEmail", email)

	if !ok {
		return user, sql.ErrNoRows
	}

	return user, nil
}
