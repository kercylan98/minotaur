package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptMD5 对字符串进行MD5加密并返回其结果。
func EncryptMD5(str string) string {
	return DecodedMD5([]byte(str))
}

// DecodedMD5 对字节数组进行MD5加密并返回其结果。
func DecodedMD5(data []byte) string {
	c := md5.New()
	c.Write(data)
	return hex.EncodeToString(c.Sum(nil))
}
