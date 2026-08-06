package userstructs

import (
	"boobles.cloud/backend/database"
)

type UserPermission struct {
	PermissionId          uint
	PermissionName        string
	PermissionDescription string
	UserId                uint
}

// Sets a new permission for the given user permission struct
func (up *UserPermission) SetNewPermission(dh *database.DbHandler) (uint, bool) {

	result := dh.ExecuteSQLStatement("InsertUserPermission", []any{up.PermissionName, up.PermissionDescription, up.UserId})
	return result.LastId, result.Ok
}
