package effect

import (
	"github.com/shopspring/decimal"
	"math/big"
)

type Attribute struct {
	v decimal.Decimal
}

func (attr Attribute) String() string {
	return attr.v.String()
}

func (attr Attribute) AddInt(v int) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromInt(int64(v))),
	}
}

func (attr Attribute) AddInt8(v int8) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) AddInt16(v int16) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) AddInt32(v int32) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromInt32(v)),
	}
}

func (attr Attribute) AddInt64(v int64) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromInt(v)),
	}
}

func (attr Attribute) AddUint(v uint) Attribute {
	return attr.AddUint64(uint64(v))
}

func (attr Attribute) AddUint8(v uint8) Attribute {
	return attr.AddUint64(uint64(v))
}

func (attr Attribute) AddUint16(v uint16) Attribute {
	return attr.AddUint64(uint64(v))
}

func (attr Attribute) AddUint32(v uint32) Attribute {
	return attr.AddUint64(uint64(v))
}

func (attr Attribute) AddUint64(v uint64) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)),
	}
}

func (attr Attribute) SubInt(v int) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromInt(int64(v))),
	}
}

func (attr Attribute) SubInt8(v int8) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) SubInt16(v int16) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) SubInt32(v int32) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromInt32(v)),
	}
}

func (attr Attribute) SubInt64(v int64) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromInt(v)),
	}
}

func (attr Attribute) SubUint(v uint) Attribute {
	return attr.SubUint64(uint64(v))
}

func (attr Attribute) SubUint8(v uint8) Attribute {
	return attr.SubUint64(uint64(v))
}

func (attr Attribute) SubUint16(v uint16) Attribute {
	return attr.SubUint64(uint64(v))
}

func (attr Attribute) SubUint32(v uint32) Attribute {
	return attr.SubUint64(uint64(v))
}

func (attr Attribute) SubUint64(v uint64) Attribute {
	return attr.SubUint64(uint64(v))
}

func (attr Attribute) MulInt(v int) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromInt(int64(v))),
	}
}

func (attr Attribute) MulInt8(v int8) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) MulInt16(v int16) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) MulInt32(v int32) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromInt32(v)),
	}
}

func (attr Attribute) MulInt64(v int64) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromInt(v)),
	}
}

func (attr Attribute) MulUint(v uint) Attribute {
	return attr.MulUint64(uint64(v))
}

func (attr Attribute) MulUint8(v uint8) Attribute {
	return attr.MulUint64(uint64(v))
}

func (attr Attribute) MulUint16(v uint16) Attribute {
	return attr.MulUint64(uint64(v))
}

func (attr Attribute) MulUint32(v uint32) Attribute {
	return attr.MulUint64(uint64(v))
}

func (attr Attribute) MulUint64(v uint64) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)),
	}
}

func (attr Attribute) DivInt(v int) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromInt(int64(v))),
	}
}

func (attr Attribute) DivInt8(v int8) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) DivInt16(v int16) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromInt32(int32(v))),
	}
}

func (attr Attribute) DivInt32(v int32) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromInt32(v)),
	}
}

func (attr Attribute) DivInt64(v int64) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromInt(v)),
	}
}

func (attr Attribute) DivUint(v uint) Attribute {
	return attr.DivUint64(uint64(v))
}

func (attr Attribute) DivUint8(v uint8) Attribute {
	return attr.DivUint64(uint64(v))
}

func (attr Attribute) DivUint16(v uint16) Attribute {
	return attr.DivUint64(uint64(v))
}

func (attr Attribute) DivUint32(v uint32) Attribute {
	return attr.DivUint64(uint64(v))
}

func (attr Attribute) DivUint64(v uint64) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)),
	}
}

func (attr Attribute) Add(v Attribute) Attribute {
	return Attribute{
		attr.v.Add(v.v),
	}
}

func (attr Attribute) Sub(v Attribute) Attribute {
	return Attribute{
		attr.v.Sub(v.v),
	}
}

