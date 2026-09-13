package repositorys

import (
	"context"
	"database/sql"
	"errors"

	"ecomps.boobles.cloud/backend/database"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
)

type OrderRepository struct {
	*DatabaseRepository[orderstructs.Order]
}

func NewOrderRepository(db *database.DbHandler, entityName string, toArgs ToArgsFunc[orderstructs.Order]) *OrderRepository {
	return &OrderRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (o *OrderRepository) GetOrderById(ctx context.Context, id uint) (orderstructs.Order, error) {

	order, ok := database.QueryOne[orderstructs.Order](ctx, o.db, "SelectOrderById", []any{id})

	if !ok {
		return order, sql.ErrNoRows
	}

	return order, nil
}

func (o *OrderRepository) GetAllOrdersByTenantId(ctx context.Context, tenantId uint) ([]orderstructs.Order, error) {

	orders, ok := database.QueryMany[orderstructs.Order](ctx, o.db, "SelectOrdersByTenantId", []any{tenantId})

	if !ok {
		return orders, sql.ErrNoRows
	}

	return orders, nil
}

func (o *OrderRepository) CreateOrder(ctx context.Context, order orderstructs.Order) error {

	if result := o.db.ExecuteSQLStatement(ctx, "InsertOrder", o.toArgs(order)); !result.Ok {
		return errors.New("Failed to insert item")
	}

	return nil
}

func (o *OrderRepository) UpdateOrder(ctx context.Context, order orderstructs.Order) error {

	if ok := database.UpdateDatabaseEntry[orderstructs.Order](o.db, "UpdateOrder", "OrderId", order); !ok {
		return errors.New("Failed to update item")
	}

	return nil
}

func (o *OrderRepository) DeleteOrder(ctx context.Context, orderId uint) error {

	if result := o.db.ExecuteSQLStatement(ctx, "DelteOrderById", []any{orderId}); !result.Ok {
		return errors.New("Failed to delete item")
	}

	return nil
}
