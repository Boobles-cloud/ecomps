package repository

import (
	"context"
	"uuid"

	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
)

type OrderRepository interface {
	GetOrderById(ctx context.Context, id uint) (orderstructs.Order, error)
	GetAllOrdersByTenantId(ctx context.Context, tenantId uuid.UUID) ([]orderstructs.Order, error)
	CreateOrder(ctx context.Context, order orderstructs.Order) (uint, error)
	UpdateOrder(ctx context.Context, order orderstructs.Order) error
	DeleteOrder(ctx context.Context, orderId uint) error
}

type OrderProductRepository interface {
	GetAllProductsByOrderId(ctx context.Context, orderId uint) ([]orderstructs.OrderProduct, error)
	CreateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error
	UpdateOrderProduct(ctx context.Context, product orderstructs.OrderProduct) error
	DeleteOrderProductsByOrderId(ctx context.Context, orderId uint) error
	DeleteOrderProduct(ctx context.Context, productId uint) error
}

type OrderStatusRepostiory interface {
	GetById(ctx context.Context, id, langId uint) (orderstructs.OrderStatus, error)
	GetAllByLangId(ctx context.Context, langId uint) ([]orderstructs.OrderStatus, error)
}
