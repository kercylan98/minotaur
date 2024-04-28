package codec

import (
	"errors"
	"math/big"
)

const (
	base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	base58Base     = 58
)

type Base58 struct{}

func (b *Base58) Encode(src []byte) ([]byte, error) {
	x := big.NewInt(0).SetBytes(src)
	mod := big.NewInt(base58Base)
	var encoded []byte

	for x.Sign() > 0 {
		remainder := big.NewInt(0)
		x.DivMod(x, mod, remainder)
		encoded = append([]byte{base58Alphabet[remainder.Int64()]}, encoded...)
	}

	// 添加前导零以保持相同的长度
	for _, v := range src {
		if v != 0 {
			break
		}
		encoded = append([]byte{base58Alphabet[0]}, encoded...)
	}

	return encoded, nil
}

func (b *Base58) Decode(src []byte) ([]byte, error) {
	x := big.NewInt(0)
	base := big.NewInt(base58Base)

	for _, c := range src {
		index := b.alphabetIndex(c)
		if index == -1 {
			return nil, errors.New("invalid Base58 character")
		}
		x.Mul(x, base)
		x.Add(x, big.NewInt(index))
	}

	return x.Bytes(), nil
}

func (b *Base58) alphabetIndex(char byte) int64 {
	for i, a := range []byte(base58Alphabet) {
		if a == char {
			return int64(i)
		}
	}
	return -1
}
