package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateID() string {
	b := make([]byte, 16) // 16 bytes = 128 bits
	rand.Read(b)
	return hex.EncodeToString(b)
}
