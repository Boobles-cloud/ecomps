package services

import (
	"context"
	"errors"

	permissionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
)

func (p *PermissionService) GetPermissionById(ctx context.Context, id uint) (permissionstructs.Permission, error) {
	return p.permissionRepo.GetById(ctx, id)
}

// Deprecated: This service doesnt support this
// TODO: Maybe implement this in the future
func (p *PermissionService) GetAllByTenantIdctx(ctx context.Context, tenantId uint) ([]permissionstructs.Permission, error) {
	return []permissionstructs.Permission{}, errors.ErrUnsupported
}

// Deprecated: This service doesnt support this
// TODO: Maybe implement this in the future
func (p *PermissionService) CreatePermission(ctx context.Context, per permissionstructs.Permission) (uint, error) {
	return 0, errors.ErrUnsupported
}

// Deprecated: This service doesnt support this
// TODO: Maybe implement this in the future
func (p *PermissionService) UpdatePermission(ctx context.Context, per permissionstructs.Permission, filtername string) error {
	return errors.ErrUnsupported
}

// Deprecated: This service doesnt support this
// TODO: Maybe implement this in the future
func (p *PermissionService) DeletePermission(ctx context.Context, id, tenanId uint) error {
	return errors.ErrUnsupported
}

func (p *PermissionService) AsignUserPermssion(ctx context.Context, userId, permissionId uint) error {

	if userId == 0 {
		return errors.New("UserId cant be 0")
	}

	allPermissions, err := p.permissionRepo.GetAllPermissionsForUserId(ctx, userId)

	if err != nil {
		return err
	}

	for i := range allPermissions {
		if allPermissions[i].PermissionId == permissionId {
			return errors.New("User already has wanted permission")
		}
	}

	return p.permissionRepo.AsingUserPermission(ctx, userId, permissionId)
}

func (p *PermissionService) GetAllPermissionsByLanguageId(ctx context.Context, langId uint) ([]permissionstructs.Permission, error) {

	if langId == 0 {
		return []permissionstructs.Permission{}, errors.New("LanguageId cant be 0")
	}

	return p.permissionRepo.GetAllPermissionsByLanguageId(ctx, langId)
}

func (p *PermissionService) GetAllPermissionsForUserId(ctx context.Context, userId uint) ([]permissionstructs.Permission, error) {

	if userId == 0 {
		return []permissionstructs.Permission{}, errors.New("UserId cant be 0")
	}
	return p.permissionRepo.GetAllPermissionsForUserId(ctx, userId)
}

func (p *PermissionService) RemoveUserPermission(ctx context.Context, userId, permissionId uint) error {
	return p.permissionRepo.RemoveUserPermission(ctx, userId, permissionId)
}
