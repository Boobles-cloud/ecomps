package handlers

import (
	"context"
	"net/http"
	"time"

	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels changing a customer
func (c *CustomerHandler) HandleCustomerChange(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Customer | HandleCustomerChange")

	customer, err := jsonutils.JsonDeserilizeHttpRequestBody[customerstructs.Customer](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := c.customerService.Update(ctx, customer); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go c.insertItem(customer)
	w.WriteHeader(http.StatusOK)
}
