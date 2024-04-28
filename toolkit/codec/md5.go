package codec

import (
	"crypto/md5"
	"errors"
)

type MD5 struct{}

func (h *MD5) Encode(src []byte) ([]byte, error) {
	hash := md5.New()
	hash.Write(src)
	return append(src, hash.Sum(nil)...), nil
}

func (h *MD5) Decode(src []byte) ([]byte, error) {
	if len(src) < md5.Size {
		return nil, errors.New("invalid data length")
	}
	data := src[:len(src)-md5.Size]
	expectedHash := src[len(src)-md5.Size:]

	hash := md5.New()
	hash.Write(data)
	calculatedHash := hash.Sum(nil)

	for i := 0; i < md5.Size; i++ {
		if expectedHash[i] != calculatedHash[i] {
			return nil, errors.New("MD5 mismatch")
		}
	}

	return data, nil
}
