package userstructs

import (
	"context"

	"ecomps.boobles.cloud/backend/database"
)

type UserStruct struct {
	UserId        uint   `json:"UserId"`
	UserName      string `json:"UserName"`
	UserPW        string `json:"UserPw"`
	UserMail      string `json:"UserMail"`
	UserTel       string `json:"UserTel"`
	UserHas2FA    bool   `json:"UserHas2Fa"`
	UserHasTenant bool   `json:"UserHasTenant"`
	TenantId      uint   `json:"TenantId"`
}

// Creates a user in the database
func (u *UserStruct) CreateUserInDB(dh *database.DbHandler) (bool, uint) {

	if result := dh.ExecuteSQLStatement("InsertUser", []any{u.UserName, u.UserPW, u.UserMail, u.UserTel, u.UserHas2FA, u.UserHasTenant, u.TenantId}); result.Ok {
		return result.Ok, result.LastId
	}
	return false, 0
}

// Update a user in the database
func (u *UserStruct) UpdateUserInDB(dh *database.DbHandler) bool {
	return database.UpdateDatabaseEntry[UserStruct](dh, "UpdateUser", "UserId", *u)
}

// Returns all permissions a user has
func (u *UserStruct) GetPermissionsByUser(ctx context.Context, dh *database.DbHandler) ([]UserPermission, bool) {
	return database.QueryMany[UserPermission](ctx, dh, "SelectPermissionsByUserId", u.UserId)
}
