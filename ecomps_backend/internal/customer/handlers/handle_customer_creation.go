package handlers

import (
	"context"
	"net/http"
	"time"

	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels creating a customer
func (c *CustomerHandler) HandleCustomerCreation(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Customer | HandleCustomerCreation")

	customer, err := jsonutils.JsonDeserilizeHttpRequestBody[customerstructs.Customer](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := ctx.Value(middleware.TenantIdContextKey).(int)

	customer.TenantId = uint(tenantId)

	cId, err := c.customerService.Create(ctx, customer)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	customer.CustomerId = cId

	go c.insertItem(customer)
	w.WriteHeader(http.StatusOK)
}
