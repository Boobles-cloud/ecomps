package repositorys

import (
	"context"

	"ecomps.boobles.cloud/backend/database"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

type OAuthTenantAppRepository struct {
	*DatabaseRepository[tenantoauthstructs.TenantOAuthApps]
}

func NewOAuthTenantAppRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantoauthstructs.TenantOAuthApps]) *OAuthTenantAppRepository {
	return &OAuthTenantAppRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

type OAuthTenantSettingsRepository struct {
	*DatabaseRepository[tenantoauthstructs.TenantSyncSettings]
}

func NewOAuthTenantSettingsRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantoauthstructs.TenantSyncSettings]) *OAuthTenantSettingsRepository {
	return &OAuthTenantSettingsRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

type OAuthTenantTokenRepository struct {
	*DatabaseRepository[tenantoauthstructs.TenantOAuthToken]
}

func NewOAuthTenantTokenRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantoauthstructs.TenantOAuthToken]) *OAuthTenantTokenRepository {
	return &OAuthTenantTokenRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

// ======================================================================
// 							All Tenant-App funcs
// ======================================================================

func (t *OAuthTenantAppRepository) GetAllOAuthAppsForTenant(ctx context.Context, tenantId uint) ([]tenantoauthstructs.TenantOAuthApps, error)

func (t *OAuthTenantAppRepository) AddOAuthApplicationForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error

func (t *OAuthTenantAppRepository) DeleteApplicationForTenant(ctx context.Context, tenantId, appId uint) error

// ======================================================================
// 							All Tenant-Settings funcs
// ======================================================================

func (t *OAuthTenantSettingsRepository) GetSettingsForApplication(ctx context.Context, tenantId, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error)

func (t *OAuthTenantSettingsRepository) AddSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error

func (t *OAuthTenantSettingsRepository) ChangeSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error

// ======================================================================
// 							All Tenant-Token funcs
// ======================================================================

func (t *OAuthTenantTokenRepository) GetTokenForTenantApp(ctx context.Context, tenantId, appId uint) (tenantoauthstructs.TenantOAuthToken, error)

func (t *OAuthTenantTokenRepository) AddToken(ctx context.Context, token tenantoauthstructs.TenantOAuthToken) error

func (t *OAuthTenantTokenRepository) DeleteToken(ctx context.Context, tokenId uint) error
