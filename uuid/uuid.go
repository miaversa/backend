package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

func New() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes[:])
}
