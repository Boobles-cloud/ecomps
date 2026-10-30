package repositorys

import (
	"context"
	"database/sql"
	"errors"
	"uuid"

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

func (t *OAuthTenantAppRepository) GetAllOAuthAppsForTenant(ctx context.Context, tenantId uuid.UUID) ([]tenantoauthstructs.TenantOAuthApps, error) {

	allApps, ok := database.QueryMany[tenantoauthstructs.TenantOAuthApps](ctx, t.db, "SelectAllAppsByTenantId", []any{tenantId})

	if !ok {
		return allApps, sql.ErrNoRows
	}

	return allApps, nil
}

func (t *OAuthTenantAppRepository) AddOAuthApplicationForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error {

	args := t.toArgs(app)
	// Remove the first entry
	args = args[1:]

	if result := t.db.ExecuteSQLStatement(ctx, "InsertOAuthAppForTenant", args); !result.Ok {
		return errors.New("Failed to create item")
	}

	return nil
}

func (t *OAuthTenantAppRepository) DeleteApplicationForTenant(ctx context.Context, tenantId uuid.UUID, appId uint) error {

	if result := t.db.ExecuteSQLStatement(ctx, "DeleteTenantOAuthAppByTenantAndAppId", []any{tenantId, appId}); !result.Ok {
		return errors.New("Failed to delete item")
	}

	return nil
}

// ======================================================================
// 							All Tenant-Settings funcs
// ======================================================================

func (t *OAuthTenantSettingsRepository) GetSettingsForApplication(ctx context.Context, tenantId uuid.UUID, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error) {

	allSettings, ok := database.QueryMany[tenantoauthstructs.TenantSyncSettings](ctx, t.db, "SelectAllSettingsByTenantAndAppId", []any{tenantId, appId})

	if !ok {
		return allSettings, sql.ErrNoRows
	}

	return allSettings, nil
}

func (t *OAuthTenantSettingsRepository) AddSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {

	args := t.toArgs(setting)
	args = args[1:]

	if result := t.db.ExecuteSQLStatement(ctx, "InsertSettingForTenantApplication", args); !result.Ok {
		return errors.New("Failed to create item")
	}

	return nil
}

func (t *OAuthTenantSettingsRepository) ChangeSettingForTenantApplication(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {

	// NOTE: we needed to implement this the custom way
	// Because the database update func doenst support two filter values at the moment
	// TODO: Come back here and change me

	_, err := t.db.DbConnection.Exec("UPDATE TenantSyncSettings SET LastChanged = ?, SettingId = ?, UserId = ? WHERE TenantSettingId = ?", []any{setting.LastChanged, setting.SettingId, setting.UserId, setting.TenantSettingId})

	return err
}

// ======================================================================
// 							All Tenant-Token funcs
// ======================================================================

// TODO: Move this to our key-management-provider

func (t *OAuthTenantTokenRepository) GetTokenForTenantApp(ctx context.Context, tenantId uuid.UUID, appId uint) (tenantoauthstructs.TenantOAuthToken, error) {

	token, ok := database.QueryOne[tenantoauthstructs.TenantOAuthToken](ctx, t.db, "SelectOAuthTokenByTenantAndAppId", []any{})

	if !ok {
		return token, sql.ErrNoRows
	}

	return token, nil
}

func (t *OAuthTenantTokenRepository) AddToken(ctx context.Context, token tenantoauthstructs.TenantOAuthToken) error {

	args := t.toArgs(token)
	args = args[1:]

	if result := t.db.ExecuteSQLStatement(ctx, "InsertOAuthToken", args); !result.Ok {
		return errors.New("Failed to insert item")
	}

	return nil
}

func (t *OAuthTenantTokenRepository) DeleteToken(ctx context.Context, tokenId uuid.UUID) error {

	if result := t.db.ExecuteSQLStatement(ctx, "DeleteOAuthTokenById", []any{tokenId}); !result.Ok {
		return errors.New("Failed to delete item")
	}

	return nil
}
