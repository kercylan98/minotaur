package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// EncryptSHA256 对字符串进行SHA256加密并返回其结果。
func EncryptSHA256(str string) string {
	return DecodedSHA256([]byte(str))
}

// DecodedSHA256 对字节数组进行SHA256加密并返回其结果。
func DecodedSHA256(data []byte) string {
	bytes := sha256.Sum256(data)
	return hex.EncodeToString(bytes[:])
}
