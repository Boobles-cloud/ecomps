package services

import (
	"context"

	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

// =========================== General Stuff ===========================

func (t *TenantOAuthService) GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error) {
	return t.oauthAppRepository.GetAllOAuthApps(ctx)
}

func (t *TenantOAuthService) GetAllOAuthSettings(ctx context.Context, langId uint) ([]tenantoauthstructs.OAuthSettings, error) {
	return t.oauthSettingsRepository.GetAllSettings(ctx, langId)
}

// ===========================
