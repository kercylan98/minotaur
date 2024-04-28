package codec

import (
	"crypto/sha256"
	"errors"
)

type Sha256 struct{}

func (s *Sha256) Encode(src []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(src)
	return append(src, hash.Sum(nil)...), nil
}

func (s *Sha256) Decode(src []byte) ([]byte, error) {
	if len(src) < sha256.Size {
		return nil, errors.New("invalid data length")
	}

	data := src[:len(src)-sha256.Size]
	expectedHash := src[len(src)-sha256.Size:]

	hash := sha256.New()
	hash.Write(data)
	calculatedHash := hash.Sum(nil)

	for i := 0; i < sha256.Size; i++ {
		if expectedHash[i] != calculatedHash[i] {
			return nil, errors.New("SHA-256 mismatch")
		}
	}

	return data, nil
}
