package services

import (
	"context"
	"errors"
	"uuid"

	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	"ecomps.boobles.cloud/backend/utils/crypto"
)

func (o *OrderService) GetOrderById(ctx context.Context, id uint) (orderstructs.Order, error) {

	if id == 0 || id == 1 {
		return orderstructs.Order{}, errors.New("Id cant be 0 or 1")
	}

	encryptedOrder, err := o.orderRepository.GetOrderById(ctx, id)

	if err != nil {
		return encryptedOrder, err
	}

	pw, err := o.tenantService.GetPw(ctx, encryptedOrder.TenantId)

	if err != nil {
		return encryptedOrder, err
	}

	order, ok := crypto.Decrypt[orderstructs.Order](&encryptedOrder, pw)

	if !ok {
		return encryptedOrder, errors.New("Failed to decrypt")
	}

	allProducts, err := o.productRespository.GetAllProductsByOrderId(ctx, order.OrderId)

	if err != nil {
		return encryptedOrder, err
	}

	order.Products = allProducts

	return *order, nil
}

func (o *OrderService) GetAllOrdersByTenantId(ctx context.Context, tenantId uuid.UUID) ([]orderstructs.Order, error) {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return []orderstructs.Order{}, errors.New("TenantId invalid")
	}

	encryptedOrder, err := o.orderRepository.GetAllOrdersByTenantId(ctx, tenantId)

	if err != nil {
		return encryptedOrder, err
	}

	pw, err := o.tenantService.GetPw(ctx, encryptedOrder[0].TenantId)

	if err != nil {
		return encryptedOrder, err
	}

	orders := make([]orderstructs.Order, 0, len(encryptedOrder))

	for i := range encryptedOrder {
		order, _ := crypto.Decrypt[orderstructs.Order](&encryptedOrder[i], pw)

		orders = append(orders, *order)
	}

	return orders, nil
}

func (o *OrderService) CreateOrder(ctx context.Context, order orderstructs.Order) (uint, error) {

	if order.TenantId.Compare(uuid.Nil()) == 0 {
		return 0, errors.New("TenantId invalid")
	}

	pw, err := o.tenantService.GetPw(ctx, order.TenantId)

	if err != nil {
		return 0, err
	}

	encrypted, ok := crypto.Encrypt[orderstructs.Order](order, pw)

	if !ok {
		return 0, errors.New("Failed to encrypt")
	}

	for i := range order.Products {
		if err := o.productRespository.CreateOrderProduct(ctx, order.Products[i]); err != nil {
			return 0, err
		}
	}

	return o.orderRepository.CreateOrder(ctx, encrypted)
}

func (o *OrderService) UpdateOrder(ctx context.Context, order orderstructs.Order) error {

	if order.OrderId <= 1 {
		return errors.New("OrderId cant be 0 or 1")
	}

	if order.TenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	pw, err := o.tenantService.GetPw(ctx, order.TenantId)

	if err != nil {
		return err
	}

	encrypted, ok := crypto.Encrypt[orderstructs.Order](order, pw)

	if !ok {
		return errors.New("Failed to encrypt product")
	}

	for i := range order.Products {
		if err := o.productRespository.UpdateOrderProduct(ctx, order.Products[i]); err != nil {
			return err
		}
	}

	return o.orderRepository.UpdateOrder(ctx, encrypted)
}

func (o *OrderService) DeleteOrder(ctx context.Context, orderId uint) error {

	if orderId >= 1 {
		return errors.New("OrderId cant be 0 or 1")
	}

	if err := o.productRespository.DeleteOrderProductsByOrderId(ctx, orderId); err != nil {
		return err
	}

	return o.orderRepository.DeleteOrder(ctx, orderId)
}
