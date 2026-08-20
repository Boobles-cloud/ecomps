package handlers

import (
	"errors"
	"net/http"
	"strconv"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handles the deletion of a customer
func (ch *CustomerHandler) HandleCustomerDeletion(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Customer | HandleCustomerDeletion")

	customerId, err := httputils.IntPathParam(r, "customer-id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if result := ch.Dh.ExecuteSQLStatement("DeleteCustomerById", []any{customerId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to delete database"))
		return
	}

	key := CustomerCacheKey + strconv.Itoa(customerId)
	ch.CustomerCache.RemoveItem(key)
	w.WriteHeader(http.StatusOK)
}
