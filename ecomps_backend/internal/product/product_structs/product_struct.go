package productstructs

type Product struct {
	ProductId          uint   `json:"ProductId"`
	ProductName        string `json:"ProductName"`
	ProductPrice       string `json:"ProductPrice"`
	ProductDescription string `json:"ProductDescription"`
	TenantId           uint   `json:"TenantId"`
}
