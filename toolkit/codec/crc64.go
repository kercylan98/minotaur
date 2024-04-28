package codec

import (
	"errors"
	"hash/crc64"
)

type Crc64 struct{}

func (c *Crc64) Encode(src []byte) ([]byte, error) {
	crcValue := crc64.Checksum(src, crc64.MakeTable(crc64.ISO))

	// 将 CRC-64 值附加到原始数据的末尾（CRC-64 是 8 字节）
	result := append(src,
		byte(crcValue>>56), byte(crcValue>>48), byte(crcValue>>40), byte(crcValue>>32),
		byte(crcValue>>24), byte(crcValue>>16), byte(crcValue>>8), byte(crcValue),
	)
	return result, nil
}

func (c *Crc64) Decode(src []byte) ([]byte, error) {
	if len(src) < 8 {
		return nil, errors.New("invalid data")
	}

	// 获取数据部分和 CRC 校验码部分
	data := src[:len(src)-8]
	crcReceived := uint64(0)

	// 将 CRC 校验码从字节数组转换为 uint64
	for i := 0; i < 8; i++ {
		crcReceived |= uint64(src[len(src)-8+i]) << (56 - (i * 8))
	}

	// 计算数据的 CRC
	crcCalculated := crc64.Checksum(data, crc64.MakeTable(crc64.ISO))

	// 验证 CRC 是否匹配
	if crcReceived != crcCalculated {
		return nil, errors.New("CRC check failed")
	}

	return data, nil
}
