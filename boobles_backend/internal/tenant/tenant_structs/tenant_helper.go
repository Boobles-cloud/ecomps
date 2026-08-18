package tenantstructs

import (
	"context"
	"database/sql"
	"encoding/base64"
	"math/rand"
	"time"

	"boobles.cloud/backend/database"
)

const (
	allCharacters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	maxLengthOfKey = 32
)

// Creates the master key for the given tenant, inside the given transaction.
func createMasterKey(ctx context.Context, tx *sql.Tx, dh *database.DbHandler, t Tenant) (uint, bool) {
	tenantPw := t.TenantName + createRandomString(maxLengthOfKey)
	tenantPwBase64 := base64.StdEncoding.EncodeToString([]byte(tenantPw))

	result := dh.ExecuteSQLStatementTx(ctx, tx, "InsertTenantPw", []any{tenantPwBase64})
	return result.LastId, result.Ok
}

// Create a random string of the given length.
func createRandomString(length int) string {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	result := make([]byte, length)
	for i := range length {
		result[i] = allCharacters[random.Intn(len(allCharacters))]
	}
	return string(result)
}
