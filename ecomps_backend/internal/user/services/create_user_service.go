package services

import (
	"ecomps.boobles.cloud/backend/internal/repository"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

// Used for all our buisness logic
type UserService struct {
	userRepo           repository.UserRepository
	userPermissionRepo repository.UserPermissionRepository
}

// TODO: We need the tenant service here -> for user deletion!!

func CreateNewUserService(r repository.UserRepository, rp repository.UserPermissionRepository) *UserService {
	return &UserService{
		userRepo:           r,
		userPermissionRepo: rp,
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
