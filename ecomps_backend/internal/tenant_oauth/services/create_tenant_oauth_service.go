package services

import "ecomps.boobles.cloud/backend/internal/repository"

type TenantOAuthService struct {
	oauthAppRepository            repository.OAuthAppRepository
	oauthSettingsRepository       repository.OAuthSettingsRepository
	oauthTenantSettingsRepository repository.OAuthTenantSettingsRepository
	oauthTenantAppRepository      repository.OAuthTenantAppRepository
	oauthTenantTokenRepository    repository.OAuthTenantTokenRepository
}

func CreateNewTenantOAuthService(app repository.OAuthAppRepository,
	setting repository.OAuthSettingsRepository,
	tenantSet repository.OAuthTenantSettingsRepository,
	tenantApp repository.OAuthTenantAppRepository,
	tenantToken repository.OAuthTenantTokenRepository) *TenantOAuthService {

	return &TenantOAuthService{
		oauthAppRepository:            app,
		oauthSettingsRepository:       setting,
		oauthTenantSettingsRepository: tenantSet,
		oauthTenantAppRepository:      tenantApp,
		oauthTenantTokenRepository:    tenantToken,
	}
}

// TODO: Add to args funcs here!
