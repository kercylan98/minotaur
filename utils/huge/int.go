package huge

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math/big"
)

var (
	IntNegativeOne = NewInt(-1)    // 默认初始化的-1值Int，应当将其当作常量使用
	IntZero        = NewInt(0)     // 默认初始化的0值Int，应当将其当作常量使用
	IntOne         = NewInt(1)     // 默认初始化的1值Int，应当将其当作常量使用
	IntTen         = NewInt(10)    // 默认初始化的10值Int，应当将其当作常量使用
	IntHundred     = NewInt(100)   // 默认初始化的100值Int，应当将其当作常量使用
	IntThousand    = NewInt(1000)  // 默认初始化的1000值Int，应当将其当作常量使用
	IntTenThousand = NewInt(10000) // 默认初始化的10000值Int，应当将其当作常量使用
)

type Int big.Int

// NewInt 创建一个 Int 对象，该对象的值为 x
func NewInt[T generic.Basic](x T) *Int {
	var xa any = x
	switch x := xa.(type) {
	case int:
		return (*Int)(big.NewInt(int64(x)))
	case int8:
		return (*Int)(big.NewInt(int64(x)))
	case int16:
		return (*Int)(big.NewInt(int64(x)))
	case int32:
		return (*Int)(big.NewInt(int64(x)))
	case int64:
		return (*Int)(big.NewInt(x))
	case uint:
		return (*Int)(big.NewInt(int64(x)))
	case uint8:
		return (*Int)(big.NewInt(int64(x)))
	case uint16:
		return (*Int)(big.NewInt(int64(x)))
	case uint32:
		return (*Int)(big.NewInt(int64(x)))
	case uint64:
		return (*Int)(big.NewInt(int64(x)))
	case string:
		si, suc := new(big.Int).SetString(x, 10)
		if !suc {
			return (*Int)(big.NewInt(0))
		}
		return (*Int)(si)
	case bool:
		if x {
			return (*Int)(big.NewInt(1))
		}
		return (*Int)(big.NewInt(0))
	case float32:
		return (*Int)(big.NewInt(int64(x)))
	case float64:
		return (*Int)(big.NewInt(int64(x)))
	}
	return (*Int)(big.NewInt(0))
}

// applyIntOperation 应用一个 Int 操作
func applyIntOperation[T generic.Number](v *Int, i T, op func(*big.Int, *big.Int) *big.Int) *Int {
	return (*Int)(op(v.ToBigint(), NewInt(i).ToBigint()))
}

// Copy 拷贝当前 Int 对象
func (slf *Int) Copy() *Int {
	return (*Int)(new(big.Int).Set(slf.ToBigint()))
}

// Set 设置当前 Int 对象的值为 i
func (slf *Int) Set(i *Int) *Int {
	return (*Int)(slf.ToBigint().Set(i.ToBigint()))
}

// SetString 设置当前 Int 对象的值为 i
func (slf *Int) SetString(i string) *Int {
	return (*Int)(slf.ToBigint().Set((*big.Int)(NewInt(i))))
}

// SetInt 设置当前 Int 对象的值为 i
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
	if slf == nil {
		return big.NewInt(0)
	}
	return (*big.Int)(slf)
}

// Cmp 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0
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
	if slf == nil {
		return "0"
	}
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

func (slf *Int) Div(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Div(x, i.ToBigint()))
}

func (slf *Int) DivInt(i int) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivInt8(i int8) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivInt16(i int16) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivInt32(i int32) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivInt64(i int64) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivUint(i uint) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivUint8(i uint8) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivUint16(i uint16) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivUint32(i uint32) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) DivUint64(i uint64) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Div)
}

func (slf *Int) Mod(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Mod(x, i.ToBigint()))
}

func (slf *Int) ModInt(i int) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModInt8(i int8) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModInt16(i int16) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModInt32(i int32) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModInt64(i int64) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModUint(i uint) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModUint8(i uint8) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModUint16(i uint16) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModUint32(i uint32) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) ModUint64(i uint64) *Int {
	return applyIntOperation(slf, i, NewInt(i).ToBigint().Mod)
}

