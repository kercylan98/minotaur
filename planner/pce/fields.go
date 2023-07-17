package pce

import (
	"github.com/kercylan98/minotaur/utils/super"
	"math"
	"strconv"
)

type Int int

func (slf Int) TypeName() string {
	return "int"
}

func (slf Int) Zero() any {
	return int(0)
}

func (slf Int) Parse(value string) any {
	return super.StringToInt(value)
}

type Int8 int8

func (slf Int8) TypeName() string {
	return "int8"
}

func (slf Int8) Zero() any {
	return int8(0)
}

func (slf Int8) Parse(value string) any {
	v := super.StringToInt(value)
	if v < 0 {
		return int8(0)
	} else if v > math.MaxInt8 {
		return int8(math.MaxInt8)
	}
	return int8(v)
}

type Int16 int16

func (slf Int16) TypeName() string {
	return "int16"
}

func (slf Int16) Zero() any {
	return int16(0)
}

func (slf Int16) Parse(value string) any {
	v := super.StringToInt(value)
	if v < 0 {
		return int16(0)
	} else if v > math.MaxInt16 {
		return int16(math.MaxInt16)
	}
	return int16(v)
}

type Int32 int32

func (slf Int32) TypeName() string {
	return "int32"
}

func (slf Int32) Zero() any {
	return int32(0)
}

func (slf Int32) Parse(value string) any {
	v := super.StringToInt(value)
	if v < 0 {
		return int32(0)
	} else if v > math.MaxInt32 {
		return int32(math.MaxInt32)
	}
	return int32(v)
}

type Int64 int64

func (slf Int64) TypeName() string {
	return "int64"
}

func (slf Int64) Zero() any {
	return int64(0)
}

func (slf Int64) Parse(value string) any {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

type Uint uint

func (slf Uint) TypeName() string {
	return "uint"
}

func (slf Uint) Zero() any {
	return uint(0)
}

func (slf Uint) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return uint(0)
	}
	return uint(v)
}

type Uint8 uint8

func (slf Uint8) TypeName() string {
	return "uint8"
}

func (slf Uint8) Zero() any {
	return uint8(0)
}

func (slf Uint8) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return uint8(0)
	} else if v > math.MaxUint8 {
		return uint8(math.MaxUint8)
	}
	return uint8(v)
}

type Uint16 uint16

func (slf Uint16) TypeName() string {
	return "uint16"
}

func (slf Uint16) Zero() any {
	return uint16(0)
}

func (slf Uint16) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return uint16(0)
	} else if v > math.MaxUint16 {
		return uint16(math.MaxUint16)
	}
	return uint16(v)
}

type Uint32 uint32

func (slf Uint32) TypeName() string {
	return "uint32"
}

func (slf Uint32) Zero() any {
	return uint32(0)
}

func (slf Uint32) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return uint32(0)
	} else if v > math.MaxUint32 {
		return uint32(math.MaxUint32)
	}
	return uint32(v)
}

type Uint64 uint64

func (slf Uint64) TypeName() string {
	return "uint64"
}

func (slf Uint64) Zero() any {
	return uint64(0)
}

func (slf Uint64) Parse(value string) any {
	v, _ := strconv.ParseUint(value, 10, 64)
	if v < 0 {
		return uint64(0)
	}
	return v
}

type Float32 float32

func (slf Float32) TypeName() string {
	return "float32"
}

func (slf Float32) Zero() any {
	return float32(0)
}

func (slf Float32) Parse(value string) any {
	v, _ := strconv.ParseFloat(value, 32)
	return v
}

type Float64 float64

func (slf Float64) TypeName() string {
	return "float64"
}

func (slf Float64) Zero() any {
	return float64(0)
}

func (slf Float64) Parse(value string) any {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

type String string

func (slf String) TypeName() string {
	return "string"
}

func (slf String) Zero() any {
	return ""
}

func (slf String) Parse(value string) any {
	return value
}

type Bool bool

func (slf Bool) TypeName() string {
	return "bool"
}

func (slf Bool) Zero() any {
	return false
}

func (slf Bool) Parse(value string) any {
	v, _ := strconv.ParseBool(value)
	return v
}

type Byte byte

func (slf Byte) TypeName() string {
	return "byte"
}

func (slf Byte) Zero() any {
	return byte(0)
}

func (slf Byte) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return byte(0)
	} else if v > math.MaxUint8 {
		return byte(math.MaxUint8)
	}
	return byte(v)
}

type Rune rune

func (slf Rune) TypeName() string {
	return "rune"
}

func (slf Rune) Zero() any {
	return rune(0)
}

func (slf Rune) Parse(value string) any {
	v, _ := strconv.Atoi(value)
	if v < 0 {
		return rune(0)
	} else if v > math.MaxInt32 {
		return rune(math.MaxInt32)
	}
	return rune(v)
}

type Complex64 complex64

func (slf Complex64) TypeName() string {
	return "complex64"
}

func (slf Complex64) Zero() any {
	return complex64(0)
}

func (slf Complex64) Parse(value string) any {
	v, _ := strconv.ParseComplex(value, 64)
	return v
}

type Complex128 complex128

func (slf Complex128) TypeName() string {
	return "complex128"
}

func (slf Complex128) Zero() any {
	return complex128(0)
}

func (slf Complex128) Parse(value string) any {
	v, _ := strconv.ParseComplex(value, 128)
	return v
}

type Uintptr uintptr

func (slf Uintptr) TypeName() string {
	return "uintptr"
}

func (slf Uintptr) Zero() any {
	return uintptr(0)
}

func (slf Uintptr) Parse(value string) any {
	v, _ := strconv.ParseUint(value, 10, 64)
	return uintptr(v)
}

type Double float64

func (slf Double) TypeName() string {
	return "double"
}

func (slf Double) Zero() any {
	return float64(0)
}

func (slf Double) Parse(value string) any {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

type Float float32

func (slf Float) TypeName() string {
	return "float"
}

func (slf Float) Zero() any {
	return float32(0)
}

func (slf Float) Parse(value string) any {
	v, _ := strconv.ParseFloat(value, 32)
	return v
}

type Long int64

func (slf Long) TypeName() string {
	return "long"
}

func (slf Long) Zero() any {
	return int64(0)
}

func (slf Long) Parse(value string) any {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

type Short int16

func (slf Short) TypeName() string {
	return "short"
}

func (slf Short) Zero() any {
	return int16(0)
}

func (slf Short) Parse(value string) any {
	v, _ := strconv.ParseInt(value, 10, 16)
	return v
}

type Char int8

func (slf Char) TypeName() string {
	return "char"
}

func (slf Char) Zero() any {
	return int8(0)
}

func (slf Char) Parse(value string) any {
	v, _ := strconv.ParseInt(value, 10, 8)
	return v
}

type Number float64

func (slf Number) TypeName() string {
	return "number"
}

func (slf Number) Zero() any {
	return float64(0)
}

func (slf Number) Parse(value string) any {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

type Integer int64

func (slf Integer) TypeName() string {
	return "integer"
}

func (slf Integer) Zero() any {
	return int64(0)
}

func (slf Integer) Parse(value string) any {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

type Boolean bool

func (slf Boolean) TypeName() string {
	return "boolean"
}

func (slf Boolean) Zero() any {
	return false
}

func (slf Boolean) Parse(value string) any {
	v, _ := strconv.ParseBool(value)
	return v
}
