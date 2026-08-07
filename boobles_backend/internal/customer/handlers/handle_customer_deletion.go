package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/logging"
)

// Handles the deletion of a customer
func (ch *CustomerHandler) HandleCustomerDeletion(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleCustomerDeletion] "+err.Error())
		w.WriteHeader(status)
	}

	customerId, err := strconv.Atoi(r.PathValue("customer-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	if result := ch.Dh.ExecuteSQLStatement("DeleteCustomerById", []any{customerId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to delete database"))
	}

	key := CustomerCacheKey + strconv.Itoa(customerId)
	ch.CustomerCache.RemoveItem(key)
	w.WriteHeader(http.StatusOK)
}
