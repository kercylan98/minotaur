package codec

import "encoding/base64"

type Base64 struct{}

func (b *Base64) Encode(src []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(src)), nil
}

func (b *Base64) Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
