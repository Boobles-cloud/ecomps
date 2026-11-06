package services

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

// TODO: Maybe in the future move this to a own service?

func (t *TenantOAuthService) createProviderConf(ctx context.Context, providerKey string) (oauth2.Config, error) {

	clientSecret, err := t.smService.GetClientSecret(ctx, providerKey)

	if err != nil {
		return oauth2.Config{}, err
	}

	clientId, err := t.smService.GetClientId(ctx, providerKey)

	if err != nil {
		return oauth2.Config{}, err
	}

	return oauth2.Config{
		ClientSecret: clientSecret.Value,
		ClientID:     clientId.Key,
		RedirectURL:  "TODO",
		Endpoint:     endpoints.Amazon,
		Scopes:       []string{},
	}, nil
}

// TODO: add a seperate func that builds our endpoint from database information
