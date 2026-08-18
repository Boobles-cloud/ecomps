package customerstructs

import (
	"time"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/utils/crypto"
	"ecomps.boobles.cloud/backend/utils/logging"
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

func (c *Customer) CreateCustomerInDatabase(key string, dh *database.DbHandler) (uint, bool) {

	c.CustomerLastChanged = time.Now()

	encryptedCustomer, ok := crypto.Encrypt[Customer](*c, key)

	if !ok {
		logging.Log(logging.Error, "[Customer | CreateCustomerInDatabase] Failed to encrypt")
		return 0, false
	}

	result := dh.ExecuteSQLStatement("InsertCustomer", []any{encryptedCustomer.CustomerName, encryptedCustomer.CustomerPostalCode,
		encryptedCustomer.CustomerStreetAndHouseNr, encryptedCustomer.CustomerCity, encryptedCustomer.CustomerLastChanged, encryptedCustomer.TenantId})

	return result.LastId, result.Ok
}
