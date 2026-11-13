package repository

import (
	"context"
	"uuid"

	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetAllByTenantId(ctx context.Context, tenantId uuid.UUID) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint, tenantId uuid.UUID) error
}

// Just for all application that we support
type OAuthAppRepository interface {
	GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error)
}

// All settings a user can chose from
type OAuthSettingsRepository interface {
	GetAllSettings(ctx context.Context, langId uint) ([]tenantoauthstructs.OAuthSettings, error)
}

type ProductPictureRepository interface {
	GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error)
	GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error)
	GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error)
	Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error)
	Delete(ctx context.Context, id uint) error
	DeleteByProductId(ctx context.Context, productId uint) error
}
