package productstructs

import (
	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/utils/crypto"
	"ecomps.boobles.cloud/backend/utils/logging"
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
func (p *Product) CreateProductInDatabase(key string, dh *database.DbHandler) (uint, bool) {

	product, ok := crypto.Encrypt(p, key)

	if !ok {
		logging.Log(logging.Error, "[Product | CreateProductInDatabase] Failed to encrypt product...")
		return 0, false
	}

	if res := dh.ExecuteSQLStatement("InserProduct", []any{product.ProductId, product.ProductName,
		product.ProductPrice, product.ProductPrice, product.ProductDescription, product.ProductPrice, product.TenantId}); res.Ok {
		return res.LastId, true
	}

	logging.Log(logging.Error, "[Product | CreateProductInDatabase] Failed to create product in database...")
	return 0, false
}
