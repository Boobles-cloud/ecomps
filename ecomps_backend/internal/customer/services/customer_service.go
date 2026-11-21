package services

import (
	"context"
	"errors"
	"uuid"

	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/utils/crypto"
)

func (c *CustomerService) GetById(ctx context.Context, id uint) (customerstructs.Customer, error) {

	if id <= 1 {
		return customerstructs.Customer{}, errors.New("Customer Id cant be 0 or 1")
	}

	encryptedCustomer, err := c.customerRepo.GetById(ctx, id)

	if err != nil {
		return encryptedCustomer, err
	}

	pw, err := c.tenantService.GetPw(ctx, encryptedCustomer.TenantId)

	if err != nil {
		return encryptedCustomer, err
	}

	customer, ok := crypto.Decrypt[customerstructs.Customer](&encryptedCustomer, pw)

	if !ok {
		return encryptedCustomer, errors.New("Failed to decrypt")
	}

	return *customer, nil
}

func (c *CustomerService) GetAllByTenantId(ctx context.Context, tenantId uuid.UUID) ([]customerstructs.Customer, error) {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return []customerstructs.Customer{}, errors.New("TenantId invalid")
	}

	encryptedCustomer, err := c.customerRepo.GetAllByTenantId(ctx, tenantId)

	if err != nil {
		return encryptedCustomer, err
	}

	// Check if the customer has the same TenantId
	if tenantId.Compare(encryptedCustomer[0].TenantId) != 0 {
		return encryptedCustomer, errors.New("Unathorized!")
	}

	pw, err := c.tenantService.GetPw(ctx, encryptedCustomer[0].TenantId)

	if err != nil {
		return encryptedCustomer, err
	}

	customers := make([]customerstructs.Customer, 0, len(encryptedCustomer))

	for i := range encryptedCustomer {

		customer, ok := crypto.Decrypt[customerstructs.Customer](&encryptedCustomer[i], pw)

		if ok {
			customers = append(customers, *customer)
		}
	}

	return customers, nil
}

func (c *CustomerService) Create(ctx context.Context, item customerstructs.Customer) (uint, error) {

	if item.TenantId.Compare(uuid.Nil()) == 0 {
		return 0, errors.New("TenantId invalid")
	}

	if item.CustomerName == "" {
		return 0, errors.New("Customer name cant be empty")
	}

	pw, err := c.tenantService.GetPw(ctx, item.TenantId)

	if err != nil {
		return 0, err
	}

	encryptedCustomer, ok := crypto.Encrypt[customerstructs.Customer](item, pw)

	if !ok {
		return 0, errors.New("Failed to encrypt customer")
	}

	return c.customerRepo.Create(ctx, encryptedCustomer)
}

func (c *CustomerService) Update(ctx context.Context, item customerstructs.Customer) error {

	if item.TenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	if item.CustomerId <= 1 {
		return errors.New("Customer Id cant be 0 or 1")
	}

	pw, err := c.tenantService.GetPw(ctx, item.TenantId)

	if err != nil {
		return err
	}

	encrypted, ok := crypto.Encrypt[customerstructs.Customer](item, pw)

	if !ok {
		return errors.New("Failed to encrypt")
	}

	return c.customerRepo.Update(ctx, encrypted, "CustomerId")
}

func (c *CustomerService) Delete(ctx context.Context, id uint, tenantId uuid.UUID) error {

	if tenantId.Compare(uuid.Nil()) == 0 {
		return errors.New("TenantId invalid")
	}

	if id <= 1 {
		return errors.New("Customer Id cant be 0 or 1")
	}

	return c.customerRepo.Delete(ctx, id, tenantId)
}
