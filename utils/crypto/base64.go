package kcrypto

import "encoding/base64"

// EncryptBase64 对数据进行Base64编码
func EncryptBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodedBase64 对数据进行Base64解码
func DecodedBase64(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	return decoded, err
}
