package gcrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))

	a := hasher.Sum(nil)
	return hex.EncodeToString(a)
}
