package repository

import (
	"context"

	oauthproviderstructs "ecomps.boobles.cloud/backend/internal/oauth_provider/oauth_provider_structs"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

type OAuthProviderRepository interface {
	GetProviderById(ctx context.Context, providerId uint) (oauthproviderstructs.OAuthProvider, error)
	GetProviderByAppId(ctx context.Context, appId uint) (oauthproviderstructs.OAuthProvider, error)
	GetProviderByKey(ctx context.Context, key string) (oauthproviderstructs.OAuthProvider, error)
}

type OAuthScopeRepository interface {
	GetScopesByProviderId(ctx context.Context, providerId uint) ([]oauthproviderstructs.ProviderScopes, error)
}

// Just for all application that we support
type OAuthAppRepository interface {
	GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error)
}

// All settings a user can chose from
type OAuthSettingsRepository interface {
	GetAllSettings(ctx context.Context, langId uint) ([]tenantoauthstructs.OAuthSettings, error)
}
