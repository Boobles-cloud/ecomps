package handler

import (
	"context"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// This handels getting a image by position and picture id
func (p *ProductPictureHandler) HandleGettingPictureByProductIdAndPosition(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleGettingPictureByProductIdAndPosition")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	positionId, err := httputils.IntPathParam(r, "position_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	picture, err := p.pictureService.GetByIdAndPosition(ctx, uint(productId), uint(positionId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=86400")
	http.ServeFile(w, r, picture.PicturePath)
}

// Handels getting all image metadata for a product
func (p *ProductPictureHandler) HandleGettingAllImageMetadataForProductId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleGettingImageCountForProductId")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	images, err := p.pictureService.GetByProductId(ctx, uint(productId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, images) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
