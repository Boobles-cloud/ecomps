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

// Handels updating a product
func (p *ProductHandler) HandleChangingProduct(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Product | HandleChangingProduct] "+err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var product productstructs.Product

	if err := json.Unmarshal(body, &product); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Get the tenant id
	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	// Get the tenant for encryption stuff
	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), p.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting Tenant"))
		return
	}

	// Encrypt the product
	encryptedProduct, ok := crypto.Encrypt[productstructs.Product](product, tenant.GetPw(p.Dh, r.Context()))

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to encrypt product"))
		return
	}

	if !database.UpdateDatabaseEntry[productstructs.Product](p.Dh, "UpdateProduct", "ProductId", encryptedProduct) {
		fail(http.StatusInternalServerError, errors.New("Failed to update product"))
		return
	}

	go p.insertItem(product)
	w.WriteHeader(http.StatusOK)
}
