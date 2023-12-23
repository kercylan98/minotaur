package huge

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math/big"
)

var (
	FloatNegativeOne = NewFloat(-1.0)    // 默认初始化的-1值Float，应当将其当作常量使用
	FloatZero        = NewFloat(0.0)     // 默认初始化的0值Float，应当将其当作常量使用
	FloatOne         = NewFloat(1.0)     // 默认初始化的1值Float，应当将其当作常量使用
	FloatTen         = NewFloat(10.0)    // 默认初始化的10值Float，应当将其当作常量使用
	FloatHundred     = NewFloat(100.0)   // 默认初始化的100值Float，应当将其当作常量使用
	FloatThousand    = NewFloat(1000.0)  // 默认初始化的1000值Float，应当将其当作常量使用
	FloatTenThousand = NewFloat(10000.0) // 默认初始化的10000值Float，应当将其当作常量使用
)

type Float big.Float

// NewFloat 创建一个 Float
func NewFloat[T generic.Number](x T) *Float {
	return (*Float)(big.NewFloat(float64(x)))
}

// NewFloatByString 通过字符串创建一个 Float
//   - 如果字符串不是一个合法的数字，则返回 0
func NewFloatByString(i string) *Float {
	v, suc := new(big.Float).SetString(i)
	if !suc {
		return FloatZero.Copy()
	}
	return (*Float)(v)
}

func (slf *Float) Copy() *Float {
	return (*Float)(new(big.Float).Copy(slf.ToBigFloat()))
}

func (slf *Float) Set(i *Float) *Float {
	return (*Float)(slf.ToBigFloat().Set(i.ToBigFloat()))
}

func (slf *Float) IsZero() bool {
	if slf == nil || slf.EqualTo(FloatZero) {
		return true
	}
	return false
}

func (slf *Float) ToBigFloat() *big.Float {
	return (*big.Float)(slf)
}

// Cmp 比较，当 slf > i 时返回 1，当 slf < i 时返回 -1，当 slf == i 时返回 0
func (slf *Float) Cmp(i *Float) int {
	return slf.ToBigFloat().Cmp(i.ToBigFloat())
}

// GreaterThan 大于
func (slf *Float) GreaterThan(i *Float) bool {
	return slf.Cmp(i) > 0
}

// GreaterThanOrEqualTo 大于或等于
func (slf *Float) GreaterThanOrEqualTo(i *Float) bool {
	return slf.Cmp(i) >= 0
}

// LessThan 小于
func (slf *Float) LessThan(i *Float) bool {
	return slf.Cmp(i) < 0
}

// LessThanOrEqualTo 小于或等于
func (slf *Float) LessThanOrEqualTo(i *Float) bool {
	return slf.Cmp(i) <= 0
}

// EqualTo 等于
func (slf *Float) EqualTo(i *Float) bool {
	return slf.Cmp(i) == 0
}

func (slf *Float) Float64() float64 {
	f, _ := slf.ToBigFloat().Float64()
	return f
}

func (slf *Float) String() string {
	return slf.ToBigFloat().String()
}

func (slf *Float) Add(i *Float) *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Add(x, i.ToBigFloat()))
}

func (slf *Float) Sub(i *Float) *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Sub(x, i.ToBigFloat()))
}

func (slf *Float) Mul(i *Float) *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Mul(x, i.ToBigFloat()))
}

func (slf *Float) Div(i *Float) *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Quo(x, i.ToBigFloat()))
}

// Sqrt 平方根
func (slf *Float) Sqrt() *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Sqrt(x))
}

// Abs 返回数字的绝对值
func (slf *Float) Abs() *Float {
	x := slf.ToBigFloat()
	return (*Float)(x.Abs(x))
}

// Sign 返回数字的符号
//   - 1：正数
//   - 0：零
//   - -1：负数
func (slf *Float) Sign() int {
	return slf.ToBigFloat().Sign()
}

// IsPositive 是否为正数
func (slf *Float) IsPositive() bool {
	return slf.Sign() > 0
}

// IsNegative 是否为负数
func (slf *Float) IsNegative() bool {
	return slf.Sign() < 0
}
