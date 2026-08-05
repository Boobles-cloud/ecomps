package handlers

import (
	"strconv"

	"boobles.cloud/backend/caching"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
)

const (
	ProductCacheKey = "PRODUCT:"
)

type ProductHandler struct {
	ProductCache *caching.CacheManager[productstructs.Product]
}

// Creates a new handler for products
func CreateNewProductHandler(c *caching.CacheManager[productstructs.Product]) *ProductHandler {
	return &ProductHandler{
		ProductCache: c,
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *ProductHandler) insertItems(t []productstructs.Product) {

	for i := range t {

		key := ProductCacheKey + strconv.Itoa(int(t[i].ProductId))
		p.ProductCache.SetOrUpdateItem(key, t[i], t[i].TenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *ProductHandler) insertItem(t productstructs.Product) {
	key := ProductCacheKey + strconv.Itoa(int(t.ProductId))
	p.ProductCache.SetOrUpdateItem(key, t, t.TenantId)
}
