package orderstructs

import (
	"context"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

type OrderStatus struct {
	StatusId   uint   `json:"StatusId"`
	StatusName string `json:"StatusName"`
	LanguageId uint   `json:"LanguageId"`
}

func (o *OrderStatus) GetStatusNameInCurrentLang(dh *database.DbHandler, ctx context.Context) string {

	status, ok := database.QueryOne[OrderStatus](ctx, dh, "SelectOrderStatusById", o.StatusId)

	if !ok {
		logging.Log(logging.Error, "[OrderStatus | GetStatusNameInCurrentLang] Failed to get status")
		return ""
	}

	return status.StatusName
}
