package services

import (
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/tenant/services"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

// Used for all our buisness logic
type UserService struct {
	userRepo      repository.UserRepository
	tenantService *services.TenantService
}

// TODO: We need the tenant service here -> for user deletion!!
// -> Also get the auth package here!

func CreateNewUserService(r repository.UserRepository, t *services.TenantService) *UserService {
	return &UserService{
		userRepo:      r,
		tenantService: t,
	}
}

func UserToArgs(u userstructs.UserStruct) []any {
	return []any{
		u.UserId,
		u.UserName,
		u.UserPW,
		u.UserMail,
		u.UserTel,
		u.UserHas2FA,
		u.UserHasTenant,
		u.TenantId,
	}
}
