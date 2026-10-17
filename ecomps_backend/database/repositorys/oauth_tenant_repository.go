package repositorys

import (
	"context"
	"database/sql"
	"errors"

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

func (t *OAuthTenantAppRepository) GetAllOAuthAppsForTenant(ctx context.Context, tenantId uint) ([]tenantoauthstructs.TenantOAuthApps, error) {

	allApps, ok := database.QueryMany[tenantoauthstructs.TenantOAuthApps](ctx, t.db, "SelectAllAppsByTenantId", []any{tenantId})

	if !ok {
		return allApps, sql.ErrNoRows
	}

	return allApps, nil
}

func (t *OAuthTenantAppRepository) AddOAuthApplicationForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error {
	// TODO
	return errors.ErrUnsupported
}

func (t *OAuthTenantAppRepository) DeleteApplicationForTenant(ctx context.Context, tenantId, appId uint) error {
	// TODO
	return errors.ErrUnsupported
}

// ======================================================================
// 							All Tenant-Settings funcs
// ======================================================================

func (t *OAuthTenantSettingsRepository) GetSettingsForApplication(ctx context.Context, tenantId, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error) {

	allSettings, ok := database.QueryMany[tenantoauthstructs.TenantSyncSettings](ctx, t.db, "SelectAllSettingsByTenantAndAppId", []any{tenantId, appId})

	if !ok {
		return allSettings, sql.ErrNoRows
	}

	return allSettings, nil
}

func (t *OAuthTenantSettingsRepository) AddSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {
	// TODO
	return errors.ErrUnsupported
}

func (t *OAuthTenantSettingsRepository) ChangeSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {
	// TODO
	return errors.ErrUnsupported
}

// ======================================================================
// 							All Tenant-Token funcs
// ======================================================================

func (t *OAuthTenantTokenRepository) GetTokenForTenantApp(ctx context.Context, tenantId, appId uint) (tenantoauthstructs.TenantOAuthToken, error) {

	token, ok := database.QueryOne[tenantoauthstructs.TenantOAuthToken](ctx, t.db, "SelectOAuthTokenByTenantAndAppId", []any{})

	if !ok {
		return token, sql.ErrNoRows
	}

	return token, nil
}

func (t *OAuthTenantTokenRepository) AddToken(ctx context.Context, token tenantoauthstructs.TenantOAuthToken) error {
	// TODO
	return errors.ErrUnsupported
}

func (t *OAuthTenantTokenRepository) DeleteToken(ctx context.Context, tokenId uint) error {
	// TODO
	return errors.ErrUnsupported
}
