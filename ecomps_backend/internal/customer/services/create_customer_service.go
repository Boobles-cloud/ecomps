package services

import (
	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/internal/repository"
	"ecomps.boobles.cloud/backend/internal/tenant/services"
)

type CustomerService struct {
	customerRepo  repository.Repository[customerstructs.Customer]
	tenantService *services.TenantService
}

func CreateNewCustomerService(r repository.Repository[customerstructs.Customer], t *services.TenantService) *CustomerService {
	return &CustomerService{
		customerRepo:  r,
		tenantService: t,
	}
}

func CustomerToArgs(c customerstructs.Customer) []any {
	return []any{
		c.CustomerId,
		c.CustomerName,
		c.CustomerPostalCode,
		c.CustomerStreetAndHouseNr,
		c.CustomerCity,
		c.CustomerLastChanged,
		c.TenantId,
	}
}
