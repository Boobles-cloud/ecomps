package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/crypto"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels updating a product
func (p *ProductHandler) HandleChangingProduct(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleChangingProduct")

	product, err := jsonutils.JsonDeserilizeHttpRequestBody[productstructs.Product](r)

	if err != nil {
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
