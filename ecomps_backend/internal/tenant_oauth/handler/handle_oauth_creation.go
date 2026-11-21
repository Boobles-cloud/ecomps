package handler

import (
	"context"
	"net/http"
	"time"

	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handles adding a tenant oauth app
func (t *TenantOAuthHandler) HandleTenantOAuthAppCreation(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuth | HandleTenantOAuthAppCreation] ")

	oauthApp, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantoauthstructs.TenantOAuthApps](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := t.oauthService.AddOAuthAppForTenant(ctx, oauthApp); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Handles adding a setting to a specified tenant app
func (t *TenantOAuthHandler) HandleSettingsCreation(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuth | HandleSettingsCreation] ")

	setting, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantoauthstructs.TenantSyncSettings](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := t.oauthService.AddSettingForTenantApp(ctx, setting); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
