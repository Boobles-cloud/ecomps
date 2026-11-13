package repository

import (
	"context"

	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	permissionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type UserRepository interface {
	Repository[userstructs.UserStruct]
	GetUserByUserNameAndPw(ctx context.Context, userName, pw string) (userstructs.UserStruct, error)
	GetUserByEmail(ctx context.Context, email string) (userstructs.UserStruct, error)
	CreateUserDeletionDate(ctx context.Context, userId uint) error
}

type PermissionRepository interface {
	Repository[permissionstructs.Permission]
	AsingUserPermission(ctx context.Context, userId, permissionId uint) error
	GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]permissionstructs.Permission, error)
	GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]permissionstructs.Permission, error)
	RemoveUserPermission(ctx context.Context, userId, permissionId uint) error
}

type AuthRepository interface {
	CreateAccessTokenInDatabase(ctx context.Context, item authstructs.JWTDatabaseStruct) error
	DeleteAccessTokenByValue(ctx context.Context, cookieVal string) error
}
