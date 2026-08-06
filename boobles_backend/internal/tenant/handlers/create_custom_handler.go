package handlers

import (
	"boobles.cloud/backend/caching"
	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
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
