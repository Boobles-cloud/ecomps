package repositorys

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

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

func (r *UserRepository) GetUserByUserNameAndPw(ctx context.Context, userName, pw string) (userstructs.UserStruct, error) {

	user, ok := database.QueryOne[userstructs.UserStruct](ctx, r.db, "SelectUserByUserNameAndPW", []any{userName, pw})

	if !ok {
		return user, sql.ErrNoRows
	}

	return user, nil
}

// Creates a user deletion date in the database
// Our background worker will then get those and delete them
func (r *UserRepository) CreateUserDeletionDate(ctx context.Context, userId uint) error {

	result := r.db.ExecuteSQLStatement(ctx, "InsertUserDeletion", []any{time.Now(), time.Now().AddDate(0, 1, 0), userId})

	if !result.Ok {
		return errors.New("Failed to insert user deletion for userId:" + strconv.Itoa(int(userId)))
	}
	return nil
}
