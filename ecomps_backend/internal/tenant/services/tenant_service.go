package services

import (
	"context"
	"errors"
	"time"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

func (t *TenantService) GetTenantById(ctx context.Context, id uint) (tenantstructs.Tenant, error) {

	if id == 0 {
		return tenantstructs.Tenant{}, errors.New("TenantId cant be 0")
	}

	return t.tenantRepo.GetById(ctx, id)
}

func (t *TenantService) GetPw(ctx context.Context, id uint) (string, error) {

	if id == 0 {
		return "", errors.New("TenantId cant be 0")
	}

	return t.tenantRepo.GetPw(ctx, id)
}

func (t *TenantService) CreateTenant(ctx context.Context, tenant tenantstructs.Tenant, userId uint) error {

	if tenant.TenantName == "" {
		return errors.New("Tenantname is empty")
	}

	tenant.TenantCreation = time.Now()

	return t.tenantRepo.Create(ctx, tenant, userId)
}

func (t *TenantService) UpdateTenant(ctx context.Context, tenant tenantstructs.Tenant) error {

	tmpTenant, err := t.GetTenantById(ctx, tenant.TenantId)

	if err != nil {
		return err
	}

	if tmpTenant.TenantPwId != tenant.TenantPwId {
		return errors.New("Illegal operation!")
	}

	return t.tenantRepo.Update(ctx, tenant)
}

func (t *TenantService) Delete(ctx context.Context, userId, tenantId uint) error {

	if userId == 0 {
		return errors.New("UserId cant be 0")
	}

	if tenantId == 0 {
		return errors.New("TenantId cant be 0")
	}

	tenant, err := t.GetTenantById(ctx, tenantId)

	if err != nil {
		return err
	}

	if !tenant.IsUserAdmin(userId) {
		return errors.New("User cant delete tenant! Needs admin permission")
	}

	return t.tenantRepo.Delete(ctx, userId, tenantId)
}
