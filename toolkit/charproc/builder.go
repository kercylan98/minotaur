package charproc

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"strings"
)

// NewBuilder 返回一个新的 Builder
func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	b   strings.Builder
	err error
	c   *color.Color
	end strings.Builder
}

// DisableColor 禁用 Builder 的颜色
func (b *Builder) DisableColor() *Builder {
	b.c = nil
	return b
}

// SetColor 设置 Builder 当前写入颜色
func (b *Builder) SetColor(c *color.Color) *Builder {
	if b.err != nil {
		return b
	}
	b.c = c
	return b
}

// Write 向 Builder 中写入字节切片 p
func (b *Builder) Write(p ...byte) *Builder {
	if b.err != nil || len(p) == 0 {
		return b
	}
	if b.c != nil {
		_, b.err = b.b.Write([]byte(b.c.Sprint(string(p))))
	} else {
		_, b.err = b.b.Write(p)
	}
	return b
}

// WriteToEnd 向 Builder 的尾部写入字节切片
func (b *Builder) WriteToEnd(p ...byte) *Builder {
	if b.err != nil {
		return b
	}
	if b.c != nil {
		_, b.err = b.end.Write([]byte(b.c.Sprint(string(p))))
	} else {
		_, b.err = b.end.Write(p)
	}
	return b
}

// WriteSprintf 向 Builder 中写入格式化字符串
func (b *Builder) WriteSprintf(format string, a ...any) *Builder {
	if b.err != nil {
		return b
	}
	if b.c != nil {
		_, b.err = b.b.WriteString(b.c.Sprint(fmt.Sprintf(format, a...)))
	} else {
		_, b.err = b.b.WriteString(fmt.Sprintf(format, a...))
	}
	return b
}

// WriteSprintfToEnd 向 Builder 的尾部写入格式化字符串
func (b *Builder) WriteSprintfToEnd(format string, a ...any) *Builder {
	if b.err != nil {
		return b
	}
	if b.c != nil {
		_, b.err = b.end.WriteString(b.c.Sprint(fmt.Sprintf(format, a...)))
	} else {
		_, b.err = b.end.WriteString(fmt.Sprintf(format, a...))
	}
	return b
}

// WriteString 向 Builder 中写入字符串 s
func (b *Builder) WriteString(s string) *Builder {
	if b.err != nil || s == "" {
		return b
	}
	if b.c != nil {
		_, b.err = b.b.WriteString(b.c.Sprint(s))
	} else {
		_, b.err = b.b.WriteString(s)
	}
	return b
}

// WriteStringToEnd 向 Builder 的尾部写入字符串
func (b *Builder) WriteStringToEnd(str string) *Builder {
	if b.err != nil {
		return b
	}
	b.end.WriteString(str)
	return b
}

// WriteRune 向 Builder 中写入 Unicode 字符 r
func (b *Builder) WriteRune(r rune) *Builder {
	if b.err != nil || r == 0 {
		return b
	}
	if b.c != nil {
		_, b.err = b.b.WriteString(b.c.Sprint(string(r)))
	} else {
		_, b.err = b.b.WriteRune(r)
	}
	return b
}

// WriteRunes 向 Builder 中写入 Unicode 字符切片 r
func (b *Builder) WriteRunes(r ...rune) *Builder {
	if b.err != nil || len(r) == 0 {
		return b
	}
	for _, v := range r {
		if b.c != nil {
			_, b.err = b.b.WriteString(b.c.Sprint(string(v)))
		} else {
			_, b.err = b.b.WriteRune(v)
		}
		if b.err != nil {
			break
		}
	}
	return b
}

// WriteInt 向 Builder 中写入 int 类型的数字
func (b *Builder) WriteInt(i int) *Builder {
	return b.WriteString(convert.IntToString(i))
}

// WriteIntToEnd 向 Builder 的尾部写入 int 类型的数字
func (b *Builder) WriteIntToEnd(i int) *Builder {
	return b.WriteStringToEnd(convert.IntToString(i))
}

// WriteInt8 向 Builder 中写入 int8 类型的数字
func (b *Builder) WriteInt8(i int8) *Builder {
	return b.WriteString(convert.Int8ToString(i))
}

// WriteInt16 向 Builder 中写入 int16 类型的数字
func (b *Builder) WriteInt16(i int16) *Builder {
	return b.WriteString(convert.Int16ToString(i))
}

// WriteInt32 向 Builder 中写入 int32 类型的数字
func (b *Builder) WriteInt32(i int32) *Builder {
	return b.WriteString(convert.Int32ToString(i))
}

// WriteInt64 向 Builder 中写入 int64 类型的数字
func (b *Builder) WriteInt64(i int64) *Builder {
	return b.WriteString(convert.Int64ToString(i))
}

// WriteUint 向 Builder 中写入 uint 类型的数字
func (b *Builder) WriteUint(i uint) *Builder {
	return b.WriteString(convert.UintToString(i))
}

// WriteUint8 向 Builder 中写入 uint8 类型的数字
func (b *Builder) WriteUint8(i uint8) *Builder {
	return b.WriteString(convert.Uint8ToString(i))
}

// WriteUint16 向 Builder 中写入 uint16 类型的数字
func (b *Builder) WriteUint16(i uint16) *Builder {
	return b.WriteString(convert.Uint16ToString(i))
}

// WriteUint32 向 Builder 中写入 uint32 类型的数字
func (b *Builder) WriteUint32(i uint32) *Builder {
	return b.WriteString(convert.Uint32ToString(i))
}

// WriteUint64 向 Builder 中写入 uint64 类型的数字
func (b *Builder) WriteUint64(i uint64) *Builder {
	return b.WriteString(convert.Uint64ToString(i))
}

// WriteFloat32 向 Builder 中写入 float32 类型的数字
func (b *Builder) WriteFloat32(f float32) *Builder {
	return b.WriteString(convert.Float32ToString(f))
}

// WriteFloat64 向 Builder 中写入 float64 类型的数字
func (b *Builder) WriteFloat64(f float64) *Builder {
	return b.WriteString(convert.Float64ToString(f))
}

// WriteBool 向 Builder 中写入 bool 类型的值
func (b *Builder) WriteBool(v bool) *Builder {
	return b.WriteString(convert.BooleanToString(v))
}

// Len 返回 Builder 中已写入的字节数
func (b *Builder) Len() int {
	return b.b.Len() + b.end.Len()
}

// Cap 返回 Builder 中的容量
func (b *Builder) Cap() int {
	return b.b.Cap() + b.end.Cap()
}

// Reset 重置 Builder
func (b *Builder) Reset() *Builder {
	b.b.Reset()
	b.end.Reset()
	b.err = nil
	b.c = nil
	return b
}

// String 返回 Builder 中的字符串
func (b *Builder) String() (string, error) {
	return b.b.String() + b.end.String(), b.err
}

// Err 返回 Builder 中的错误
func (b *Builder) Err() error {
	return b.err
}

// Bytes 返回 Builder 中的字节切片
func (b *Builder) Bytes() ([]byte, error) {
	str, err := b.String()
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
