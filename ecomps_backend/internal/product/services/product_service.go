package services

import (
	"context"
	"errors"

	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	"ecomps.boobles.cloud/backend/utils/crypto"
)

func (p *ProductService) GetById(ctx context.Context, id uint) (productstructs.Product, error) {

	if id == 0 {
		return productstructs.Product{}, errors.New("Product Id cant be 0")
	}

	encryptedProduct, err := p.productRepo.GetById(ctx, id)

	if err != nil {
		return encryptedProduct, err
	}

	pw, err := p.tenantService.GetPw(ctx, encryptedProduct.TenantId)

	if err != nil {
		return encryptedProduct, err
	}

	product, ok := crypto.Decrypt[productstructs.Product](&encryptedProduct, pw)

	if !ok {
		return encryptedProduct, errors.New("Failed to encrypt")
	}

	return *product, nil
}

func (p *ProductService) GetAllByTenantId(ctx context.Context, tenantId uint) ([]productstructs.Product, error) {

	if tenantId == 0 {
		return []productstructs.Product{}, errors.New("Tenant Id cant be 0")
	}

	encryptedProduct, err := p.productRepo.GetAllByTenantId(ctx, tenantId)

	if err != nil {
		return encryptedProduct, err
	}

	pw, err := p.tenantService.GetPw(ctx, encryptedProduct[0].TenantId)

	if err != nil {
		return encryptedProduct, err
	}

	products := make([]productstructs.Product, 0, len(encryptedProduct))

	for i := range encryptedProduct {

		product, ok := crypto.Decrypt[productstructs.Product](&products[i], pw)

		if ok {
			products = append(products, *product)
		}
	}

	return products, nil
}

func (p *ProductService) Create(ctx context.Context, item productstructs.Product) (uint, error) {

	if item.TenantId == 0 {
		return 0, errors.New("Tenant Id cant be 0")
	}

	if item.ProductName == "" {
		return 0, errors.New("Product name is emtpy")
	}

	pw, err := p.tenantService.GetPw(ctx, item.TenantId)

	if err != nil {
		return 0, err
	}

	encrypted, ok := crypto.Encrypt[productstructs.Product](item, pw)

	if !ok {
		return 0, errors.New("Failed to encrypt")
	}

	return p.productRepo.Create(ctx, encrypted)
}

func (p *ProductService) Update(ctx context.Context, item productstructs.Product, filterName string) error {

	pw, err := p.tenantService.GetPw(ctx, item.TenantId)

	if err != nil {
		return err
	}

	encrypted, ok := crypto.Encrypt[productstructs.Product](item, pw)

	if !ok {
		return errors.New("Failed to encrypt")
	}

	return p.productRepo.Update(ctx, encrypted, "ProductId")
}

func (p *ProductService) Delete(ctx context.Context, id uint, tenantId uint) error {

	if id == 0 {
		return errors.New("Product Id cant be 0")
	}

	return p.productRepo.Delete(ctx, id, tenantId)
}
