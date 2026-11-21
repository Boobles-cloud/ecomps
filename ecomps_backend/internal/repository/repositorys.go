package repository

import (
	"context"
	"uuid"

	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
)

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetAllByTenantId(ctx context.Context, tenantId uuid.UUID) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint, tenantId uuid.UUID) error
}

type ProductPictureRepository interface {
	GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error)
	GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error)
	GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error)
	Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error)
	Delete(ctx context.Context, id uint) error
	DeleteByProductId(ctx context.Context, productId uint) error
}
