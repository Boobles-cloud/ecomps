package repository

import (
	"context"
	"uuid"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

type TenantRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (tenantstructs.Tenant, error)
	GetTenantDeletion(ctx context.Context, id uuid.UUID) (tenantstructs.TenantDeletionStruct, error)
	Create(ctx context.Context, tenant *tenantstructs.Tenant, userId uint) error
	Update(ctx context.Context, tenant tenantstructs.Tenant) error
	Delete(ctx context.Context, userId uint, tenantId uuid.UUID) error
}

// All apps a tenant has active
type OAuthTenantAppRepository interface {
	GetAllOAuthAppsForTenant(ctx context.Context, tenantId uuid.UUID) ([]tenantoauthstructs.TenantOAuthApps, error)
	AddOAuthApplicationForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error
	DeleteApplicationForTenant(ctx context.Context, tenantId uuid.UUID, appId uint) error
	// TODO: add deletion
}

// All settings for the tenant apps
type OAuthTenantSettingsRepository interface {
	GetSettingsForApplication(ctx context.Context, tenantId uuid.UUID, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error)
	AddSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error
	ChangeSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error
	// TODO: add deletion
}

type OAuthTenantTokenRepository interface {
	GetTokenForTenantApp(ctx context.Context, tenantId uuid.UUID, appId uint) (tenantoauthstructs.TenantOAuthToken, error)
	AddToken(ctx context.Context, token tenantoauthstructs.TenantOAuthToken) error
	DeleteToken(ctx context.Context, tokenId uint) error
}
