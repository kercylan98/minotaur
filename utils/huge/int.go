package huge

import "math/big"

type Int big.Int

func NewInt[T uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64](x T, exp ...T) *Int {
	num := int64(x)
	i := big.NewInt(num)
	for _, t := range exp {
		i = i.Exp(i, big.NewInt(int64(t)), nil)
	}
	return (*Int)(i)
}
