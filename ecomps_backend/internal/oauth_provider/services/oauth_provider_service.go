package services

import (
	"context"
	"errors"

	oauthproviderstructs "ecomps.boobles.cloud/backend/internal/oauth_provider/oauth_provider_structs"
	"golang.org/x/oauth2"
)

// This struct is just used inside here!
type clientIdCliendSecretMapper struct {
	ClientId     string
	ClientSecret string
	Error        error
}

func (s *OAuthProviderService) CreateProviderConfigByKey(ctx context.Context, key string) (oauth2.Config, error) {

	if key == "" {
		return oauth2.Config{}, errors.New("Empty string")
	}

	provider, err := s.providerRepository.GetProviderByKey(ctx, key)

	if err != nil {
		return oauth2.Config{}, err
	}

	scopes, err := s.scopeRepository.GetScopesByProviderId(ctx, provider.ProviderId)

	if err != nil {
		return oauth2.Config{}, err
	}

	tmpConfig := s.getClientIdAndSecret(ctx, provider.ProviderKey)

	if tmpConfig.Error != nil {
		return oauth2.Config{}, tmpConfig.Error
	}

	return oauth2.Config{
		ClientID:     tmpConfig.ClientId,
		ClientSecret: tmpConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  provider.AuthUrl,
			TokenURL: provider.TokenUrl,
		},
		RedirectURL: "TODO: please add meeeeeee",
		Scopes:      mapScopesToStringArray(scopes),
	}, errors.ErrUnsupported
}

func (s *OAuthProviderService) CreateProviderConfigByAppId(ctx context.Context, appId uint) (oauth2.Config, error) {

	if appId == 0 {
		return oauth2.Config{}, errors.New("AppId invalid")
	}

	provider, err := s.providerRepository.GetProviderByAppId(ctx, appId)

	if err != nil {
		return oauth2.Config{}, err
	}

	scopes, err := s.scopeRepository.GetScopesByProviderId(ctx, provider.ProviderId)

	if err != nil {
		return oauth2.Config{}, err
	}

	tmpConfig := s.getClientIdAndSecret(ctx, provider.ProviderKey)

	if tmpConfig.Error != nil {
		return oauth2.Config{}, tmpConfig.Error
	}

	return oauth2.Config{
		ClientID:     tmpConfig.ClientId,
		ClientSecret: tmpConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  provider.AuthUrl,
			TokenURL: provider.TokenUrl,
		},
		RedirectURL: "TODO: please add meeeeeee",
		Scopes:      mapScopesToStringArray(scopes),
	}, errors.ErrUnsupported
}

func (s *OAuthProviderService) CheckProviderConfig(ctx context.Context, conf oauth2.Config) error {

	// TODO: could be usefull!
	// -> But do we need this?
	return errors.ErrUnsupported
}

func (s *OAuthProviderService) getClientIdAndSecret(ctx context.Context, key string) clientIdCliendSecretMapper {

	clientId, err := s.secretManager.GetClientId(ctx, key)

	if err != nil {
		return clientIdCliendSecretMapper{Error: err}
	}

	clientSecret, err := s.secretManager.GetClientSecret(ctx, key)

	if err != nil {
		return clientIdCliendSecretMapper{Error: err}
	}

	return clientIdCliendSecretMapper{
		ClientSecret: clientSecret.Value,
		ClientId:     clientId.Value,
		Error:        nil,
	}
}

func mapScopesToStringArray(scope []oauthproviderstructs.ProviderScopes) []string {

	tmp := make([]string, 0, len(scope))

	for i := range scope {
		tmp = append(tmp, scope[i].Scope)
	}

	return tmp
}
