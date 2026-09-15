package services

import (
	"context"
	"errors"

	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
)

func (s *StatusService) GetById(ctx context.Context, id, langId uint) (orderstructs.OrderStatus, error) {

	if id == 0 {
		return orderstructs.OrderStatus{}, errors.New("Id cant be 0")
	}

	if langId == 0 {
		return orderstructs.OrderStatus{}, errors.New("LanguageId cant be 0")
	}

	return s.statusRepository.GetById(ctx, id, langId)
}

func (s *StatusService) GetAllByLangId(ctx context.Context, langId uint) ([]orderstructs.OrderStatus, error) {

	if langId == 0 {
		return []orderstructs.OrderStatus{}, errors.New("LanguageId cant be 0")
	}

	return s.GetAllByLangId(ctx, langId)
}
