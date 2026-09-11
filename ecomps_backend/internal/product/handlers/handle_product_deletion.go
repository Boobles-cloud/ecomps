package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the deletion of a prodcut
func (p *ProductHandler) HandleDeletingProduct(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	// TODO: also delete all product pictures

	fail := httputils.NewFailHandler(w, "Product | HandleDeletingProduct")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	p.productService.Delete(ctx, uint(productId), ctx.Value(middleware.TenantIdContextKey).(uint))

	key := ProductCacheKey + strconv.Itoa(productId)
	p.productCache.RemoveItem(key)

	w.WriteHeader(http.StatusOK)
}
