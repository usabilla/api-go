package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func hash(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func hexHash(data []byte) string {
	return hex.EncodeToString(hash(data))
}

func keyedHash(key, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func hexKeyedHash(key, data []byte) string {
	return hex.EncodeToString(keyedHash(key, data))
}
