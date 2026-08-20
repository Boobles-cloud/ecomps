package handlers

import (
	"errors"
	"net/http"
	"strconv"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the deletion of a prodcut
func (p *ProductHandler) HandleDeletingProduct(w http.ResponseWriter, r *http.Request) {

	// TODO: also delete all product pictures

	fail := httputils.NewFailHandler(w, "Product | HandleDeletingProduct")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if result := p.Dh.ExecuteSQLStatement("DeleteProductById", []any{productId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed deleting product"))
		return
	}

	key := ProductCacheKey + strconv.Itoa(productId)
	p.ProductCache.RemoveItem(key)

	w.WriteHeader(http.StatusOK)
}
