package handler

import (
	"ecomps.boobles.cloud/backend/internal/product_pictures/services"
)

const (
	ProductFormFileKey     = "file-data"
	ProductFormMetaDataKey = "meta-data"
)

type ProductPictureHandler struct {
	pictureService *services.ProductPictureService
}

// Creates a new handler for products
func CreateNewProductHandler(p *services.ProductPictureService) *ProductPictureHandler {
	return &ProductPictureHandler{
		pictureService: p,
	}
}
