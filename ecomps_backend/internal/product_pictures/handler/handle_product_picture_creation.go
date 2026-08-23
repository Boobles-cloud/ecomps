package handler // Handles the upload of a picture for a product
import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/internal/middleware"
	productstructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

func (p *ProductPictureHandler) HandleCreatingProductPicture(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Product | HandleCreatingProductPicture")

	picture, err := httputils.GetMetadataAndFileFromFormValues[productstructs.ProductPictures](r, ProductFormMetaDataKey,
		ProductFormFileKey, "PicturePath")

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	if result := p.Dh.ExecuteSQLStatement("InsertProductPicture", []any{picture.PictureName, picture.PicturePath, picture.PicturePosition,
		picture.ProductId, tenantId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed inserting item in database"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
