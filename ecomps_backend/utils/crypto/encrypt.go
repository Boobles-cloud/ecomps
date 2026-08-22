package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"reflect"

	"ecomps.boobles.cloud/backend/utils/logging"
)

func Encrypt[T any](toEncrypt T, key string) (T, bool) {

	valOfT := reflect.ValueOf(toEncrypt).Elem()

	if valOfT.Kind() != reflect.Struct {
		logging.Log(logging.Error, "[Crypto | Encrypt] T isn't of kind struct!")
		return toEncrypt, false
	}

	for i := 0; i < valOfT.NumField(); i++ {

		// Get the field
		field := valOfT.Field(i)

		// Check if the field is a string
		// We only decrypt strings
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			decryptedMsg, ok := encryptionHelper(field.String(), key)

			if !ok {
				logging.Log(logging.Error, "[Crypto | Encrypt] Failed to encrypt: "+field.Type().Name())
				return toEncrypt, false
			} else {
				field.SetString(decryptedMsg)
			}
		}
	}

	return toEncrypt, true
}

// Encryptes the given value with the given key.
// Returns the encrypted value and a bool
func encryptionHelper(fieldVal, key string) (string, bool) {

	// Masive thanks to: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

	// Creates our cipher block
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		logging.Log(logging.Error, "[Crypto | Encrypt] "+err.Error())
		return "", false
	}

	// Returns our symetric cipher
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		logging.Log(logging.Error, "[Crypto | Encrypt] "+err.Error())
		return "", false
	}

	nonce := make([]byte, gcm.NonceSize())

	// Populates the nonce with random secure sequences
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// Encryptes the given byte array
	encrypted := gcm.Seal(nonce, nonce, []byte(fieldVal), nil)

	return base64.StdEncoding.EncodeToString(encrypted), true
}
