package services

import (
	secretmanager "ecomps.boobles.cloud/backend/internal/secret_manager"
	smstructs "ecomps.boobles.cloud/backend/internal/secret_manager/sm_structs"
	"ecomps.boobles.cloud/backend/utils/caching"
)

type SecretManagerService struct {
	secretManagerRepo secretmanager.SecretManagerRepository
	cacheManager      *caching.CacheManager[smstructs.SecretEntry]
}

// IMPORTANT-NOTE:
//
// only call this func at the start of the server!!!!
func CreateNewSecretManagerService(s secretmanager.SecretManagerRepository) *SecretManagerService {
	return &SecretManagerService{
		secretManagerRepo: s,
		cacheManager:      caching.CreateNewCacheManager[smstructs.SecretEntry](),
	}
}
