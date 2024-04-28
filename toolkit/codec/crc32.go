package codec

import (
	"errors"
	"hash/crc32"
)

type Crc32 struct{}

func (c *Crc32) Encode(src []byte) ([]byte, error) {
	crcTable := crc32.MakeTable(crc32.IEEE)
	crcValue := crc32.Checksum(src, crcTable)

	// 将 CRC 值附加到原始数据的末尾
	result := append(src, byte(crcValue>>24), byte(crcValue>>16), byte(crcValue>>8), byte(crcValue))
	return result, nil
}

func (c *Crc32) Decode(src []byte) ([]byte, error) {
	if len(src) < 4 {
		return nil, errors.New("invalid data")
	}

	// 获取数据部分和 CRC 校验码部分
	data := src[:len(src)-4]
	crcReceived := uint32(src[len(src)-4])<<24 | uint32(src[len(src)-3])<<16 | uint32(src[len(src)-2])<<8 | uint32(src[len(src)-1])

	// 计算数据的 CRC
	crcTable := crc32.MakeTable(crc32.IEEE)
	crcCalculated := crc32.Checksum(data, crcTable)

	// 验证 CRC 是否匹配
	if crcReceived != crcCalculated {
		return nil, errors.New("CRC check failed")
	}

	return data, nil
}
