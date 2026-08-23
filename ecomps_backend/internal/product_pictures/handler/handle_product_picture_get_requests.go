package handler

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	productstructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// This handels getting a image by position and picture id
func (p *ProductPictureHandler) HandleGettingPictureByProductIdAndPosition(w http.ResponseWriter, r *http.Request) {

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

	picture, ok := database.QueryOne[productstructs.ProductPictures](r.Context(), p.Dh, "SelectProductPictureByProductIdAndPosition", []any{productId, positionId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get picture"))
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=86400")
	http.ServeFile(w, r, picture.PicturePath)
}

// Handels getting all image metadata for a product
func (p *ProductPictureHandler) HandleGettingAllImageMetadataForProductId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleGettingImageCountForProductId")

	productId, err := httputils.IntPathParam(r, "product_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	images, ok := database.QueryMany[productstructs.ProductPictures](r.Context(), p.Dh, "SelectProductPicturesByProductId", []any{productId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get items from database"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, images) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
