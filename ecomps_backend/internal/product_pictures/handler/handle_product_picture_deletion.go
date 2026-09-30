package handler

import (
	"context"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the deletion of a product picture
func (p *ProductPictureHandler) HandleDeletingProductPicture(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleDeletingProductPicture")

	productPictureId, err := httputils.IntPathParam(r, "picture_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := p.pictureService.Delete(ctx, uint(productPictureId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
