package gcrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func MD5(m string) string {
	hasher := md5.New()
	hasher.Write([]byte(m))
	c := hasher.Sum(nil)
	return hex.EncodeToString(c)
}

func SHA1(m string) string {
	hasher := sha1.New()
	hasher.Write([]byte(m))
	c := hasher.Sum(nil)
	return hex.EncodeToString(c)
}

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
	r, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		panic(err)
	}
	return string(r[:])
}

func HmacSHA256(m string, k string) string {
	hasher := hmac.New(sha256.New, []byte(k))
	hasher.Write([]byte(m))
	a := hasher.Sum(nil)
	return hex.EncodeToString(a)
}
