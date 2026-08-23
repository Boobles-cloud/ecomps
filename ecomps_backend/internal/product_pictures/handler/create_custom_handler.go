package handler

import (
	"ecomps.boobles.cloud/backend/database"
)

const (
	ProductFormFileKey     = "file-data"
	ProductFormMetaDataKey = "meta-data"
)

type ProductPictureHandler struct {
	Dh *database.DbHandler
}

// Creates a new handler for products
func CreateNewProductHandler(d *database.DbHandler) *ProductPictureHandler {
	return &ProductPictureHandler{
		Dh: d,
	}
}
