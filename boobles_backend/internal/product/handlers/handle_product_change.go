package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"boobles.cloud/backend/database"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
	"boobles.cloud/backend/logging"
)

// Handels updating a product
func (p *ProductHandler) HandleChangingProduct(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Product | HandleChangingProduct] "+err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	var product productstructs.Product

	if err := json.Unmarshal(body, &product); err != nil {
		fail(http.StatusBadRequest, err)
	}

	if !database.UpdateDatabaseEntry[productstructs.Product]("UpdateProduct", "ProductId", product) {
		fail(http.StatusInternalServerError, errors.New("Failed to update product"))
	}

	go p.insertItem(product)
	w.WriteHeader(http.StatusOK)
}
