package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	productstructs "ecomps.boobles.cloud/backend/internal/product/product_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels creating a new product
func (p *ProductHandler) HandleCreatingProduct(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleCreatingProduct")

	product, err := jsonutils.JsonDeserilizeHttpRequestBody[productstructs.Product](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), p.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting Tenant"))
		return
	}

	// We create a copy here so we dont set the ecrypted stuff into the cache
	copyOfProduct := product

	id, ok := product.CreateProductInDatabase(tenant.GetPw(p.Dh, r.Context()), p.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to create product"))
		return
	}

	copyOfProduct.ProductId = id

	go p.insertItem(copyOfProduct)
	w.WriteHeader(http.StatusOK)
}

// TODO: we want to accept more then one picture here
func (p *ProductHandler) HandleCreatingProductPicture(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
	w.WriteHeader(http.StatusBadRequest)
}
