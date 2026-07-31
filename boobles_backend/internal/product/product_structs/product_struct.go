package productstructs

import (
	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

type Product struct {
	ProductId          uint   `json:"ProductId"`
	ProductName        string `json:"ProductName"`
	ProductPrice       string `json:"ProductPrice"`
	ProductDescription string `json:"ProductDescription"`
	ProductPicturePath string `json:"-"` // TODO: set this path as a picture is added
	TenantId           uint   `json:"TenantId"`
}

// Encrypts a product and stores it in database
func (p *Product) CreateProductInDatabase(key string) (uint, bool) {

	product, ok := crypto.Encrypt(p, key)

	if !ok {
		logging.Log(logging.Error, "[Product | CreateProductInDatabase] Failed to encrypt product...")
		return 0, false
	}

	if res := database.ExecuteSQLStatement("InserProduct", database.Insert, []any{product.ProductId, product.ProductName,
		product.ProductPrice, product.ProductPrice, product.ProductDescription, product.ProductPrice, product.TenantId}); res.Ok {
		return res.LastId, true
	}

	logging.Log(logging.Error, "[Product | CreateProductInDatabase] Failed to create product in database...")
	return 0, false
}
