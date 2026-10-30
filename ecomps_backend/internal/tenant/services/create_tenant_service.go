package services

import (
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/secret_manager/services"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

type TenantService struct {
	tenantRepo repository.TenantRepository
	smService  *services.SecretManagerService
}

func CreateNewTenantService(r repository.TenantRepository, s *services.SecretManagerService) *TenantService {
	return &TenantService{
		tenantRepo: r,
		smService:  s,
	}
}

func TenantToArgs(t tenantstructs.Tenant) []any {
	return []any{
		t.TenantId,
		t.TenantName,
		t.TenantCreation,
		t.TenantAdminUserId,
	}
}
