package handler

import (
	"context"
	"errors"
	"net/http"
	"time"
	"uuid"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handles getting all avaliable oauth settings
func (t *TenantOAuthHandler) HandleGettingAllOAuthApps(w http.ResponseWriter, r *http.Request) {

	// TODO: Maybe add getting languageId first

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "TenantOAuth | HandleGettingAllOAuthApps")

	allApps, err := t.oauthService.GetAllOAuthApps(ctx)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if ok := jsonutils.RespondWithJson(w, http.StatusOK, allApps); !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to respont with json"))
	}
}

// Handles getting all avaliable oauth settings
func (t *TenantOAuthHandler) HandleGettingAllOAuthSettings(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "TenantOAuth | HandleGettingAllOAuthSettings")

	langId, err := httputils.IntPathParam(r, "lang_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	allSettings, err := t.oauthService.GetAllOAuthSettings(ctx, uint(langId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if ok := jsonutils.RespondWithJson(w, http.StatusOK, allSettings); !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to respont with json"))
	}
}

// Handles getting all oauth-apps a tenant has activated
func (t *TenantOAuthHandler) HandleGettingAllTenantOAuthApps(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "TenantOAuth | HandleGettingAllTenantOAuthApps")

	tenantId := ctx.Value(middleware.TenantIdContextKey).(uuid.UUID)

	allApps, err := t.oauthService.GetAllOAuthAppsForTenant(ctx, tenantId)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if ok := jsonutils.RespondWithJson(w, http.StatusOK, allApps); !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to respont with json"))
	}
}

// Handles getting all settings for a specific tenant oauth-app
func (t *TenantOAuthHandler) HandleGettingAllSettingsForTenantOAuthApp(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "TenantOAuth | HandleGettingAllSettingsForTenantOAuthApp")

	tenantId := ctx.Value(middleware.TenantIdContextKey).(uuid.UUID)

	appId, err := httputils.IntPathParam(r, "app_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	allSettings, err := t.oauthService.GetSettingsForTenantApp(ctx, tenantId, uint(appId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if ok := jsonutils.RespondWithJson(w, http.StatusOK, allSettings); !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to respont with json"))
	}
}
