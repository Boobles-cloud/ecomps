package orderstructs

import "ecomps.boobles.cloud/backend/database"

// This struct is used for the relationship between order and product inside the db
type OrderProduct struct {
	OPId      uint `json:"-"`
	ProductId uint `json:"ProductId"`
	Amount    uint `json:"Amount"`
	OrderId   uint `json:"-"`
}

// Inserts the product in the databas
func (op *OrderProduct) InsertIntoDatabase(orderId uint, dh *database.DbHandler) bool {

	result := dh.ExecuteSQLStatement("InsertOrderProduct", []any{op.ProductId, op.Amount, op.OrderId})
	return result.Ok
}

func (op *OrderProduct) UpdateOrderProduct(dh *database.DbHandler) {
	database.UpdateDatabaseEntry[OrderProduct](dh, "UpdateOrderProduct", "OPId", *op)
}
