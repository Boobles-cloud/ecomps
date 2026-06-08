package crypto

import (
	"encoding/base64"
	"math/rand"
	"time"

	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
)

const (
	allCharacters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	maxLengthOfKey = 32
)

// TODO: Rethink this -> is there a better way to do this?

// Creates the master key for the given tenant
func CreateMasterKey(t tenantstructs.Tenant) (string, bool) {

	tenantPw := t.TenantName
	tenantPw += createRandomString(len(t.TenantName) - maxLengthOfKey)

	tenantPwBase64 := base64.StdEncoding.EncodeToString([]byte(tenantPw))

	return tenantPwBase64, true
}

// Create a random string
func createRandomString(length int) string {

	// Creates our random source
	randSource := rand.NewSource(time.Now().Unix())
	random := rand.New(randSource)

	result := make([]byte, maxLengthOfKey/2)

	for i := range length {
		result[i] = allCharacters[random.Intn(len(allCharacters))]
	}

	return string(result)
}
