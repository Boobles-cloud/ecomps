package services

import (
	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
)

type ProductPictureService struct {
	productPictureRepo repository.ProductPictureRepository
}

func CreateNewProductPictureService(r repository.ProductPictureRepository) *ProductPictureService {
	return &ProductPictureService{
		productPictureRepo: r,
	}
}

func ProductPictureToArgs(p productpicturetructs.ProductPictures) []any {
	return []any{
		p.PictureId,
		p.PictureName,
		p.PicturePath,
		p.PicturePosition,
		p.ProductId,
		p.TenantId,
	}

}
