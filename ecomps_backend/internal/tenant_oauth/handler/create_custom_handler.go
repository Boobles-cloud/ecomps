package handler

import "ecomps.boobles.cloud/backend/internal/tenant_oauth/services"

const (
	stateCookieKey = "oauth_state"
)

type TenantOAuthHandler struct {
	oauthService *services.TenantOAuthService
}

func CreateNewTenantOAuthHandler(o *services.TenantOAuthService) *TenantOAuthHandler {
	return &TenantOAuthHandler{
		oauthService: o,
	}
}
