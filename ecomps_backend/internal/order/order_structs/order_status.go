package orderstructs

type OrderStatus struct {
	StatusId   uint   `json:"StatusId"`
	StatusName string `json:"StatusName"`
	LanguageId uint   `json:"LanguageId"`
}
