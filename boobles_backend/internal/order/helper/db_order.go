package helper

import (
	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
	"boobles.cloud/backend/logging"
)

// Gets a order by its id
// Decrypts the given order
func GetOrder(orderId uint, key string) (*orderstructs.Order, bool) {

	order, ok := database.QueryDatabase[orderstructs.Order]("SelectOrderById", []any{orderId})

	if !ok || len(order) != 1 {
		logging.Log(logging.Error, "[Order helper | GetOrder] Failed getting order from db...")
		return nil, false
	}

	orderEncrypted, ok := crypto.Decrypt(&order[0], key)

	if !ok {
		logging.Log(logging.Error, "[Order helper | GetOrder] Failed decrypting")
		return nil, false
	}

	return orderEncrypted, true
}

// TODO: implement getting all orders for tenant id
