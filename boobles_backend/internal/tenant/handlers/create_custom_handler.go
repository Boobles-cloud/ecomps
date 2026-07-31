package handlers

import (
	"boobles.cloud/backend/caching"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
)

type TenantHandler struct {
	TenantCache *caching.CacheManager[tenantstructs.Tenant]
}

// Creates a new Tenant handler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(tc *caching.CacheManager[tenantstructs.Tenant]) *TenantHandler {
	return &TenantHandler{
		TenantCache: tc,
	}
}
