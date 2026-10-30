package services

import (
	"context"
	"errors"
	"uuid"

	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
)

// =========================== General Stuff ===========================

func (t *TenantOAuthService) GetAllOAuthApps(ctx context.Context) ([]tenantoauthstructs.OAuthApps, error) {
	return t.oauthAppRepository.GetAllOAuthApps(ctx)
}

func (t *TenantOAuthService) GetAllOAuthSettings(ctx context.Context, langId uint) ([]tenantoauthstructs.OAuthSettings, error) {
	return t.oauthSettingsRepository.GetAllSettings(ctx, langId)
}

// =========================== OAuth tenant app stuff ===========================

func (t *TenantOAuthService) GetAllOAuthAppsForTenant(ctx context.Context, tenantId uuid.UUID) ([]tenantoauthstructs.TenantOAuthApps, error) {

	if tenantId.Compare(uuid.Nil()) == 0 || tenantId.Compare(uuid.Max()) == 0 {
		return []tenantoauthstructs.TenantOAuthApps{}, errors.New("TenantId invalid")
	}

	return t.oauthTenantAppRepository.GetAllOAuthAppsForTenant(ctx, tenantId)
}

// This func is used to add a new application to the tenant
// -> it does not add any accesstokens!
func (t *TenantOAuthService) AddOAuthAppForTenant(ctx context.Context, app tenantoauthstructs.TenantOAuthApps) error {

	if app.UserId <= 1 {
		return errors.New("UserId invalid")
	}

	if app.AppId == 0 {
		return errors.New("AppId invalid")
	}

	if app.TenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	return t.oauthTenantAppRepository.AddOAuthApplicationForTenant(ctx, app)
}

func (t *TenantOAuthService) DeleteOAuthAppForTenant(ctx context.Context, tenantId uuid.UUID, appId uint) error {

	if appId == 0 {
		return errors.New("AppId invalid")
	}

	if tenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	appToken, err := t.GetTokenForTenantApp(ctx, tenantId, appId)

	if err != nil {
		return err
	}

	if err := t.DeleteToken(ctx, appToken.TokenId); err != nil {
		return err
	}

	return t.oauthTenantAppRepository.DeleteApplicationForTenant(ctx, tenantId, appId)
}

// =========================== OAuth tenant token stuff ===========================

func (t *TenantOAuthService) GetTokenForTenantApp(ctx context.Context, tenantId uuid.UUID, appId uint) (tenantoauthstructs.TenantOAuthToken, error) {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return tenantoauthstructs.TenantOAuthToken{}, errors.New("TenantId invalid")
	}

	if appId == 0 {
		return tenantoauthstructs.TenantOAuthToken{}, errors.New("AppId invalid")
	}

	return t.oauthTenantTokenRepository.GetTokenForTenantApp(ctx, tenantId, appId)
}

func (t *TenantOAuthService) AddToken(ctx context.Context, code string) error {

	if code == "" {
		return errors.New("Invalid code")
	}

	// TODO: we need the config for the specified client here!
	// -> then get the access token and a refresh token
	// -> store both in the db

	return nil
}

func (t *TenantOAuthService) DeleteToken(ctx context.Context, tokenId uint) error {

	if tokenId == 0 {
		return errors.New("TokenId invalid")
	}

	return t.oauthTenantTokenRepository.DeleteToken(ctx, tokenId)
}

// =========================== OAuth tenant setting stuff ===========================

func (t *TenantOAuthService) GetSettingsForTenantApp(ctx context.Context, tenantId uuid.UUID, appId uint) ([]tenantoauthstructs.TenantSyncSettings, error) {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return []tenantoauthstructs.TenantSyncSettings{}, errors.New("TenantId invalid")
	}

	if appId == 0 {
		return []tenantoauthstructs.TenantSyncSettings{}, errors.New("AppId invalid")
	}

	return t.oauthTenantSettingsRepository.GetSettingsForApplication(ctx, tenantId, appId)
}

func (t *TenantOAuthService) AddSettingForTenantApp(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {

	if setting.AppId == 0 {
		return errors.New("AppId invalid")
	}

	if setting.UserId == 0 {
		return errors.New("UserId invalid")
	}

	if setting.TenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	return t.oauthTenantSettingsRepository.AddSettingForTenantApplication(ctx, setting)
}

func (t *TenantOAuthService) ChangeSettingForTenantApp(ctx context.Context, setting tenantoauthstructs.TenantSyncSettings) error {

	if setting.AppId == 0 {
		return errors.New("AppId invalid")
	}

	if setting.UserId == 0 {
		return errors.New("UserId invalid")
	}

	if setting.TenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	return t.oauthTenantSettingsRepository.ChangeSettingForTenantApplication(ctx, setting)
}