func (slf *Int) Pow(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, i.ToBigint(), nil))
}

func (slf *Int) PowInt(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowInt8(i int8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowInt16(i int16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowInt32(i int32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowInt64(i int64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowUint(i uint) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowUint8(i uint8) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowUint16(i uint16) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowUint32(i uint32) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

func (slf *Int) PowUint64(i uint64) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Exp(x, NewInt(i).ToBigint(), nil))
}

// Lsh 左移
func (slf *Int) Lsh(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Lsh(x, uint(i)))
}

// Rsh 右移
func (slf *Int) Rsh(i int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Rsh(x, uint(i)))
}

// And 与
func (slf *Int) And(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.And(x, i.ToBigint()))
}

// AndNot 与非
func (slf *Int) AndNot(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.AndNot(x, i.ToBigint()))
}

// Or 或
func (slf *Int) Or(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Or(x, i.ToBigint()))
}

// Xor 异或
func (slf *Int) Xor(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.Xor(x, i.ToBigint()))
}

// Not 非
func (slf *Int) Not() *Int {
	x := slf.ToBigint()
	return (*Int)(x.Not(x))
}

// Sqrt 平方根
func (slf *Int) Sqrt() *Int {
	x := slf.ToBigint()
	return (*Int)(x.Sqrt(x))
}

// GCD 最大公约数
func (slf *Int) GCD(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.GCD(nil, nil, x, i.ToBigint()))
}

// LCM 最小公倍数
func (slf *Int) LCM(i *Int) *Int {
	sb := slf.ToBigint()
	ib := i.ToBigint()
	gcd := new(big.Int).GCD(nil, nil, sb, ib)
	absProduct := new(big.Int).Mul(sb, ib).Abs(new(big.Int).Mul(sb, ib))
	lcm := new(big.Int).Div(absProduct, gcd)
	return (*Int)(lcm)
}

// ModInverse 模反元素
func (slf *Int) ModInverse(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.ModInverse(x, i.ToBigint()))
}

// ModSqrt 模平方根
func (slf *Int) ModSqrt(i *Int) *Int {
	x := slf.ToBigint()
	return (*Int)(x.ModSqrt(x, i.ToBigint()))
}

// BitLen 二进制长度
func (slf *Int) BitLen() int {
	return slf.ToBigint().BitLen()
}

// Bit 二进制位
func (slf *Int) Bit(i int) uint {
	return slf.ToBigint().Bit(i)
}

// SetBit 设置二进制位
func (slf *Int) SetBit(i int, v uint) *Int {
	x := slf.ToBigint()
	return (*Int)(x.SetBit(x, i, v))
}

// Neg 返回数字的相反数
func (slf *Int) Neg() *Int {
	x := slf.ToBigint()
	return (*Int)(x.Neg(x))
}

// Abs 返回数字的绝对值
func (slf *Int) Abs() *Int {
	x := slf.ToBigint()
	return (*Int)(x.Abs(x))
}

// Sign 返回数字的符号
//   - 1：正数
//   - 0：零
//   - -1：负数
func (slf *Int) Sign() int {
	return slf.ToBigint().Sign()
}

// IsPositive 是否为正数
func (slf *Int) IsPositive() bool {
	return slf.Sign() > 0
}

// IsNegative 是否为负数
func (slf *Int) IsNegative() bool {
	return slf.Sign() < 0
}

// IsEven 是否为偶数
func (slf *Int) IsEven() bool {
	return slf.ToBigint().Bit(0) == 0
}

// IsOdd 是否为奇数
func (slf *Int) IsOdd() bool {
	return slf.ToBigint().Bit(0) == 1
}

// ProportionalCalc 比例计算，该函数会再 formula 返回值的基础上除以 proportional
//   - formula 为计算公式，该公式的参数为调用该函数的 Int 的拷贝
func (slf *Int) ProportionalCalc(proportional *Int, formula func(v *Int) *Int) *Int {
	return formula(slf.Copy()).Div(proportional)
}
