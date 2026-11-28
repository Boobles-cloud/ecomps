package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"reflect"

	"ecomps.boobles.cloud/mango/global"
)

func Encrypt[T any](toEncrypt T, key string) (T, error) {

	valOfT := reflect.ValueOf(toEncrypt).Elem()

	if valOfT.Kind() != reflect.Struct {
		return toEncrypt, global.ERR_INVALID_TYPE
	}

	for i := 0; i < valOfT.NumField(); i++ {

		// Get the field
		field := valOfT.Field(i)

		// Check if the field is a string
		// We only decrypt strings
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			decryptedMsg, err := encryptionHelper(field.String(), key)
			if err != nil {
				return toEncrypt, err
			} else {
				field.SetString(decryptedMsg)
			}
		}
	}

	return toEncrypt, nil
}

// Encryptes the given value with the given key.
// Returns the encrypted value and a bool
func encryptionHelper(fieldVal, key string) (string, error) {

	// Masive thanks to: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

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

	nonce := make([]byte, gcm.NonceSize())

	// Populates the nonce with random secure sequences
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encryptes the given byte array
	encrypted := gcm.Seal(nonce, nonce, []byte(fieldVal), nil)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
