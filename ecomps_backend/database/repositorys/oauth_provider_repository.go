package repositorys

import (
	"context"
	"database/sql"

	"ecomps.boobles.cloud/backend/database"
	oauthproviderstructs "ecomps.boobles.cloud/backend/internal/oauth_provider/oauth_provider_structs"
)

type OAuhtProviderRepository struct {
	*DatabaseRepository[oauthproviderstructs.OAuthProvider]
}

func NewOAuthProviderRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[oauthproviderstructs.OAuthProvider]) *OAuhtProviderRepository {
	return &OAuhtProviderRepository{
		DatabaseRepository: NewDatabaseRepository[oauthproviderstructs.OAuthProvider](db, entityName, toArgs),
	}
}

type OAuthScopeRepository struct {
	*DatabaseRepository[oauthproviderstructs.ProviderScopes]
}

func NewOAuhtScopeRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[oauthproviderstructs.ProviderScopes]) *OAuthScopeRepository {
	return &OAuthScopeRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

// ======================================================================
// 							All OAuthProvider funcs
// ======================================================================

func (p *OAuhtProviderRepository) GetProviderById(ctx context.Context, providerId uint) (oauthproviderstructs.OAuthProvider, error) {

	provider, ok := database.QueryOne[oauthproviderstructs.OAuthProvider](ctx, p.db, "SelectProviderByProviderId", []any{providerId})

	if !ok {
		return provider, sql.ErrNoRows
	}

	return provider, nil
}

func (p *OAuhtProviderRepository) GetProviderByAppId(ctx context.Context, appId uint) (oauthproviderstructs.OAuthProvider, error) {

	provider, ok := database.QueryOne[oauthproviderstructs.OAuthProvider](ctx, p.db, "SelectProviderByAppId", []any{appId})

	if !ok {
		return provider, sql.ErrNoRows
	}

	return provider, nil
}

func (p *OAuhtProviderRepository) GetProviderByKey(ctx context.Context, key string) (oauthproviderstructs.OAuthProvider, error) {

	provider, ok := database.QueryOne[oauthproviderstructs.OAuthProvider](ctx, p.db, "SelectProviderByKey", []any{key})

	if !ok {
		return provider, sql.ErrNoRows
	}

	return provider, nil
}

// ======================================================================
// 							All OAuhtScope funcs
// ======================================================================

func (s *OAuthScopeRepository) GetScopesByProviderId(ctx context.Context, providerId uint) ([]oauthproviderstructs.ProviderScopes, error) {

	scopes, ok := database.QueryMany[oauthproviderstructs.ProviderScopes](ctx, s.db, "SelectScopesByProviderId", []any{providerId})

	if !ok {
		return scopes, sql.ErrNoRows
	}

	return scopes, nil
}
