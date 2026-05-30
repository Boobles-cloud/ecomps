package userstructs

import "boobles.cloud/backend/database"

type UserPermission struct {
	PermissionId          uint
	PermissionName        string
	PermissionDescription string
	UserId                uint
}

// Sets a new permission for the given user permission struct
func (up *UserPermission) SetNewPermission() (uint, bool) {

	result := database.ExecuteSQLStatement("InsertUserPermission", database.Insert, []any{up.PermissionName, up.PermissionDescription, up.UserId})
	return result.LastId, result.Ok
}

// TODO
func (up *UserPermission) UpdatePermission() bool
