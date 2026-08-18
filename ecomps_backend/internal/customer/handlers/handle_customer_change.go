package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"ecomps.boobles.cloud/backend/crypto"
	"ecomps.boobles.cloud/backend/database"
	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Handels changing a customer
func (ch *CustomerHandler) HandleCustomerChange(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleCustomerChange] "+err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var tmpCustomer customerstructs.Customer

	if err := json.Unmarshal(body, &tmpCustomer); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
		return
	}

	encryptedCustomer, ok := crypto.Encrypt(tmpCustomer, tenant.GetPw(ch.Dh, r.Context()))

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed encrypting customer"))
		return
	}

	if !database.UpdateDatabaseEntry[customerstructs.Customer](ch.Dh, "UpdateCustomer", "CustomerId", encryptedCustomer) {
		fail(http.StatusInternalServerError, errors.New("Failed updating customer"))
		return
	}

	go ch.insertItem(tmpCustomer)
	w.WriteHeader(http.StatusOK)
}