func (attr Attribute) Mul(v Attribute) Attribute {
	return Attribute{
		attr.v.Mul(v.v),
	}
}

func (attr Attribute) Div(v Attribute) Attribute {
	return Attribute{
		attr.v.Div(v.v),
	}
}

func (attr Attribute) AddFloat32(v float32) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromFloat32(v)),
	}
}

func (attr Attribute) AddFloat64(v float64) Attribute {
	return Attribute{
		attr.v.Add(decimal.NewFromFloat(v)),
	}
}

func (attr Attribute) SubFloat32(v float32) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromFloat32(v)),
	}
}

func (attr Attribute) SubFloat64(v float64) Attribute {
	return Attribute{
		attr.v.Sub(decimal.NewFromFloat(v)),
	}
}

func (attr Attribute) MulFloat32(v float32) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromFloat32(v)),
	}
}

func (attr Attribute) MulFloat64(v float64) Attribute {
	return Attribute{
		attr.v.Mul(decimal.NewFromFloat(v)),
	}
}

func (attr Attribute) DivFloat32(v float32) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromFloat32(v)),
	}
}

func (attr Attribute) DivFloat64(v float64) Attribute {
	return Attribute{
		attr.v.Div(decimal.NewFromFloat(v)),
	}
}

func (attr Attribute) Max(v Attribute) Attribute {
	if attr.v.GreaterThan(v.v) {
		return attr
	} else {
		return v
	}
}

func (attr Attribute) Min(v Attribute) Attribute {
	if attr.v.LessThan(v.v) {
		return attr
	} else {
		return v
	}
}

func (attr Attribute) MaxFloat32(v float32) Attribute {
	return attr.Max(Attribute{decimal.NewFromFloat32(v)})
}

func (attr Attribute) MaxFloat64(v float64) Attribute {
	return attr.Max(Attribute{decimal.NewFromFloat(v)})
}

func (attr Attribute) MinFloat32(v float32) Attribute {
	return attr.Min(Attribute{decimal.NewFromFloat32(v)})
}

func (attr Attribute) MinFloat64(v float64) Attribute {
	return attr.Min(Attribute{decimal.NewFromFloat(v)})
}

func (attr Attribute) MaxInt(v int) Attribute {
	return attr.Max(Attribute{decimal.NewFromInt(int64(v))})
}

func (attr Attribute) MinInt(v int) Attribute {
	return attr.Min(Attribute{decimal.NewFromInt(int64(v))})
}

