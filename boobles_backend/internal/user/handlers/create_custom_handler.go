package handlers

import (
	"boobles.cloud/backend/caching"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
)

type UserHandler struct {
	UserCache       *caching.CacheManager[userstructs.UserStruct]
	PermissionCache *caching.CacheManager[userstructs.UserPermission]
}

// Creates a new UserHandler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(uc *caching.CacheManager[userstructs.UserStruct], pc *caching.CacheManager[userstructs.UserPermission]) *UserHandler {
	return &UserHandler{
		UserCache:       uc,
		PermissionCache: pc,
	}
}
