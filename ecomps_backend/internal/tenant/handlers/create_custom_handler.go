package handlers

import (
	"ecomps.boobles.cloud/backend/internal/tenant/services"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
)

type TenantHandler struct {
	tenantCache   *caching.CacheManager[tenantstructs.Tenant]
	tenantService *services.TenantService
}

// Creates a new Tenant handler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(tc *caching.CacheManager[tenantstructs.Tenant], ts *services.TenantService) *TenantHandler {
	return &TenantHandler{
		tenantCache:   tc,
		tenantService: ts,
	}
}
