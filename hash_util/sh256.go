package hash_util

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(input []byte) []byte {
	hash := sha256.Sum256(input)
	return hash[:]
}

func SHA256String(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
