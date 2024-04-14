package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// comparePassword is a function that compares the password
func ComparePassword(hashedPassword string, password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)) == hashedPassword
}