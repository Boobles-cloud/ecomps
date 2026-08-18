package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Handels creating a customer
func (ch *CustomerHandler) HandleCustomerCreation(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleCustomerCreation]"+err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var customer customerstructs.Customer

	if err := json.Unmarshal(body, &customer); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
		return
	}

	id, ok := customer.CreateCustomerInDatabase(tenant.GetPw(ch.Dh, r.Context()), ch.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed creating customer"))
		return
	}

	customer.CustomerId = id

	go ch.insertItem(customer)
	w.WriteHeader(http.StatusOK)
}
