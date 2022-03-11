package jwt

var secretKey []byte

// SetSecretKey sets the secret key used for hashing JWT.
// It must be used before any usage of this package.
func SetSecretKey(key []byte) {
	secretKey = key
}

// key is a safe way to access secretKey. It panics at runtime
// if secretKey is not set.
func key() []byte {
	if len(secretKey) == 0 {
		panic("secretKey must be set before the use of package jwt")
	}
	return secretKey
}
