package orderstructs

// This struct is used for the relationship between order and product inside the db
type OrderProduct struct {
	OPId      uint `json:"-"`
	ProductId uint `json:"ProductId"`
	Amount    uint `json:"Amount"`
	OrderId   uint `json:"-"`
}
