package gcrypt

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SHA256(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))

	a := hasher.Sum(nil)
	return hex.EncodeToString(a)
}

func Base64Enc(p string) string {
	return base64.StdEncoding.EncodeToString([]byte(p))
}

func Base64Dec(p string) string {
	return base64.StdEncoding.DecodeString([]byte(p))
}
