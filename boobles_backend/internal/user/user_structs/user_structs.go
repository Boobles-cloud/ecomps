package userstructs

import (
	"boobles.cloud/backend/database"
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
func (u *UserStruct) CreateUserInDB() (bool, uint) {

	if result := database.ExecuteSQLStatement("InsertUser", database.Insert, []any{u.UserName, u.UserPW, u.UserMail, u.UserTel, u.UserHas2FA, u.UserHasTenant, u.TenantId}); result.Ok {
		return result.Ok, result.LastId
	}
	return false, 0
}

// Update a user in the database
func (u *UserStruct) UpdateUserInDB() bool {
	return database.UpdateDatabaseEntry[UserStruct]("UpdateUser", "UserId", *u)
}

// Returns all permissions a user has
func (u *UserStruct) GetPermissionsByUser() ([]UserPermission, bool) {
	return database.QueryDatabase[UserPermission]("SelectPermissionsByUserId", []any{u.UserId})
}
