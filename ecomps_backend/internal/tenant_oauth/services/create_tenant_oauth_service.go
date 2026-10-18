package services

import (
	"ecomps.boobles.cloud/backend/internal/repository"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

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

func OAuthAppsToArgs(item tenantoauthstructs.OAuthApps) []any {
	return []any{
		item.AppId,
		item.AppName,
		item.AppLink,
	}
}

func OAuthSettingsToArgs(item tenantoauthstructs.OAuthSettings) []any {
	return []any{
		item.SettingId,
		item.SettingName,
		item.SettingDescription,
		item.LangId,
	}
}

func OAuthTenantToken(item tenantoauthstructs.TenantOAuthToken) []any {
	return []any{
		item.TokenId,
		item.TokenVal,
		item.RefreshTokenVal,
		item.LastChanged,
		item.AppId,
		item.TenantId,
	}
}

func OAuthTenantApps(item tenantoauthstructs.TenantOAuthApps) []any {
	return []any{
		item.TenantAppId,
		item.SyncInterval,
		item.SyncActivated,
		item.SyncSuccessfull,
		item.SyncErrorMsg,
		item.UserId,
		item.AppId,
		item.TenantId,
	}
}

func OAuhtTEnantSyncSettings(item tenantoauthstructs.TenantSyncSettings) []any {
	return []any{
		item.TenantSettingId,
		item.LastChanged,
		item.SettingId,
		item.UserId,
		item.TenantId,
	}
}
