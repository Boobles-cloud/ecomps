package orderstructs

import (
	"context"
	"time"

	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
)

type Order struct {
	OrderId               uint           `json:"OrderId"`
	OrderName             string         `json:"OrderName"`
	OrderDate             time.Time      `json:"OderDate"`
	OrderStatus           uint           `json:"OrderStatus"`
	OrderPostalCode       string         `json:"OrderPostalCode"`
	OrderStreetAndHouseNr string         `json:"OrderStreetAndHouseNr"`
	OrderCity             string         `json:"OrderCity"`
	OrderLastChanged      time.Time      `json:"OrderLastChanged"`
	Products              []OrderProduct `json:"Products,omitempty"`
}

// Encrypts the order and creates a order in database
func (o *Order) CreateOrderInDatabase(key string, dh *database.DbHandler) (uint, bool) {

	o.OrderLastChanged = time.Now()
	order, ok := crypto.Encrypt(o, key)

	if !ok {
		return 0, false
	}

	result := dh.ExecuteSQLStatement("InsertOrder", []any{order.OrderId, order.OrderName, order.OrderDate, order.OrderStatus, order.OrderPostalCode, order.OrderStreetAndHouseNr, order.OrderCity, order.OrderLastChanged})

	return result.LastId, result.Ok
}

// Gets all product ids and amount for a order
func (o *Order) GetAllProducts(ctx context.Context, dh *database.DbHandler) {
	allProducts, _ := database.QueryMany[OrderProduct](ctx, dh, "SelectOrderProductsByOrderId", []any{o.OrderId})
	o.Products = allProducts
}
