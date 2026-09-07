package services

import (
	"context"
	"errors"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

func (t *TenantService) GetTenantById(ctx context.Context, id uint) (tenantstructs.Tenant, error) {
	return t.tenantRepo.GetById(ctx, id)
}

// Deprecated: For this service use "GetTenantById"
func (t *TenantService) GetTenant(ctx context.Context, tenantId uint) (tenantstructs.Tenant, error) {
	return tenantstructs.Tenant{}, errors.ErrUnsupported
}

// Deprecated: For this service use "GetTenantById"
func (t *TenantService) GetAllByTenantId(ctx context.Context, tenantId uint) ([]tenantstructs.Tenant, error) {
	return []tenantstructs.Tenant{}, errors.ErrUnsupported
}

// TODO
func (t *TenantService) CreateTenant(ctx context.Context, tenant tenantstructs.Tenant) (uint, error) {
	return 0, errors.ErrUnsupported
}

// TODO
func (t *TenantService) UpdateTenant(ctx context.Context, tenant tenantstructs.Tenant) error {
	return errors.ErrUnsupported
}

// TODO
func (t *TenantService) Delete(ctx context.Context, userId, tenantId uint) error {
	return errors.ErrUnsupported
}
