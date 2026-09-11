package repositorys

import (
	"context"
	"database/sql"
	"errors"

	"ecomps.boobles.cloud/backend/database"
	productpicturetructs "ecomps.boobles.cloud/backend/internal/product_pictures/product_pictures_structs"
)

type ProductPictureRepository struct {
	*DatabaseRepository[productpicturetructs.ProductPictures]
}

func NewProductPictureRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[productpicturetructs.ProductPictures]) *ProductPictureRepository {
	return &ProductPictureRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (p *ProductPictureRepository) GetByIdAndPosition(ctx context.Context, productId, position uint) (productpicturetructs.ProductPictures, error) {

	product, ok := database.QueryOne[productpicturetructs.ProductPictures](ctx, p.db, "SelectProductPictureByProductIdAndPosition", []any{productId, position})

	if !ok {
		return product, sql.ErrNoRows
	}

	return product, nil
}

func (p *PermissionRepository) GetById(ctx context.Context, id uint) (productpicturetructs.ProductPictures, error) {

	product, ok := database.QueryOne[productpicturetructs.ProductPictures](ctx, p.db, "SelectProductPictureById", []any{id})

	if !ok {
		return product, sql.ErrNoRows
	}

	return product, nil
}

func (p *ProductPictureRepository) GetByProductId(ctx context.Context, id uint) ([]productpicturetructs.ProductPictures, error) {

	product, ok := database.QueryMany[productpicturetructs.ProductPictures](ctx, p.db, "SelectProductPicturesByProductId", []any{id})

	if !ok {
		return product, sql.ErrNoRows
	}

	return product, nil
}

func (p *ProductPictureRepository) Create(ctx context.Context, item productpicturetructs.ProductPictures) (uint, error) {

	result := p.db.ExecuteSQLStatement(ctx, "InsertProductPicture", p.toArgs(item))

	if !result.Ok {
		return 0, sql.ErrNoRows
	}

	return result.LastId, nil
}

func (p *ProductPictureRepository) Delete(ctx context.Context, id uint) error {

	if result := p.db.ExecuteSQLStatement(ctx, "DeleteProductPictureById", []any{id}); !result.Ok {
		return errors.New("Failed to delete item")
	}

	return nil
}
