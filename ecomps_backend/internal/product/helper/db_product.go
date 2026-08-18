package helper

import (
	"context"

	"ecomps.boobles.cloud/backend/crypto"
	"ecomps.boobles.cloud/backend/database"
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Gets a product by its id
func GetProductById(productId uint, key string, ctx context.Context, dh *database.DbHandler) (*productstructs.Product, bool) {

	p, ok := database.QueryOne[productstructs.Product](ctx, dh, "SelectProductById", productId)

	if !ok {
		logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to get product from database")
		return nil, false
	}

	decrypted, ok := crypto.Decrypt[productstructs.Product](&p, key)

	if !ok {
		logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to decrypt product")
		return nil, false
	}

	return decrypted, true
}

// Gets all products for a tenant
func GetAllProductsForTenant(tenantId uint, key string, ctx context.Context, dh *database.DbHandler) ([]productstructs.Product, bool) {

	products, ok := database.QueryMany[productstructs.Product](ctx, dh, "SelectProductByTenantId", tenantId)

	if !ok {
		logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to get products for tenant")
		return []productstructs.Product{}, false
	}

	decryptedProducts := make([]productstructs.Product, len(products))

	for i := range products {

		p, ok := crypto.Decrypt[productstructs.Product](&products[i], key)

		if !ok {
			logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to decrypt product")
			continue
		}

		decryptedProducts = append(decryptedProducts, *p)
	}

	return decryptedProducts, true
}
