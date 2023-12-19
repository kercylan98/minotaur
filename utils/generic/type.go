package generic

// Ordered 可排序类型
type Ordered interface {
	Integer | Float | ~string
}

// Number 数字类型
type Number interface {
	Integer | Float
}

// SignedNumber 有符号数字类型
type SignedNumber interface {
	Signed | Float
}

// Integer 整数类型
type Integer interface {
	Signed | Unsigned
}

// Signed 有符号整数类型
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 无符号整数类型
type Unsigned interface {
	UnsignedNumber | ~uintptr
}

// UnsignedNumber 无符号数字类型
type UnsignedNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float 浮点类型
type Float interface {
	~float32 | ~float64
}

// Basic 基本类型
type Basic interface {
	Signed | Unsigned | Float | ~string | ~bool | ~byte
}