func (attr Attribute) MaxInt8(v int8) Attribute {
	return attr.Max(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) MinInt8(v int8) Attribute {
	return attr.Min(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) MaxInt16(v int16) Attribute {
	return attr.Max(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) MinInt16(v int16) Attribute {
	return attr.Min(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) MaxInt32(v int32) Attribute {
	return attr.Max(Attribute{decimal.NewFromInt32(v)})
}

func (attr Attribute) MinInt32(v int32) Attribute {
	return attr.Min(Attribute{decimal.NewFromInt32(v)})
}

func (attr Attribute) MaxInt64(v int64) Attribute {
	return attr.Max(Attribute{decimal.NewFromInt(v)})
}

func (attr Attribute) MinInt64(v int64) Attribute {
	return attr.Min(Attribute{decimal.NewFromInt(v)})
}

func (attr Attribute) MaxUint(v uint) Attribute {
	return attr.Max(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MaxUint8(v uint8) Attribute {
	return attr.Max(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MaxUint16(v uint16) Attribute {
	return attr.Max(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MaxUint32(v uint32) Attribute {
	return attr.Max(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MaxUint64(v uint64) Attribute {
	return attr.Max(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)})
}

func (attr Attribute) MinUint(v uint) Attribute {
	return attr.Min(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MinUint8(v uint8) Attribute {
	return attr.Min(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MinUint16(v uint16) Attribute {
	return attr.Min(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MinUint32(v uint32) Attribute {
	return attr.Min(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) MinUint64(v uint64) Attribute {
	return attr.Min(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)})
}

func (attr Attribute) Equal(v Attribute) bool {
	return attr.v.Equal(v.v)
}

func (attr Attribute) EqualFloat32(v float32) bool {
	return attr.Equal(Attribute{decimal.NewFromFloat32(v)})
}

func (attr Attribute) EqualFloat64(v float64) bool {
	return attr.Equal(Attribute{decimal.NewFromFloat(v)})
}

func (attr Attribute) EqualInt(v int) bool {
	return attr.Equal(Attribute{decimal.NewFromInt(int64(v))})
}

func (attr Attribute) EqualInt8(v int8) bool {
	return attr.Equal(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) EqualInt16(v int16) bool {
	return attr.Equal(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) EqualInt32(v int32) bool {
	return attr.Equal(Attribute{decimal.NewFromInt32(v)})
}

func (attr Attribute) EqualInt64(v int64) bool {
	return attr.Equal(Attribute{decimal.NewFromInt(v)})
}

func (attr Attribute) EqualUint(v uint) bool {
	return attr.Equal(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) EqualUint8(v uint8) bool {
	return attr.Equal(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) EqualUint16(v uint16) bool {
	return attr.Equal(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) EqualUint32(v uint32) bool {
	return attr.Equal(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) EqualUint64(v uint64) bool {
	return attr.Equal(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)})
}

func (attr Attribute) GreaterThan(v Attribute) bool {
	return attr.v.GreaterThan(v.v)
}

func (attr Attribute) GreaterThanFloat32(v float32) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromFloat32(v)})
}

func (attr Attribute) GreaterThanFloat64(v float64) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromFloat(v)})
}

func (attr Attribute) GreaterThanInt(v int) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromInt(int64(v))})
}

func (attr Attribute) GreaterThanInt8(v int8) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) GreaterThanInt16(v int16) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) GreaterThanInt32(v int32) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromInt32(v)})
}

func (attr Attribute) GreaterThanInt64(v int64) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromInt(v)})
}

func (attr Attribute) GreaterThanUint(v uint) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) GreaterThanUint8(v uint8) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) GreaterThanUint16(v uint16) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) GreaterThanUint32(v uint32) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}
func (attr Attribute) GreaterThanUint64(v uint64) bool {
	return attr.GreaterThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)})
}

func (attr Attribute) LessThan(v Attribute) bool {
	return attr.v.LessThan(v.v)
}

func (attr Attribute) LessThanFloat32(v float32) bool {
	return attr.LessThan(Attribute{decimal.NewFromFloat32(v)})
}

func (attr Attribute) LessThanFloat64(v float64) bool {
	return attr.LessThan(Attribute{decimal.NewFromFloat(v)})
}

func (attr Attribute) LessThanInt(v int) bool {
	return attr.LessThan(Attribute{decimal.NewFromInt(int64(v))})
}

func (attr Attribute) LessThanInt8(v int8) bool {
	return attr.LessThan(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) LessThanInt16(v int16) bool {
	return attr.LessThan(Attribute{decimal.NewFromInt32(int32(v))})
}

func (attr Attribute) LessThanInt32(v int32) bool {
	return attr.LessThan(Attribute{decimal.NewFromInt32(v)})
}

func (attr Attribute) LessThanInt64(v int64) bool {
	return attr.LessThan(Attribute{decimal.NewFromInt(v)})
}

func (attr Attribute) LessThanUint(v uint) bool {
	return attr.LessThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) LessThanUint8(v uint8) bool {
	return attr.LessThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) LessThanUint16(v uint16) bool {
	return attr.LessThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) LessThanUint32(v uint32) bool {
	return attr.LessThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(v)), 0)})
}

func (attr Attribute) LessThanUint64(v uint64) bool {
	return attr.LessThan(Attribute{decimal.NewFromBigInt(new(big.Int).SetUint64(v), 0)})
}
