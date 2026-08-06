package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/logging"
)

// Handles deleting a order
func (ho OrderHandler) HandleOrderDeletion(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleOrderDeletion] "+err.Error())
		w.WriteHeader(status)
	}

	orderId, err := strconv.Atoi(r.PathValue("order-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	if result := ho.Dh.ExecuteSQLStatement("DeleteOrderById", []any{orderId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to delete order"))
	}

	ho.Dh.ExecuteSQLStatement("DeleteOrderProductsByOrderId", []any{orderId})

	key := OrderCacheKey + strconv.Itoa(orderId)
	ho.OrderCache.RemoveItem(key)

	w.WriteHeader(http.StatusOK)
}
