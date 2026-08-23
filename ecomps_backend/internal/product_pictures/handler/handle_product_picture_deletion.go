package handler

import (
	"errors"
	"net/http"
	"os"

	"ecomps.boobles.cloud/backend/database"
	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the deletion of a product picture
func (p *ProductPictureHandler) HandleDeletingProductPicture(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleDeletingProductPicture")

	productPictureId, err := httputils.IntPathParam(r, "picture_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// We first need to query the database to get the image path
	image, ok := database.QueryOne[productpicturetructs.ProductPictures](r.Context(), p.Dh, "SelectProductPictureById", []any{productPictureId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting item from database"))
		return
	}

	if err := os.Remove(image.PicturePath); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if result := p.Dh.ExecuteSQLStatement("DeleteProductById", []any{productPictureId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to delete item in database"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
