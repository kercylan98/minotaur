package huge

import "math/big"

var (
	IntNegativeOne = NewInt(-1)    // 默认初始化的-1值Int，应当将其当作常量使用
	IntZero        = NewInt(0)     // 默认初始化的0值Int，应当将其当作常量使用
	IntOne         = NewInt(1)     // 默认初始化的1值Int，应当将其当作常量使用
	IntTen         = NewInt(10)    // 默认初始化的10值Int，应当将其当作常量使用
	IntHundred     = NewInt(100)   // 默认初始化的100值Int，应当将其当作常量使用
	IntThousand    = NewInt(1000)  // 默认初始化的1000值Int，应当将其当作常量使用
	IntTenThousand = NewInt(10000) // 默认初始化的10000值Int，应当将其当作常量使用
)

type IntRestrain interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type Int big.Int

func NewInt[T IntRestrain](x T, exp ...T) *Int {
	num := int64(x)
	i := big.NewInt(num)
	for _, t := range exp {
		i = i.Exp(i, big.NewInt(int64(t)), nil)
	}
	return (*Int)(i)
}

func (slf *Int) Copy() *Int {
	return (*Int)(new(big.Int).Set(slf.ToBigint()))
}

func (slf *Int) Set(i *Int) *Int {
	return (*Int)(slf.ToBigint().Set(i.ToBigint()))
}

func (slf *Int) SetInt(i int) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetInt8(i int8) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetInt16(i int16) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetInt32(i int32) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetInt64(i int64) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetUint(i uint) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetUint8(i uint8) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetUint16(i uint16) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetUint32(i uint32) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) SetUint64(i uint64) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

func (slf *Int) IsZero() bool {
	if slf == nil || slf.EqualTo(IntZero) {
		return true
	}
	return false
}

func (slf *Int) ToBigint() *big.Int {
	return (*big.Int)(slf)
}

func (slf *Int) Cmp(i *Int) int {
	return slf.ToBigint().Cmp(i.ToBigint())
}

// GreaterThan 大于
func (slf *Int) GreaterThan(i *Int) bool {
	return slf.Cmp(i) > 0
}

// GreaterThanOrEqualTo 大于或等于
func (slf *Int) GreaterThanOrEqualTo(i *Int) bool {
	return slf.Cmp(i) >= 0
}

// LessThan 小于
func (slf *Int) LessThan(i *Int) bool {
	return slf.Cmp(i) < 0
}

// LessThanOrEqualTo 小于或等于
func (slf *Int) LessThanOrEqualTo(i *Int) bool {
	return slf.Cmp(i) <= 0
}

// EqualTo 等于
func (slf *Int) EqualTo(i *Int) bool {
	return slf.Cmp(i) == 0
}

func (slf *Int) Int64() int64 {
	return slf.ToBigint().Int64()
}

func (slf *Int) String() string {
	return slf.ToBigint().String()
}

func (slf *Int) Add(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, i.ToBigint()))
}

func (slf *Int) AddInt(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddInt8(i int8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddInt16(i int16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddInt32(i int32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddInt64(i int64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddUint(i uint) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddUint8(i uint8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddUint16(i uint16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddUint32(i uint32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) AddUint64(i uint64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Add(x, NewInt(i).ToBigint()))
}

func (slf *Int) Mul(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, i.ToBigint()))
}

func (slf *Int) MulInt(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulInt8(i int8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulInt16(i int16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulInt32(i int32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulInt64(i int64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulUint(i uint) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulUint8(i uint8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulUint16(i uint16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulUint32(i uint32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}

func (slf *Int) MulUint64(i uint64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mul(x, NewInt(i).ToBigint()))
}
func (slf *Int) Sub(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, i.ToBigint()))
}

func (slf *Int) SubInt(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubInt8(i int8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubInt16(i int16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubInt32(i int32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubInt64(i int64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubUint(i uint) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubUint8(i uint8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubUint16(i uint16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubUint32(i uint32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}

func (slf *Int) SubUint64(i uint64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sub(x, NewInt(i).ToBigint()))
}
