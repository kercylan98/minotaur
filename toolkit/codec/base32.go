package codec

import (
	"encoding/base32"
	"errors"
)

var base32Encoding = base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")

type Base32 struct{}

func (b *Base32) Encode(src []byte) ([]byte, error) {
	encoded := base32Encoding.EncodeToString(src)
	return []byte(encoded), nil
}

func (b *Base32) Decode(src []byte) ([]byte, error) {
	decoded, err := base32Encoding.DecodeString(string(src))
	if err != nil {
		return nil, errors.New("invalid Base32 data")
	}
	return decoded, nil
}
