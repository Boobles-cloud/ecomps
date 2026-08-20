package handlers

import (
	"errors"
	"net/http"
	"strconv"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handles deleting a order
func (ho OrderHandler) HandleOrderDeletion(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Order | HandleOrderDeletion")

	orderId, err := httputils.IntPathParam(r, "order_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if result := ho.Dh.ExecuteSQLStatement("DeleteOrderById", []any{orderId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed to delete order"))
		return
	}

	ho.Dh.ExecuteSQLStatement("DeleteOrderProductsByOrderId", []any{orderId})

	key := OrderCacheKey + strconv.Itoa(orderId)
	ho.OrderCache.RemoveItem(key)

	w.WriteHeader(http.StatusOK)
}
