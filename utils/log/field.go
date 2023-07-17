package log

import (
	"go.uber.org/zap"
)

type Field = zap.Field

var (
	// Skip 构造一个无操作字段，这在处理其他 Field 构造函数中的无效输入时通常很有用
	Skip = zap.Skip

	// Binary 构造一个携带不透明二进制 blob 的字段。二进制数据以适合编码的格式进行序列化。例如，JSON 编码器对二进制 blob 进行 base64 编码。要记录 UTF-8 编码文本，请使用 ByteString
	Binary = zap.Binary

	// Bool 构造一个带有布尔值的字段
	Bool = zap.Bool

	// BoolP 构造一个带有布尔值的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	BoolP = zap.Boolp

	// ByteString 构造一个将 UTF-8 编码文本作为 [] 字节传送的字段。要记录不透明的二进制 blob（不一定是有效的 UTF-8），请使用 Binary
	ByteString = zap.ByteString

	// Complex128 构造一个带有复数的字段。与大多数数字字段不同，这需要分配（将complex128转换为interface{}）
	Complex128 = zap.Complex128

	// Complex128P 构造一个带有complex128 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Complex128P = zap.Complex128p

	// Complex64 构造一个带有复数的字段。与大多数数字字段不同，这需要分配（将complex64转换为interface{}）
	Complex64 = zap.Complex64

	// Complex64P 构造一个带有complex64 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Complex64P = zap.Complex64p

	// Float64 构造一个带有 float64 的字段。浮点值的表示方式取决于编码器，因此封送处理必然是惰性的
	Float64 = zap.Float64

	// Float64P 构造一个带有 float64 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Float64P = zap.Float64p

	// Float32 构造一个带有 float32 的字段。浮点值的表示方式取决于编码器，因此封送处理必然是惰性的
	Float32 = zap.Float32

	// Float32P 构造一个带有 float32 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Float32P = zap.Float32p

	// Int constructs a field with the given key and value.
	Int = zap.Int

	// IntP 构造一个带有 int 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	IntP = zap.Intp

	// Int64 使用给定的键和值构造一个字段.
	Int64 = zap.Int64

	// Int64P 构造一个带有 int64 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Int64P = zap.Int64p

	// Int32 使用给定的键和值构造一个字段
	Int32 = zap.Int32

	// Int32P 构造一个带有 int32 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”.
	Int32P = zap.Int32p

	// Int16 使用给定的键和值构造一个字段
	Int16 = zap.Int16

	// Int16P 构造一个带有 int16 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Int16P = zap.Int16p

	// Int8 使用给定的键和值构造一个字段
	Int8 = zap.Int8

	// Int8P 构造一个带有 int8 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Int8P = zap.Int8p

	// String 使用给定的键和值构造一个字段
	String = zap.String

	// StringP 构造一个带有字符串的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	StringP = zap.Stringp

	// Uint 使用给定的键和值构造一个字段
	Uint = zap.Uint

	// UintP 构造一个带有 uint 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	UintP = zap.Uintp

	// Uint64 使用给定的键和值构造一个字段
	Uint64 = zap.Uint64

	// Uint64P 构造一个带有 uint64 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Uint64P = zap.Uint64p

	// Uint32 使用给定的键和值构造一个字段
	Uint32 = zap.Uint32

	// Uint32P 构造一个带有 uint32 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Uint32P = zap.Uint32p

	// Uint16 使用给定的键和值构造一个字段
	Uint16 = zap.Uint16

	// Uint16P 构造一个带有 uint16 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Uint16P = zap.Uint16p

	// Uint8 使用给定的键和值构造一个字段
	Uint8 = zap.Uint8

	// Uint8P 构造一个带有 uint8 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	Uint8P = zap.Uint8p

	// Uintptr 使用给定的键和值构造一个字段
	Uintptr = zap.Uintptr

	// UintptrP 构造一个带有 uintptr 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	UintptrP = zap.Uintptrp

	// Reflect 使用给定的键和任意对象构造一个字段。它使用适当的编码、基于反射的方式将几乎任何对象延迟序列化到日志记录上下文中，但它相对较慢且分配繁重。在测试之外，Any 始终是更好的选择
	//  - 如果编码失败（例如，尝试将 map[int]string 序列化为 JSON），Reflect 将在最终日志输出中包含错误消息
	Reflect = zap.Reflect

	// Namespace 命名空间在记录器的上下文中创建一个命名的、隔离的范围。所有后续字段都将添加到新的命名空间中
	//  - 这有助于防止将记录器注入子组件或第三方库时发生按键冲突
	Namespace = zap.Namespace

	// Stringer 使用给定的键和值的 String 方法的输出构造一个字段。 Stringer 的 String 方法被延迟调用
	Stringer = zap.Stringer

	// Time 使用给定的键和值构造一个 Field。编码器控制时间的序列化方式
	Time = zap.Time

	// TimeP 构造一个带有 time.Time 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	TimeP = zap.Timep

	// Stack 构造一个字段，在提供的键下存储当前 goroutine 的堆栈跟踪。请记住，进行堆栈跟踪是急切且昂贵的（相对而言）；此操作会进行分配并花费大约两微秒的时间
	Stack = zap.Stack

	// StackSkip 构造一个与 Stack 类似的字段，但也会从堆栈跟踪顶部跳过给定数量的帧
	StackSkip = zap.StackSkip

	// Duration 使用给定的键和值构造一个字段。编码器控制持续时间的序列化方式
	Duration = zap.Duration

	// DurationP 构造一个带有 time.Duration 的字段。返回的 Field 将在适当的时候安全且显式地表示“nil”
	DurationP = zap.Durationp

	// Object 使用给定的键和 ObjectMarshaler 构造一个字段。它提供了一种灵活但仍然类型安全且高效的方法来将类似映射或结构的用户定义类型添加到日志记录上下文。该结构的 MarshalLogObject 方法被延迟调用
	Object = zap.Object

	// Inline 构造一个与 Object 类似的 Field，但它会将提供的 ObjectMarshaler 的元素添加到当前命名空间
	Inline = zap.Inline

	// Any 接受一个键和一个任意值，并选择将它们表示为字段的最佳方式，仅在必要时才回退到基于反射的方法。
	// 由于 byteuint8 和 runeint32 是别名，Any 无法区分它们。为了尽量减少意外情况，[]byte 值被视为二进制 blob，字节值被视为 uint8，而 runes 始终被视为整数
	Any = zap.Any

	// Err 是常见习语 NamedError("error", err) 的简写
	Err = zap.Error
)
