package handler // Handles the upload of a picture for a product
import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	productstructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

func (p *ProductPictureHandler) HandleCreatingProductPicture(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Product | HandleCreatingProductPicture")

	picture, err := httputils.GetMetadataAndFileFromFormValues[productstructs.ProductPictures](r, ProductFormMetaDataKey,
		ProductFormFileKey, "PicturePath")

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	picture.TenantId = uint(tenantId)

	if _, err := p.pictureService.Create(ctx, picture); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
