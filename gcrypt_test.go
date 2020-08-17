package gcrypt

import (
	"testing"
)

const (
	md5TestResult        = "098f6bcd4621d373cade4e832627b4f6"
	sha1TestResult       = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	sha256TestResult     = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	base64EncTestResult  = "dGVzdA=="
	base64DecTestResult  = "test"
	hmacSHA256TestResult = "88cd2108b5347d973cf39cdf9053d7dd42704876d8c9a9bd8e2d168259d3ddf7"
)

func TestMD5(t *testing.T) {
	ans := MD5("test")
	if ans != md5TestResult {
		t.Errorf("MD5(\"test\") = %s; want %s", ans, md5TestResult)
	}
}

func TestSHA1(t *testing.T) {
	ans := SHA1("test")
	if ans != sha1TestResult {
		t.Errorf("SHA1(\"test\") = %s; want %s", ans, sha1TestResult)
	}

}

func TestSHA256(t *testing.T) {
	ans := SHA256("test")
	if ans != sha256TestResult {
		t.Errorf("SHA256(\"test\") = %s; want %s", ans, sha256TestResult)
	}
}

func TestBase64Enc(t *testing.T) {
	ans := Base64Enc("test")
	if ans != base64EncTestResult {
		t.Errorf("Base64Enc(\"test\") = %s; want %s", ans, base64EncTestResult)
	}
}

func TestBase64Dec(t *testing.T) {
	ans := Base64Dec("dGVzdA==")
	if ans != base64DecTestResult {
		t.Errorf("Base64Dec(\"dGVzdA==\") = %s; want %s", ans, base64DecTestResult)
	}
}

func TestHmacSHA256(t *testing.T) {
	ans := HmacSHA256("test", "test")
	if ans != hmacSHA256TestResult {
		t.Errorf("HmacSHA256(\"test\", \"test\") = %s; want %s", ans, hmacSHA256TestResult)
	}
}
