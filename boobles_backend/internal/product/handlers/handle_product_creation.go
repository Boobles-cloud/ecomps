package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	productstructs "boobles.cloud/backend/internal/product/product_structs"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
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

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantById", []any{tenantId})

	if !ok || len(tenant) != 1 {
		fail(http.StatusInternalServerError, errors.New("Failed getting Tenant"))
	}

	encryptedProduct, ok := crypto.Encrypt[productstructs.Product](product, tenant[0].GetPw())

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to encrypt product"))
	}

	result := database.ExecuteSQLStatement("InsertProduct", database.Insert, []any{encryptedProduct.ProductName, encryptedProduct.ProductPrice,
		encryptedProduct.ProductDescription, encryptedProduct.ProductPicturePath, encryptedProduct.TenantId})

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
