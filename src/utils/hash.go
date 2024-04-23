package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword is a function that hashes the password
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
