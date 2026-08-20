package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	"ecomps.boobles.cloud/backend/internal/product/helper"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Gets a Product by the id
func (p *ProductHandler) HandleGettingProductById(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleGettingProductById")

	id, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey)

	key := ProductCacheKey + strconv.Itoa(id)
	cacheItem, ok := p.ProductCache.GetItem(key)

	if ok {

		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), p.Dh, "SelectTenantById", tenantId.(int))

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	product, ok := helper.GetProductById(uint(id), tenant.GetPw(p.Dh, r.Context()), r.Context(), p.Dh)

	go p.insertItem(*product)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting Product"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, product) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Gets all products by the given tenant id
func (p *ProductHandler) HandleGettingAllProductsByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleGettingAllProductByTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := p.ProductCache.GetItems(uint(tenantId))

	if ok || len(cacheItems) != 0 {

		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	// We always need to get the tenant, because all items are encrypted in the database
	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), p.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	// Gets all decrypted products
	allProducts, ok := helper.GetAllProductsForTenant(tenant.TenantId, tenant.GetPw(p.Dh, r.Context()), r.Context(), p.Dh)

	// Write the stuff to the cache
	go p.insertItems(allProducts)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get all products"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, allProducts) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// TODO: Change this, so we can get more pictures from one item
func (p *ProductHandler) HandleGettingPictureByProductId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
	w.WriteHeader(http.StatusBadRequest)
}
