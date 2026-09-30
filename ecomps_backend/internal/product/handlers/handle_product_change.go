package handlers

import (
	"context"
	"net/http"
	"time"

	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels updating a product
func (p *ProductHandler) HandleChangingProduct(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleChangingProduct")

	product, err := jsonutils.JsonDeserilizeHttpRequestBody[productstructs.Product](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := p.productService.Update(ctx, product, "ProductId"); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go p.insertItem(product)
	w.WriteHeader(http.StatusOK)
}
