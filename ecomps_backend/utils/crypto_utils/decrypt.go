package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"reflect"

	"ecomps.boobles.cloud/mango/global"
)

// Decrypts a given struct
func Decrypt[T any](toDecrypt *T, key string) (*T, error) {

	valOfT := reflect.ValueOf(toDecrypt).Elem()

	if valOfT.Kind() != reflect.Struct {
		return toDecrypt, global.ERR_INVALID_TYPE
	}

	for i := 0; i < valOfT.NumField(); i++ {

		// Get the field
		field := valOfT.Field(i)

		if field.String() == "" {
			continue
		}

		// Check if the field is a string
		// We only decrypt strings
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			decryptedMsg, err := decryptHelper(field.String(), key)

			if err != nil {
				return toDecrypt, err
			}
			field.SetString(decryptedMsg)
		}
	}

	return toDecrypt, nil
}

// Decryptes the given string with the given key.
// Returns a bool and the decrypted val
func decryptHelper(fieldVal, key string) (string, error) {

	// Masive thanks to: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

	fieldValByte, err := base64.StdEncoding.DecodeString(fieldVal)

	if err != nil {
		return "", err
	}

	// Creates our cipher block
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	// Returns our symetric cipher
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	if len(fieldValByte) < nonceSize {
		return "", global.ERR_INVALID_LENGTH
	}

	nonce, fieldValByte := fieldValByte[:nonceSize], fieldValByte[nonceSize:]

	encryptedFieldVal, err := gcm.Open(nil, nonce, fieldValByte, nil)

	if err != nil {
		return "", err
	}

	return string(encryptedFieldVal), nil
}
