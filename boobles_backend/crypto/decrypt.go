package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"reflect"

	"boobles.cloud/backend/logging"
)

// Decrypts a given struct
func Decrypt[T any](toDecrypt *T, key string) (*T, bool) {

	valOfT := reflect.ValueOf(toDecrypt).Elem()

	if valOfT.Kind() != reflect.Struct {
		logging.Log(logging.Error, "[Crypto | Decrypt] T isn't of kind struct!")
		return toDecrypt, false
	}

	for i := 0; i < valOfT.NumField(); i++ {

		// Get the field
		field := valOfT.Field(i)

		// Check if the field is a string
		// We only decrypt strings
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			decryptedMsg, ok := decryptHelper(field.String(), key)

			if !ok {
				logging.Log(logging.Error, "[Crypto | Decrypt] Failed to decrypt: "+field.Type().Name())
				return toDecrypt, false
			} else {
				field.SetString(decryptedMsg)
			}
		}
	}

	return toDecrypt, true
}

// Decryptes the given string with the given key.
// Returns a bool and the decrypted val
func decryptHelper(fieldVal, key string) (string, bool) {

	fieldValByte := []byte(fieldVal)

	// Creates our cipher block
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		logging.Log(logging.Error, "[Crypto | Decrypt] "+err.Error())
		return "", false
	}

	// Returns our symetric
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		logging.Log(logging.Error, "[Crypto | Decrypt] "+err.Error())
		return "", false
	}

	nonceSize := gcm.NonceSize()

	if len(fieldValByte) < nonceSize {
		logging.Log(logging.Error, "[Crypto | Decrypt] Wrong length for nonce size")
		return "", false
	}

	nonce, fieldValByte := fieldValByte[:nonceSize], fieldValByte[nonceSize:]

	encryptedFieldVal, err := gcm.Open(nil, nonce, fieldValByte, nil)

	if err != nil {
		logging.Log(logging.Error, "[Crypto | Decrypt] "+err.Error())
		return "", false
	}

	return string(encryptedFieldVal), true
}
