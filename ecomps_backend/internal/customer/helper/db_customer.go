package helper

import (
	"context"

	"ecomps.boobles.cloud/backend/crypto"
	"ecomps.boobles.cloud/backend/database"
	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Gets and decrypts a customer
func GetCustomer(customerId uint, key string, ctx context.Context, dh *database.DbHandler) (customerstructs.Customer, bool) {

	customer, ok := database.QueryOne[customerstructs.Customer](ctx, dh, "SelectCustomerById", customerId)

	if !ok {
		logging.Log(logging.Error, "[Customer | GetCustomer] Failed getting customer from db")
		return customerstructs.Customer{}, false
	}

	decrypted, ok := crypto.Decrypt[customerstructs.Customer](&customer, key)

	if !ok {
		logging.Log(logging.Error, "[Customer | GetCustomer] Failed decrypting customer")
		return customerstructs.Customer{}, false
	}

	return *decrypted, true
}

// Gets and decrypts all customer for a tenant
func GetAllCustomerForTenant(tenantId uint, key string, ctx context.Context, dh *database.DbHandler) ([]customerstructs.Customer, bool) {

	allCustomer, ok := database.QueryMany[customerstructs.Customer](ctx, dh, "SelectAllCustomerByTenantId", tenantId)

	if !ok {
		logging.Log(logging.Error, "[Customer | GetAllCustomerForTenant] Failed getting customer from db")
		return []customerstructs.Customer{}, false
	}

	allDecrypted := make([]customerstructs.Customer, len(allCustomer))

	for i := range allCustomer {

		decrypted, ok := crypto.Decrypt[customerstructs.Customer](&allCustomer[i], key)

		if !ok {
			logging.Log(logging.Error, "[Customer | GetAllCustomerForTenant] Failed getting customer from db")
			continue
		}

		allCustomer = append(allCustomer, *decrypted)
	}

	return allDecrypted, len(allDecrypted) != 0
}
