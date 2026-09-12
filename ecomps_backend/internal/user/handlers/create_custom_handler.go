package handlers

import (
	authService "ecomps.boobles.cloud/backend/internal/auth/services"
	"ecomps.boobles.cloud/backend/internal/user/services"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
)

type UserHandler struct {
	UserCache   *caching.CacheManager[userstructs.UserStruct]
	userService *services.UserService
	authService *authService.AuthService
}

// Creates a new UserHandler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(uc *caching.CacheManager[userstructs.UserStruct], us *services.UserService, a *authService.AuthService) *UserHandler {
	return &UserHandler{
		UserCache:   uc,
		userService: us,
		authService: a,
	}
}
