package orderstructs

import (
	"time"
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
	TenantId              uint           `json:"TenantId"`
}
