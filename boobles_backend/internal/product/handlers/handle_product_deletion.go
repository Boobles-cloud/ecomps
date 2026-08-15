package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/logging"
)

// Handels the deletion of a prodcut
func (p *ProductHandler) HandleDeletingProduct(w http.ResponseWriter, r *http.Request) {

	// TODO: also delete all product pictures

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	productId, err := strconv.Atoi(r.PathValue("product_id"))

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
