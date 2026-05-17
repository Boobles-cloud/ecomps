package userstructs

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

func (u *UserStruct) CreateUserInDB() bool

// Returns the tenant for the given user
// TODO
func (u *UserStruct) GetTenantByUser()

// Returns all permissions a user has
// TODO
func (u *UserStruct) GetPermissionsByUser() []UserPermission
