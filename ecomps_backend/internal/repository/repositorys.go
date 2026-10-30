package repository

import (
	"context"
	"uuid"

	authstructs "ecomps.boobles.cloud/backend/internal/auth/auth_structs"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	permissionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetAllByTenantId(ctx context.Context, tenantId uuid.UUID) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint, tenantId uuid.UUID) error
}

type TenantRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (tenantstructs.Tenant, error)
	GetTenantDeletion(ctx context.Context, id uuid.UUID) (tenantstructs.TenantDeletionStruct, error)
	Create(ctx context.Context, tenant *tenantstructs.Tenant, userId uint) error
	Update(ctx context.Context, tenant tenantstructs.Tenant) error
	Delete(ctx context.Context, userId uint, tenantId uuid.UUID) error
}

// Just for all application that we support
type OAuthAppRepository interface {
	GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error)
}

// All settings a user can chose from
type OAuthSettingsRepository interface {
	GetAllSettings(ctx context.Context, langId uint) ([]tenantoauthstructs.OAuthSettings, error)
}

// All apps a tenant has active
type OAuthTenantAppRepository interface {
	GetAllOAuthAppsForTenant(ctx context.Context, tenantId uuid.UUID) ([]tenantoauthstructs.TenantOAuthApps, error)
	AddOAuthApplicationForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error
	DeleteApplicationForTenant(ctx context.Context, tenantId uuid.UUID, appId uint) error
}

// All settings for the tenant apps
type OAuthTenantSettingsRepository interface {
	GetSettingsForApplication(ctx context.Context, tenantId uuid.UUID, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error)
	AddSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error
	ChangeSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error
}

type OAuthTenantTokenRepository interface {
	GetTokenForTenantApp(ctx context.Context, tenantId uuid.UUID, appId uint) (tenantoauthstructs.TenantOAuthToken, error)
	AddToken(ctx context.Context, token tenantoauthstructs.TenantOAuthToken) error
	DeleteToken(ctx context.Context, tokenId uint) error
}

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

type ProductPictureRepository interface {
	GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error)
	GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error)
	GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error)
	Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error)
	Delete(ctx context.Context, id uint) error
	DeleteByProductId(ctx context.Context, productId uint) error
}

type AuthRepository interface {
	CreateAccessTokenInDatabase(ctx context.Context, item authstructs.JWTDatabaseStruct) error
	DeleteAccessTokenByValue(ctx context.Context, cookieVal string) error
}

type OrderRepository interface {
	GetOrderById(ctx context.Context, id uint) (orderstructs.Order, error)
	GetAllOrdersByTenantId(ctx context.Context, tenantId uuid.UUID) ([]orderstructs.Order, error)
	CreateOrder(ctx context.Context, order orderstructs.Order) (uint, error)
	UpdateOrder(ctx context.Context, order orderstructs.Order) error
	DeleteOrder(ctx context.Context, orderId uint) error
}

type OrderProductRepository interface {
	GetAllProductsByOrderId(ctx context.Context, orderId uint) ([]orderstructs.OrderProduct, error)
	CreateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error
	UpdateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error
	DeleteOrderProductsByOrderId(ctx context.Context, orderId uint) error
	DeleteOrderProduct(ctx context.Context, productId uint) error
}

type OrderStatusRepostiory interface {
	GetById(ctx context.Context, id, langId uint) (orderstructs.OrderStatus, error)
	GetAllByLangId(ctx context.Context, langId uint) ([]orderstructs.OrderStatus, error)
}
