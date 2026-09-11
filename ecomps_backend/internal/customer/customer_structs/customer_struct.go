package customerstructs

import (
	"time"
)

type Customer struct {
	CustomerId               uint      `json:"CustomerId"`
	CustomerName             string    `json:"CustomerName"`
	CustomerPostalCode       string    `json:"CustomerPostalCode"`
	CustomerStreetAndHouseNr string    `json:"CustomerStreetAndHouseNr"`
	CustomerCity             string    `json:"CustomerCity"`
	CustomerLastChanged      time.Time `json:"CustomerLastChanged"`
	TenantId                 uint      `json:"TenantId"`
}
