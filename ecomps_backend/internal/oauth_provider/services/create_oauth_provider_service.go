package services

import "ecomps.boobles.cloud/backend/internal/secret_manager/services"

type OAuthProviderService struct {
	secretManager *services.SecretManagerService
}

func CreateNewOAuthProviderService(s *services.SecretManagerService) *OAuthProviderService {
	return &OAuthProviderService{
		secretManager: s,
	}
}
