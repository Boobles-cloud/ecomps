package handlers

import (
	authService "ecomps.boobles.cloud/backend/internal/auth/services"
	"ecomps.boobles.cloud/backend/internal/tenant/services"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
)

type TenantHandler struct {
	tenantCache   *caching.CacheManager[tenantstructs.Tenant]
	tenantService *services.TenantService
	authService   *authService.AuthService
}

// Creates a new Tenant handler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(tc *caching.CacheManager[tenantstructs.Tenant], ts *services.TenantService, a *authService.AuthService) *TenantHandler {
	return &TenantHandler{
		tenantCache:   tc,
		tenantService: ts,
		authService:   a,
	}
}
