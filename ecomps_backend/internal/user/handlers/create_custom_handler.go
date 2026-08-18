package handlers

import (
	"ecomps.boobles.cloud/backend/caching"
	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

type UserHandler struct {
	UserCache       *caching.CacheManager[userstructs.UserStruct]
	PermissionCache *caching.CacheManager[userstructs.UserPermission]
	Dh              *database.DbHandler
}

// Creates a new UserHandler
// NOTE: We dont use the cache here, but for future stuff its already there
func CreateNewUserHander(uc *caching.CacheManager[userstructs.UserStruct], pc *caching.CacheManager[userstructs.UserPermission], d *database.DbHandler) *UserHandler {
	return &UserHandler{
		UserCache:       uc,
		PermissionCache: pc,
		Dh:              d,
	}
}
