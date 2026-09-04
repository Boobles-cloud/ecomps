package repositorys

import (
	"context"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type UserPermissionRepository struct {
	*DatabaseRepository[userstructs.UserPermission]
}

// Creates a new user permission repository
func NewUserPermissionRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[userstructs.UserPermission]) *UserPermissionRepository {
	return &UserPermissionRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (u *UserPermissionRepository) GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]userstructs.UserPermission, error)

func (u *UserPermissionRepository) GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]userstructs.UserPermission, error)
