package handlers

import (
	"ecomps.boobles.cloud/backend/database"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
)

type TenantHandler struct {
	TenantCache *caching.CacheManager[tenantstructs.Tenant]
	Dh          *database.DbHandler
}

// Creates a new Tenant handler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(tc *caching.CacheManager[tenantstructs.Tenant], d *database.DbHandler) *TenantHandler {
	return &TenantHandler{
		TenantCache: tc,
		Dh:          d,
	}
}
