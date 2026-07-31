package handlers

import (
	"boobles.cloud/backend/caching"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
)

type ProductHandler struct {
	ProductCache *caching.CacheManager[productstructs.Product]
}

// Creates a new handler for products
func CreateNewHandler(c *caching.CacheManager[productstructs.Product]) *ProductHandler {
	return &ProductHandler{
		ProductCache: c,
	}
}
