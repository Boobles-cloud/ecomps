package services

import (
	"context"
	"errors"
	"os"

	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
)

func (p *ProductPictureService) GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error) {

	if productId == 0 {
		return productpicturetructs.ProductPictures{}, errors.New("Product Id cant be 0")
	}

	return p.productPictureRepo.GetByIdAndPosition(ctx, productId, position)
}

func (p *ProductPictureService) GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error) {

	if id == 0 {
		return []productpicturetructs.ProductPictures{}, errors.New("Product id cant be 0")
	}

	return p.productPictureRepo.GetByProductId(ctx, id)
}

func (p *ProductPictureService) GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error) {

	if id == 0 {
		return productpicturetructs.ProductPictures{}, errors.New("Picture id cant be 0")
	}

	return p.productPictureRepo.GetById(ctx, id)

}

func (p *ProductPictureService) Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error) {

	if item.PicturePath == "" {
		return 0, errors.New("Picture path cant be empty")
	}

	if item.ProductId == 0 {
		return 0, errors.New("Product Id cant be 0")
	}

	if item.TenantId == 0 {
		return 0, errors.New("Tenant Id cant be 0")
	}

	return p.productPictureRepo.Create(ctx, item)
}

func (p *ProductPictureService) Delete(ctx context.Context, id uint) error {

	if id == 0 {
		return errors.New("Product Picture Id cant be 0")
	}

	picture, err := p.GetById(ctx, id)

	if err != nil {
		return err
	}

	if err := os.Remove(picture.PicturePath); err != nil {
		return err
	}

	return p.productPictureRepo.Delete(ctx, id)
}
