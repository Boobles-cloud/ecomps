package services

import (
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/tenant/services"
)

// TODO: Maybe split this into multiple services
type OrderService struct {
	orderRepository    repository.OrderRepository
	productRespository repository.OrderProductRepository
	statusRepository   repository.OrderStatusRepostiory
	tenantService      *services.TenantService
}

func CreateNewOrderService(o repository.OrderRepository, os repository.OrderStatusRepostiory, op repository.OrderProductRepository, t *services.TenantService) *OrderService {
	return &OrderService{
		orderRepository:    o,
		productRespository: op,
		statusRepository:   os,
		tenantService:      t,
	}
}

func OrderToArgs(o orderstructs.Order) []any {
	return []any{
		o.OrderId,
		o.OrderName,
		o.OrderDate,
		o.OrderStatus,
		o.OrderPostalCode,
		o.OrderStreetAndHouseNr,
		o.OrderCity,
		o.OrderLastChanged,
		o.TenantId,
	}
}

func OrderProductsToArgs(o *orderstructs.OrderProduct) []any {
	return []any{
		o.OPId,
		o.ProductId,
		o.Amount,
		o.OrderId,
	}
}

func OrderStatusToArgs(o *orderstructs.OrderStatus) []any {
	return []any{
		o.StatusId,
		o.StatusName,
		o.LanguageId,
	}
}
