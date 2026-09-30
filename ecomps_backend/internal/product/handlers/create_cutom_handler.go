package handlers

import (
	"strconv"

	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	"ecomps.boobles.cloud/backend/internal/product/services"
	"ecomps.boobles.cloud/backend/utils/caching"
)

const (
	ProductCacheKey = "PRODUCT:"
)

type ProductHandler struct {
	productCache   *caching.CacheManager[productstructs.Product]
	productService *services.ProductService
}

// Creates a new handler for products
func CreateNewProductHandler(c *caching.CacheManager[productstructs.Product], p *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productCache:   c,
		productService: p,
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

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *ProductHandler) insertItems(t []productstructs.Product) {

	for i := range t {
		key := ProductCacheKey + strconv.Itoa(int(t[i].ProductId))
		p.productCache.SetOrUpdateItem(key, t[i], t[i].TenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *ProductHandler) insertItem(t productstructs.Product) {
	key := ProductCacheKey + strconv.Itoa(int(t.ProductId))
	p.productCache.SetOrUpdateItem(key, t, t.TenantId)
}
