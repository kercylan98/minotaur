package charproc

import (
	"github.com/kercylan98/minotaur/toolkit/convert"
	"strings"
	"unsafe"
)

// NewBuilder 返回一个新的 Builder
func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	b   strings.Builder
	err error
}

// Write 向 Builder 中写入字节切片 p
func (b *Builder) Write(p ...byte) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.Write(p)
	return b
}

// WriteString 向 Builder 中写入字符串 s
func (b *Builder) WriteString(s string) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(s)
	return b
}

// WriteRune 向 Builder 中写入 Unicode 字符 r
func (b *Builder) WriteRune(r rune) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteRune(r)
	return b
}

// WriteRunes 向 Builder 中写入 Unicode 字符切片 r
func (b *Builder) WriteRunes(r ...rune) *Builder {
	if b.err != nil {
		return b
	}
	for _, v := range r {
		b.WriteRune(v)
	}
	return b
}

// WriteInt 向 Builder 中写入 int 类型的数字
func (b *Builder) WriteInt(i int) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.IntToString(i))
	return b
}

// WriteInt8 向 Builder 中写入 int8 类型的数字
func (b *Builder) WriteInt8(i int8) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Int8ToString(i))
	return b
}

// WriteInt16 向 Builder 中写入 int16 类型的数字
func (b *Builder) WriteInt16(i int16) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Int16ToString(i))
	return b
}

// WriteInt32 向 Builder 中写入 int32 类型的数字
func (b *Builder) WriteInt32(i int32) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Int32ToString(i))
	return b
}

// WriteInt64 向 Builder 中写入 int64 类型的数字
func (b *Builder) WriteInt64(i int64) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Int64ToString(i))
	return b
}

// WriteUint 向 Builder 中写入 uint 类型的数字
func (b *Builder) WriteUint(i uint) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.UintToString(i))
	return b
}

// WriteUint8 向 Builder 中写入 uint8 类型的数字
func (b *Builder) WriteUint8(i uint8) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Uint8ToString(i))
	return b
}

// WriteUint16 向 Builder 中写入 uint16 类型的数字
func (b *Builder) WriteUint16(i uint16) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Uint16ToString(i))
	return b
}

// WriteUint32 向 Builder 中写入 uint32 类型的数字
func (b *Builder) WriteUint32(i uint32) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Uint32ToString(i))
	return b
}

// WriteUint64 向 Builder 中写入 uint64 类型的数字
func (b *Builder) WriteUint64(i uint64) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Uint64ToString(i))
	return b
}

// WriteFloat32 向 Builder 中写入 float32 类型的数字
func (b *Builder) WriteFloat32(f float32) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Float32ToString(f))
	return b
}

// WriteFloat64 向 Builder 中写入 float64 类型的数字
func (b *Builder) WriteFloat64(f float64) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.Float64ToString(f))
	return b
}

// WriteBool 向 Builder 中写入 bool 类型的值
func (b *Builder) WriteBool(v bool) *Builder {
	if b.err != nil {
		return b
	}
	_, b.err = b.b.WriteString(convert.BooleanToString(v))
	return b
}

// Len 返回 Builder 中已写入的字节数
func (b *Builder) Len() int {
	return b.b.Len()
}

// Cap 返回 Builder 中的容量
func (b *Builder) Cap() int {
	return b.b.Cap()
}

// Reset 重置 Builder
func (b *Builder) Reset() *Builder {
	b.b.Reset()
	b.err = nil
	return b
}

// String 返回 Builder 中的字符串
func (b *Builder) String() string {
	return b.b.String()
}

// Err 返回 Builder 中的错误
func (b *Builder) Err() error {
	return b.err
}

// Bytes 返回 Builder 中的字节切片
func (b *Builder) Bytes() []byte {
	str := b.String()
	return *(*[]byte)(unsafe.Pointer(&str))
}
