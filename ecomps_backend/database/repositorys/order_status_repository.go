package repositorys

import (
	"context"
	"database/sql"

	"ecomps.boobles.cloud/backend/database"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
)

type OrderStatusRepository struct {
	*DatabaseRepository[orderstructs.OrderStatus]
}

func NewOrderStatusRepostiory(db *database.DbHandler, entityName string, toArgs ToArgsFunc[orderstructs.OrderStatus]) *OrderStatusRepository {
	return &OrderStatusRepository{
		DatabaseRepository: NewDatabaseRepository(db, entityName, toArgs),
	}
}

func (o *OrderStatusRepository) GetById(ctx context.Context, id, langId uint) (orderstructs.OrderStatus, error) {

	status, ok := database.QueryOne[orderstructs.OrderStatus](ctx, o.db, "SelectOrderStatusByIdAndLanguageId", []any{id, langId})

	if !ok {
		return status, sql.ErrNoRows
	}

	return status, nil
}

func (o *OrderStatusRepository) GetAllByLangId(ctx context.Context, langId uint) ([]orderstructs.OrderStatus, error) {

	status, ok := database.QueryMany[orderstructs.OrderStatus](ctx, o.db, "SelectAllStatusByLanguageId", []any{langId})

	if !ok {
		return status, sql.ErrNoRows
	}

	return status, nil
}
