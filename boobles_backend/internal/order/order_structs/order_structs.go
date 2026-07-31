package orderstructs

import (
	"time"

	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
)

type Order struct {
	OrderId               uint      `json:"OrderId"`
	OrderName             string    `json:"OrderName"`
	OrderDate             time.Time `json:"OderDate"`
	OrderStatus           uint      `json:"OrderStatus"`
	OrderPostalCode       string    `json:"OrderPostalCode"`
	OrderStreetAndHouseNr string    `json:"OrderStreetAndHouseNr"`
	OrderCity             string    `json:"OrderCity"`
	OrderLastChanged      time.Time `json:"OrderLastChanged"`
}

// Encrypts the order and creates a order in database
func (o *Order) CreateOrderInDatabase(key string) (uint, bool) {

	order, ok := crypto.Encrypt(o, key)

	if !ok {
		return 0, false
	}

	result := database.ExecuteSQLStatement("InsertOrder", database.Insert, []any{order.OrderId, order.OrderName, order.OrderDate, order.OrderStatus, order.OrderPostalCode, order.OrderStreetAndHouseNr, order.OrderCity, order.OrderLastChanged})

	return result.LastId, result.Ok
}

func (o *Order) GetOrderStatus() string

// func (o *Order) GetAllProducts() []productstructs.Product

func (o *Order) GetCurrentStatusAsString() string
