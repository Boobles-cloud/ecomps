package services

import (
	permissionstruct "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
)

type PermissionService struct {
	permissionRepo repository.PermissionRepository
}

func CreateNewPermissionService(r repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		permissionRepo: r,
	}
}

func PermissionToArgs(p permissionstruct.Permission) []any {
	return []any{
		p.PermissionId,
		p.PermissionName,
		p.PermissionDescription,
		p.UserId,
	}
}
