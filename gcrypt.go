package gcrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func MD5(message string) string {
	hasher := md5.New()
	hasher.Write([]byte(message))
	checksum := hasher.Sum(nil)
	return hex.EncodeToString(checksum)
}

func SHA1(message string) string {
	hasher := sha1.New()
	hasher.Write([]byte(message))
	checksum := hasher.Sum(nil)
	return hex.EncodeToString(checksum)
}

func SHA256(message string) string {
	hasher := sha256.New()
	hasher.Write([]byte(message))
	checksum := hasher.Sum(nil)
	return hex.EncodeToString(checksum)
}

func Base64Enc(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func Base64Dec(payload string) string {
	r, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}
	return string(r[:])
}

func HmacSHA256(message string, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(message))
	checksum := hasher.Sum(nil)
	return hex.EncodeToString(checksum)
}
