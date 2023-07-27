package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

// EncryptSHA1 对字符串进行SHA1加密并返回其结果。
func EncryptSHA1(str string) string {
	return DecodedSHA1([]byte(str))
}

// DecodedSHA1 对字节数组进行SHA1加密并返回其结果。
func DecodedSHA1(data []byte) string {
	c := sha1.New()
	c.Write(data)
	return hex.EncodeToString(c.Sum(nil))
}
