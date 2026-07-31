package orderstructs

// This struct is used for the relationship between order and product inside the db
type OrderProduct struct {
	OPId      uint
	ProductId uint
	OrderId   uint
}
