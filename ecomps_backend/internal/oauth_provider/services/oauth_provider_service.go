package services

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
)

func (s *OAuthProviderService) CreateProviderConfigByKey(ctx context.Context, key string) (oauth2.Config, error) {

	if key == "" {
		return oauth2.Config{}, errors.New("Empty string")
	}

	return oauth2.Config{}, errors.ErrUnsupported
}

func (s *OAuthProviderService) CreateProviderConfigById(ctx context.Context, appId uint) (oauth2.Config, error) {

	if appId == 0 {
		return oauth2.Config{}, errors.New("AppId invalid")
	}

	return oauth2.Config{}, errors.ErrUnsupported
}

func (s *OAuthProviderService) CheckProviderConfig(ctx context.Context, conf oauth2.Config) error {

	// TODO: could be usefull!
	return errors.ErrUnsupported
}
