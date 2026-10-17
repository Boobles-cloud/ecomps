package services

import (
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	pictureservice "ecomps.boobles.cloud/backend/internal/product_pictures/services"
	"ecomps.boobles.cloud/backend/internal/repository"
	tenantservice "ecomps.boobles.cloud/backend/internal/tenant/services"
)

type ProductService struct {
	productRepo    repository.Repository[productstructs.Product]
	tenantService  *tenantservice.TenantService
	pictureService *pictureservice.ProductPictureService
}

func CreateNewProductService(r repository.Repository[productstructs.Product], t *tenantservice.TenantService, p *pictureservice.ProductPictureService) *ProductService {
	return &ProductService{
		productRepo:    r,
		tenantService:  t,
		pictureService: p,
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
