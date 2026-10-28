package services

import secretmanager "ecomps.boobles.cloud/backend/internal/secret_manager"

type SecretManagerService struct {
	// TODO: add a cache manager here!!
	secretManagerRepo secretmanager.SecretManagerRepository
}

func CreateNewSecretManagerService(s secretmanager.SecretManagerRepository) *SecretManagerService {
	return &SecretManagerService{
		secretManagerRepo: s,
	}
}
