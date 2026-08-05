package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	"boobles.cloud/backend/internal/product/helper"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Gets a Product by the id
func (p *ProductHandler) HandleGettingProductById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Product | HandleGettingProductById] "+err.Error())
		w.WriteHeader(status)
	}

	id, err := strconv.Atoi(r.PathValue("product-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey)

	key := ProductCacheKey + strconv.Itoa(id)
	cacheItem, ok := p.ProductCache.GetItem(key)

	if ok {

		jsonData, err := json.Marshal(cacheItem)

		if err != nil {
			// If there is an error, just continue there
			goto withOutCache
		}

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:

	tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantById", []any{tenantId.(int)})

	if !ok || len(tenant) != 1 {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
	}

	product, ok := helper.GetProductById(uint(id), tenant[0].GetPw())

	go p.insertItem(*product)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting Product"))
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Gets all products by the given tenant id
func (p *ProductHandler) HandleGettingAllProductsByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		logging.Log(logging.Error, "[Product | HandleGettingAllProductsByTenantId] "+err.Error())

		w.WriteHeader(status)
	}

	tenantId, err := strconv.Atoi(r.PathValue("tenant-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	cacheItems, ok := p.ProductCache.GetItems(uint(tenantId))

	if ok || len(cacheItems) != 0 {

		jsonData, err := json.Marshal(cacheItems)

		if err != nil {
			// If there is an error, just continue there
			goto withOutCache
		}

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:

	// We always need to get the tenant, because all items are encrypted in the database
	tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantById", []any{tenantId})

	if !ok || len(tenant) != 1 {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
	}

	// Gets all decrypted products
	allProducts, ok := helper.GetAllProductsForTenant(tenant[0].TenantId, tenant[0].GetPw())

	// Write the stuff to the cache
	go p.insertItems(allProducts)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get all products"))
	}

	jsonData, err := json.Marshal(allProducts)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// TODO: Change this, so we can get more pictures from one item
func (p *ProductHandler) HandleGettingPictureById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
	w.WriteHeader(http.StatusBadRequest)
}
