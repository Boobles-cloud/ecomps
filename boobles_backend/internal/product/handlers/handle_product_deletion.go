package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

func (p *ProductHandler) HandleDeletingProduct(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	productId, err := strconv.Atoi(r.PathValue("product-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	if result := database.ExecuteSQLStatement("DeleteProductById", database.Delete, []any{productId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed deleting product"))
	}

	w.WriteHeader(http.StatusOK)
}
