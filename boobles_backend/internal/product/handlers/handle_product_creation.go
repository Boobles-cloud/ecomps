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

// Handels creating a new product
func (p *ProductHandler) HandleCreatingProduct(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Product | HandleCreatingProduct] "+err.Error())
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

	result := database.ExecuteSQLStatement("InsertProduct", database.Insert, []any{product.ProductName, product.ProductPrice,
		product.ProductDescription, product.ProductPicturePath, product.TenantId})

	if !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to create product"))
	}

	product.ProductId = result.LastId

	go p.insertItem(product)
	w.WriteHeader(http.StatusOK)
}

// TODO: we want to accept more then one picture here
func (p *ProductHandler) HandleCreatingProductPicture(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
	w.WriteHeader(http.StatusBadRequest)
}
