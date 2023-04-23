package kcrypto

import (
	"hash/crc32"
)

// EncryptCRC32 对字符串进行CRC加密并返回其结果。
func EncryptCRC32(str string) uint32 {
	return DecodedCRC32([]byte(str))
}

// DecodedCRC32 对字节数组进行CRC加密并返回其结果。
func DecodedCRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
