package datasheet

import "math/big"

type BigInt big.Int

func (b *BigInt) MarshalJSON() ([]byte, error) {
	return (*big.Int)(b).MarshalText()
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	return (*big.Int)(b).UnmarshalText(data)
}

func (b *BigInt) Value() *big.Int {
	return (*big.Int)(b)
}

type BigFloat big.Float

func (b *BigFloat) MarshalJSON() ([]byte, error) {
	return (*big.Float)(b).MarshalText()
}

func (b *BigFloat) UnmarshalJSON(data []byte) error {
	return (*big.Float)(b).UnmarshalText(data)
}

func (b *BigFloat) Value() *big.Float {
	return (*big.Float)(b)
}
