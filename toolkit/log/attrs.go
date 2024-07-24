package log

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"log/slog"
	"runtime/debug"
	"time"
)

type Attr = slog.Attr

// Skip 构造一个无操作字段，这在处理其他 Attr 构造函数中的无效输入时通常很有用
//   - 该函数还支持将其他字段快捷的转换为 Skip 字段
func Skip(vs ...any) Attr {
	return Attr{Key: ""}
}

// Duration 使用给定的键和值构造一个字段。编码器控制持续时间的序列化方式
func Duration(key string, val time.Duration) Attr {
	return slog.String(key, val.String())
}

// DurationP 构造一个带有 time.Duration 的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func DurationP(key string, val *time.Duration) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Duration(key, *val)
}

// Bool 构造一个带有布尔值的字段
func Bool(key string, val bool) Attr {
	return slog.Bool(key, val)
}

// BoolP 构造一个带有布尔值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func BoolP(key string, val *bool) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Bool(key, *val)
}

// String 构造一个带有字符串值的字段
func String(key, val string) Attr {
	return slog.String(key, val)
}

// StringP 构造一个带有字符串值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func StringP(key string, val *string) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return String(key, *val)
}

// Int 构造一个带有整数值的字段
func Int[I constraints.Int](key string, val I) Attr {
	return slog.Int(key, int(val))
}

// IntP 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func IntP[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Int(key, *val)
}

// Int8 构造一个带有整数值的字段
func Int8[I constraints.Int](key string, val I) Attr {
	return slog.Int(key, int(val))
}

// Int8P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Int8P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Int8(key, *val)
}

// Int16 构造一个带有整数值的字段
func Int16[I constraints.Int](key string, val I) Attr {
	return slog.Int(key, int(val))
}

// Int16P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Int16P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Int16(key, *val)
}

// Int32 构造一个带有整数值的字段
func Int32[I constraints.Int](key string, val I) Attr {
	return slog.Int(key, int(val))
}

// Int32P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Int32P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Int32(key, *val)
}

// Int64 构造一个带有整数值的字段
func Int64[I constraints.Int](key string, val I) Attr {
	return slog.Int64(key, int64(val))
}

// Int64P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Int64P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Int64(key, *val)
}

// Uint 构造一个带有整数值的字段
func Uint[I constraints.Int](key string, val I) Attr {
	return slog.Uint64(key, uint64(val))
}

// UintP 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func UintP[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Uint(key, *val)
}

// Uint8 构造一个带有整数值的字段
func Uint8[I constraints.Int](key string, val I) Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint8P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Uint8P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Uint8(key, *val)
}

// Uint16 构造一个带有整数值的字段
func Uint16[I constraints.Int](key string, val I) Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint16P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Uint16P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Uint16(key, *val)
}

// Uint32 构造一个带有整数值的字段
func Uint32[I constraints.Int](key string, val I) Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint32P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Uint32P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Uint32(key, *val)
}

// Uint64 构造一个带有整数值的字段
func Uint64[I constraints.Int](key string, val I) Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint64P 构造一个带有整数值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Uint64P[I constraints.Int](key string, val *I) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Uint64(key, *val)
}

// Float 构造一个带有浮点值的字段
func Float[F constraints.Float](key string, val F) Attr {
	return slog.Float64(key, float64(val))
}

// FloatP 构造一个带有浮点值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func FloatP[F constraints.Float](key string, val *F) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Float(key, *val)
}

// Float32 构造一个带有浮点值的字段
func Float32[F constraints.Float](key string, val F) Attr {
	return slog.Float64(key, float64(val))
}

// Float32P 构造一个带有浮点值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Float32P[F constraints.Float](key string, val *F) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Float32(key, *val)
}

// Float64 构造一个带有浮点值的字段
func Float64[F constraints.Float](key string, val F) Attr {
	return slog.Float64(key, float64(val))
}

// Float64P 构造一个带有浮点值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func Float64P[F constraints.Float](key string, val *F) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Float64(key, *val)
}

// Time 构造一个带有时间值的字段
func Time(key string, val time.Time) Attr {
	return slog.Time(key, val)
}

// TimeP 构造一个带有时间值的字段。返回的 Attr 将在适当的时候安全且显式地表示 "null"
func TimeP(key string, val *time.Time) Attr {
	if val == nil {
		return slog.Any(key, nil)
	}
	return Time(key, *val)
}

// Any 构造一个带有任意值的字段
func Any(key string, val any) Attr {
	return slog.Any(key, val)
}

// Group 返回分组字段
func Group(key string, args ...any) Attr {
	return slog.Group(key, args...)
}

// Stack 返回堆栈字段
func Stack(key string) Attr {
	return slog.Any(key, stack(debug.Stack()))
}

// StackData 返回堆栈字段
func StackData(key string, data []byte) Attr {
	return slog.Any(key, stack(data))
}

// Err 构造一个带有错误值的字段
func Err(err error) Attr {
	return slog.Any("error", err)
}
