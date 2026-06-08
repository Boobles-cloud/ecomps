package crypto

func Encrypt[T any](toEncrypt T) (T, bool) {
	return toEncrypt, false
}
