package helper

import (
	"context"

	"ecomps.boobles.cloud/backend/crypto"
	"ecomps.boobles.cloud/backend/database"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Gets a order by its id
// Decrypts the given order
func GetOrder(orderId uint, key string, dh *database.DbHandler, ctx context.Context) (*orderstructs.Order, bool) {

	order, ok := database.QueryOne[orderstructs.Order](ctx, dh, "SelectOrderById", orderId)

	if !ok {
		logging.Log(logging.Error, "[Order helper | GetOrder] Failed getting order from db...")
		return nil, false
	}

	orderEncrypted, ok := crypto.Decrypt(&order, key)

	if !ok {
		logging.Log(logging.Error, "[Order helper | GetOrder] Failed decrypting")
		return nil, false
	}

	return orderEncrypted, true
}

// Gets and decrypts all orders
func GetAllOrders(tenantId uint, key string, dh *database.DbHandler, ctx context.Context) ([]orderstructs.Order, bool) {
	orders, ok := database.QueryMany[orderstructs.Order](ctx, dh, "SelectOrdersByTenantId", tenantId)

	if !ok {
		logging.Log(logging.Error, "[Order helper | GetAllOrders] Failed getting order from db...")
		return nil, false
	}

	allDecryptedOrders := make([]orderstructs.Order, len(orders))

	for i := range orders {
		o, ok := crypto.Decrypt[orderstructs.Order](&orders[i], key)

		if !ok {
			logging.Log(logging.Error, "[Order helper | GetAllOrders] Failed decrypting order...")
			continue
		}

		allDecryptedOrders = append(allDecryptedOrders, *o)
	}

	return allDecryptedOrders, len(allDecryptedOrders) != 0
}
