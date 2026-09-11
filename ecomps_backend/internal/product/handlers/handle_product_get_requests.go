package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Gets a Product by the id
func (p *ProductHandler) HandleGettingProductById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleGettingProductById")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	key := ProductCacheKey + strconv.Itoa(productId)
	cacheItem, ok := p.productCache.GetItem(key)

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	product, err := p.productService.GetById(ctx, uint(productId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go p.insertItem(product)

	if !jsonutils.RespondWithJson(w, http.StatusOK, product) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Gets all products by the given tenant id
func (p *ProductHandler) HandleGettingAllProductsByTenantId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleGettingAllProductByTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := p.productCache.GetItems(uint(tenantId))

	if ok || len(cacheItems) != 0 {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	allProducts, err := p.productService.GetAllByTenantId(ctx, uint(tenantId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	// Write the stuff to the cache
	go p.insertItems(allProducts)

	if !jsonutils.RespondWithJson(w, http.StatusOK, allProducts) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
