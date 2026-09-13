package repositorys

import (
	"context"
	"database/sql"
	"errors"

	"ecomps.boobles.cloud/backend/database"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
)

type OrderProductRepository struct {
	*DatabaseRepository[orderstructs.OrderProduct]
}

func NewOrderProductRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[orderstructs.OrderProduct]) *OrderProductRepository {
	return &OrderProductRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (o *OrderProductRepository) GetAllProductsByOrderId(ctx context.Context, orderId uint) ([]orderstructs.OrderProduct, error) {

	products, ok := database.QueryMany[orderstructs.OrderProduct](ctx, o.db, "SelectOrderProductsByOrderId", []any{orderId})

	if !ok {
		return products, sql.ErrNoRows
	}

	return products, nil
}

func (o *OrderProductRepository) CreateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error {

	if result := o.db.ExecuteSQLStatement(ctx, "InsertOrderProduct", o.toArgs(product)); !result.Ok {
		return errors.New("Failed to inerst item")
	}

	return nil
}

func (o *OrderProductRepository) UpdateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error {

	if ok := database.UpdateDatabaseEntry[orderstructs.OrderProduct](o.db, "UpdateOrderProduct", "OPId", product); !ok {
		return errors.New("Failed to update item")
	}

	return nil
}

func (o *OrderProductRepository) DeleteOrderProductsByOrderId(ctx context.Context, orderId uint) error {

	if result := o.db.ExecuteSQLStatement(ctx, "DeleteOrderProductsByOrderId", []any{orderId}); !result.Ok {
		return errors.New("Failed to delte items")
	}

	return nil
}

func (o *OrderProductRepository) DeleteOrderProduct(ctx context.Context, productId uint) error {

	if result := o.db.ExecuteSQLStatement(ctx, "DeleteOrderProductById", []any{productId}); !result.Ok {
		return errors.New("Failed to delte item")
	}

	return nil
}
