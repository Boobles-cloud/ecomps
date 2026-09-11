package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels creating a new product
func (p *ProductHandler) HandleCreatingProduct(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleCreatingProduct")

	product, err := jsonutils.JsonDeserilizeHttpRequestBody[productstructs.Product](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := uint(ctx.Value(middleware.TenantIdContextKey).(int))

	product.TenantId = tenantId

	pId, err := p.productService.Create(ctx, product)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	product.ProductId = pId

	go p.insertItem(product)
	w.WriteHeader(http.StatusOK)
}
