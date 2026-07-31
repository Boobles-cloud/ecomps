package helper

import (
	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
	"boobles.cloud/backend/logging"
)

// Gets a product by its id
func GetProductById(productId uint, key string) (*productstructs.Product, bool) {

	p, ok := database.QueryDatabase[productstructs.Product]("SelectProductById", []any{productId})

	if !ok || len(p) != 1 {
		logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to get product from database")
		return nil, false
	}

	decrypted, ok := crypto.Decrypt[productstructs.Product](&p[0], key)

	if !ok {
		logging.Log(logging.Error, "[Product Helper | GetProductById] Failed to decrypt product")
		return nil, false
	}

	return decrypted, true
}

// Gets all products for a tenant
func GetAllProductsForTenant(tenantId uint, key string) ([]productstructs.Product, bool) {

	products, ok := database.QueryDatabase[productstructs.Product]("SelectProductByTenantId", []any{tenantId})

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
