package services

import (
	"ecomps.boobles.cloud/backend/internal/repository"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

// TODO: Create own TenantRepository!!!

type TenantService struct {
	tenantRepo repository.Repository[tenantstructs.Tenant]
}

func CreateNewTenantService(r repository.Repository[tenantstructs.Tenant]) *TenantService {
	return &TenantService{
		tenantRepo: r,
	}
}

func TenantToArgs(t tenantstructs.Tenant) []any {
	return []any{
		t.TenantId,
		t.TenantName,
		t.TenantCreation,
		t.TenantAdminUserId,
		t.TenantPwId,
	}
}
