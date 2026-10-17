package repositorys

import (
	"context"
	"database/sql"

	"ecomps.boobles.cloud/backend/database"
	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

type OAuthAppRepository struct {
	*DatabaseRepository[tenantoauthstructs.OAuthApps]
}

func NewOAuthAppRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantoauthstructs.OAuthApps]) *OAuthAppRepository {
	return &OAuthAppRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

type OAuthSettingsRepository struct {
	*DatabaseRepository[tenantoauthstructs.OAuthSettings]
}

func NewOAuthSettingsRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[tenantoauthstructs.OAuthSettings]) *OAuthSettingsRepository {
	return &OAuthSettingsRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

// ======================================================================
// 							All OAuhtApp funcs
// ======================================================================

func (o *OAuthAppRepository) GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error) {

	allApps, ok := database.QueryMany[tenantoauthstructs.OAuthApps](ctx, o.db, "SelectAllOAuthApps", []any{})

	if !ok {
		return allApps, sql.ErrNoRows
	}

	return allApps, nil
}

// ======================================================================
// 							All OAuhtSettings funcs
// ======================================================================

func (o *OAuthSettingsRepository) GetAllSettings(ctx context.Context) ([]tenantoauthstructs.OAuthSettings, error) {

	allSettings, ok := database.QueryMany[tenantoauthstructs.OAuthSettings](ctx, o.db, "SelectAllOAuthSettings", []any{})

	if !ok {
		return allSettings, sql.ErrNoRows
	}

	return allSettings, nil
}
