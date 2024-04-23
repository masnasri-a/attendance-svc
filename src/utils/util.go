package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// comparePassword is a function that compares the password
func ComparePassword(hashedPassword string, password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)) == hashedPassword
}

func GenerateUUID() string {
	return primitive.NewObjectID().Hex()
}

func GetTimeNow() int64 {
	return time.Now().Unix()
}
