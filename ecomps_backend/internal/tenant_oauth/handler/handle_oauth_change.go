package handler

import (
	"context"
	"net/http"
	"time"

	tenantoauthstructs "ecomps.boobles.cloud/backend/internal/tenant_oauth/tenant_oauth_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

func (t *TenantOAuthHandler) HandleTenantOAuthAppChange(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuth | HandleTenantOAuthAppChange]")

	oauthApp, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantoauthstructs.TenantOAuthApps](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// TODO: implement me
	oauthApp.AppId = 1

	ctx.Done()
}

func (t *TenantOAuthHandler) HandleTenantSettingChange(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuth | HandleTenantSettingChange]")

	setting, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantoauthstructs.TenantSyncSettings](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := t.oauthService.ChangeSettingForTenantApp(ctx, setting); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
