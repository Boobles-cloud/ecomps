package orderstructs

import "boobles.cloud/backend/database"

// This struct is used for the relationship between order and product inside the db
type OrderProduct struct {
	OPId      uint
	ProductId uint
	Amount    uint
	OrderId   uint
}

// Struct used only for creating a order
// So we can directly create all "OrderProduct" entrys in database
type OrderProductCreation struct {
	ProductId uint `json:"ProductId"`
	Amount    uint `json:"Amount"`
}

// Transforms and inserts into database
func (op *OrderProductCreation) InsertIntoDatabase(orderId uint) bool {

	orderProduct := OrderProduct{
		OPId:      0,
		ProductId: op.ProductId,
		Amount:    op.Amount,
		OrderId:   orderId,
	}

	result := database.ExecuteSQLStatement("InsertOrderProduct", database.Insert, []any{orderProduct.ProductId, orderProduct.Amount, orderProduct.OrderId})
	return result.Ok
}
