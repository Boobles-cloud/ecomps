package services

import (
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/tenant/services"
)

type ProductService struct {
	productRepo   repository.Repository[productstructs.Product]
	tenantService *services.TenantService
}

func CreateNewProductService(r repository.Repository[productstructs.Product], t *services.TenantService) *ProductService {
	return &ProductService{
		productRepo:   r,
		tenantService: t,
	}
}

func ProductToArgs(p productstructs.Product) []any {
	return []any{
		p.ProductId,
		p.ProductName,
		p.ProductPrice,
		p.ProductDescription,
		p.TenantId,
	}
}
