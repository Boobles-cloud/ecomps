package services

import (
	oauthproviderstructs "ecomps.boobles.cloud/backend/internal/oauth_provider/oauth_provider_structs"
	"ecomps.boobles.cloud/backend/internal/secret_manager/services"
)

type OAuthProviderService struct {
	secretManager *services.SecretManagerService
}

func CreateNewOAuthProviderService(s *services.SecretManagerService) *OAuthProviderService {
	return &OAuthProviderService{
		secretManager: s,
	}
}

func OAuthProviderToArgs(p oauthproviderstructs.OAuthProvider) []any {
	return []any{
		p.ProviderId,
		p.ProviderKey,
		p.AuthUrl,
		p.TokenUrl,
		p.AppId,
	}
}

func ProviderScopesToArgs(s oauthproviderstructs.ProviderScopes) []any {
	return []any{
		s.ScopeId,
		s.Scope,
		s.ProviderId,
	}
}
