package codec

import (
	"crypto/sha1"
	"errors"
)

type Sha1 struct{}

func (h *Sha1) Encode(src []byte) ([]byte, error) {
	hash := sha1.New()
	hash.Write(src)
	return append(src, hash.Sum(nil)...), nil
}

func (h *Sha1) Decode(src []byte) ([]byte, error) {
	if len(src) < sha1.Size {
		return nil, errors.New("invalid data length")
	}

	data := src[:len(src)-sha1.Size]
	expectedHash := src[len(src)-sha1.Size:]

	hash := sha1.New()
	hash.Write(data)
	calculatedHash := hash.Sum(nil)

	for i := 0; i < sha1.Size; i++ {
		if expectedHash[i] != calculatedHash[i] {
			return nil, errors.New("SHA-1 mismatch")
		}
	}

	return data, nil
}
