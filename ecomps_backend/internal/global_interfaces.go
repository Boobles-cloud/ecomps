package internal

import "context"

type Repository[T any] interface {
	GetById(ctx context.Context, id uint) (T, error)
	GetAllByTenantId(ctx context.Context, tenantId uint) ([]T, error)
	Create(ctx context.Context, item T) (uint, error)
	Update(ctx context.Context, item T, filterName string) error
	Delete(ctx context.Context, id uint) error
}
