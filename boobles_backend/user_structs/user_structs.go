package userstructs

import "boobles.cloud/backend/database"

type UserStruct struct {
	UserId        uint
	UserName      string
	UserPW        string
	UserMail      string
	UserTel       string
	UserHas2FA    bool
	UserHasTenant bool
	TenantId      uint
}

// Creates a user in the database
func (u *UserStruct) CreateUserInDB() (bool, uint) {

	if result := database.ExecuteSQLStatement("InsertUser", database.Insert, []any{u.UserName, u.UserPW, u.UserMail, u.UserTel, u.UserHas2FA, u.UserHasTenant, u.TenantId}); result.Ok {
		return result.Ok, result.LastId
	}
	return false, 0
}

// Update a user in the database
// TODO: Update this with the new update stuff
// func (u *UserStruct) UpdateUserInDB() bool {
// 	result := database.ExecuteSQL("UPDATE Users SET UserName = ?, UserPW = ?, UserMail = ?, UserTel = ?, UserHas2Fa = ?, UserHasTenant = ?, TenantId = ? WHERE UserId = ?", []any{u.UserName, u.UserPW, u.UserTel, u.UserHas2FA, u.UserHasTenant, u.TenantId, u.UserId})
// 	return result.Ok
// }

// Returns the tenant for the given user
// TODO
func (u *UserStruct) GetTenantByUser()

// Returns all permissions a user has
func (u *UserStruct) GetPermissionsByUser() ([]UserPermission, bool) {
	return database.QueryDatabase[UserPermission]("SelectPermissionsByUserId", []any{u.UserId})
}
