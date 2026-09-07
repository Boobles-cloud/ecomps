package handlers

import "ecomps.boobles.cloud/backend/internal/permission/services"

type PermissionHandler struct {
	permissionService *services.PermissionService
}

func CreateNewPermissionHandler(ps *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: ps,
	}
}
