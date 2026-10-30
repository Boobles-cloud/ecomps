package productstructs

import "uuid"

type Product struct {
	ProductId          uint      `json:"ProductId"`
	ProductName        string    `json:"ProductName"`
	ProductPrice       string    `json:"ProductPrice"`
	ProductDescription string    `json:"ProductDescription"`
	TenantId           uuid.UUID `json:"TenantId"`
}
