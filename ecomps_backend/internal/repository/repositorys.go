package repository

import (
	"context"

	permssionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetTenant(ctx context.Context, tenantId uint) (tenantstructs.Tenant, error)
	GetAllByTenantId(ctx context.Context, tenantId uint) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint, tenantId uint) error
}

type TenantRepository interface {
	GetById(ctx context.Context, id uint) (tenantstructs.Tenant, error)
	GetPw(ctx context.Context, tenantId uint) (string, error)
	Create(ctx context.Context, tenant tenantstructs.Tenant, userId uint) error
	Update(ctx context.Context, tenant tenantstructs.Tenant) error
	Delete(ctx context.Context, userId, tenantId uint) error
}

type UserRepository interface {
	Repository[userstructs.UserStruct]
	GetUserByEmail(ctx context.Context, email string) (userstructs.UserStruct, error)
	CreateUserDeletionDate(ctx context.Context, userId uint) error
}

type PermissionRepository interface {
	Repository[permssionstructs.Permission]
	AsingUserPermission(ctx context.Context, userId, permissionId uint) error
	GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]permssionstructs.Permission, error)
	GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]permssionstructs.Permission, error)
	RemoveUserPermission(ctx context.Context, userId, permissionId uint) error
}

type ProductPictureRepository interface {
	GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error)
	GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error)
	GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error)
	Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error)
	Delete(ctx context.Context, id uint) error
}
