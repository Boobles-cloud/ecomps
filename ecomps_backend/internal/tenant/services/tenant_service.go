package services

import (
	"context"
	"errors"
	"time"
	"uuid"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
)

func (t *TenantService) GetTenantById(ctx context.Context, id uuid.UUID) (tenantstructs.Tenant, error) {

	if id.Compare(uuid.Nil()) == 0 {
		return tenantstructs.Tenant{}, errors.New("TenantId isnt valid")
	}

	return t.tenantRepo.GetById(ctx, id)
}

func (t *TenantService) GetPw(ctx context.Context, id uuid.UUID) (string, error) {

	// TODO: change me!

	if id.Compare(uuid.Nil()) == 0 {
		return "", errors.New("TenantId isnt valid")
	}

	return "", errors.ErrUnsupported
}

func (t *TenantService) CreateTenant(ctx context.Context, tenant *tenantstructs.Tenant, userId uint) error {

	if tenant.TenantName == "" {
		return errors.New("Tenantname is empty")
	}

	tenant.TenantCreation = time.Now()

	if err := t.tenantRepo.Create(ctx, tenant, userId); err != nil {
		return err
	}

	// TODO: dont forget mee
	// return t.secerService.CreateTenantKey(ctx, tenant.TenantId)
	return nil
}

func (t *TenantService) UpdateTenant(ctx context.Context, tenant tenantstructs.Tenant) error {

	// TODO: also get the user id and check both!

	return t.tenantRepo.Update(ctx, tenant)
}

func (t *TenantService) Delete(ctx context.Context, userId uint, tenantId uuid.UUID) error {

	if userId <= 1 {
		return errors.New("UserId cant be 0 or 1")
	}

	if tenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId isnt valid")
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

func (t *TenantService) GetTenantDeletion(ctx context.Context, tenantId uuid.UUID) (tenantstructs.TenantDeletionStruct, error) {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return tenantstructs.TenantDeletionStruct{}, errors.New("TenantId inst valid")
	}

	return t.tenantRepo.GetTenantDeletion(ctx, tenantId)
}
