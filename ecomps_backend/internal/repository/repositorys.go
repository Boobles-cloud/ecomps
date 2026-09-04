package repository

import (
	"context"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetAllByTenantId(ctx context.Context, tenantId uint) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint) error
}

type UserRepository interface {
	Repository[userstructs.UserStruct]
	GetUserByEmail(ctx context.Context, email string) (userstructs.UserStruct, error)
}

type UserPermissionRepository interface {
	Repository[userstructs.UserPermission]
	GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]userstructs.UserPermission, error)
	GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]userstructs.UserPermission, error)
}
