# Log

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[CallerBasicFormat](#CallerBasicFormat)|返回调用者的基本格式
|[Println](#Println)|暂无描述...
|[Default](#Default)|获取默认的日志记录器
|[SetDefault](#SetDefault)|设置默认的日志记录器
|[SetDefaultBySlog](#SetDefaultBySlog)|设置默认的日志记录器
|[Debug](#Debug)|在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Info](#Info)|在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Warn](#Warn)|在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Error](#Error)|在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[DPanic](#DPanic)|在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Panic](#Panic)|在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Fatal](#Fatal)|在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
|[Skip](#Skip)|构造一个无操作字段，这在处理其他 Field 构造函数中的无效输入时通常很有用
|[Duration](#Duration)|使用给定的键和值构造一个字段。编码器控制持续时间的序列化方式
|[DurationP](#DurationP)|构造一个带有 time.Duration 的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Bool](#Bool)|构造一个带有布尔值的字段
|[BoolP](#BoolP)|构造一个带有布尔值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[String](#String)|构造一个带有字符串值的字段
|[StringP](#StringP)|构造一个带有字符串值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Int](#Int)|构造一个带有整数值的字段
|[IntP](#IntP)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Int8](#Int8)|构造一个带有整数值的字段
|[Int8P](#Int8P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Int16](#Int16)|构造一个带有整数值的字段
|[Int16P](#Int16P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Int32](#Int32)|构造一个带有整数值的字段
|[Int32P](#Int32P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Int64](#Int64)|构造一个带有整数值的字段
|[Int64P](#Int64P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Uint](#Uint)|构造一个带有整数值的字段
|[UintP](#UintP)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Uint8](#Uint8)|构造一个带有整数值的字段
|[Uint8P](#Uint8P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Uint16](#Uint16)|构造一个带有整数值的字段
|[Uint16P](#Uint16P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Uint32](#Uint32)|构造一个带有整数值的字段
|[Uint32P](#Uint32P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Uint64](#Uint64)|构造一个带有整数值的字段
|[Uint64P](#Uint64P)|构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Float](#Float)|构造一个带有浮点值的字段
|[FloatP](#FloatP)|构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Float32](#Float32)|构造一个带有浮点值的字段
|[Float32P](#Float32P)|构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Float64](#Float64)|构造一个带有浮点值的字段
|[Float64P](#Float64P)|构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Time](#Time)|构造一个带有时间值的字段
|[TimeP](#TimeP)|构造一个带有时间值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"
|[Any](#Any)|构造一个带有任意值的字段
|[Group](#Group)|返回分组字段
|[Stack](#Stack)|返回堆栈字段
|[Err](#Err)|构造一个带有错误值的字段
|[NewHandler](#NewHandler)|创建一个更偏向于人类可读的处理程序，该处理程序也是默认的处理程序
|[NewLogger](#NewLogger)|创建一个新的日志记录器
|[NewMultiHandler](#NewMultiHandler)|创建一个新的多处理程序
|[NewOptions](#NewOptions)|创建一个新的日志选项


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Field](#struct_Field)|暂无描述...
|`STRUCT`|[Logger](#struct_Logger)|暂无描述...
|`STRUCT`|[MultiHandler](#struct_MultiHandler)|暂无描述...
|`STRUCT`|[Option](#struct_Option)|暂无描述...

</details>


***
## 详情信息
#### func CallerBasicFormat(file string, line int) (repFile string, refLine string)
<span id="CallerBasicFormat"></span>
> 返回调用者的基本格式

***
#### func Println(str string, color string, desc string)
<span id="Println"></span>

***
#### func Default() *Logger
<span id="Default"></span>
> 获取默认的日志记录器

***
#### func SetDefault(l *Logger)
<span id="SetDefault"></span>
> 设置默认的日志记录器

***
#### func SetDefaultBySlog(l *slog.Logger)
<span id="SetDefaultBySlog"></span>
> 设置默认的日志记录器

***
#### func Debug(msg string, args ...any)
<span id="Debug"></span>
> 在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段

***
#### func Info(msg string, args ...any)
<span id="Info"></span>
> 在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段

***
#### func Warn(msg string, args ...any)
<span id="Warn"></span>
> 在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段

***
#### func Error(msg string, args ...any)
<span id="Error"></span>
> 在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段

***
#### func DPanic(msg string, args ...any)
<span id="DPanic"></span>
> 在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
>   - 如果记录器处于开发模式，它就会出现 panic（DPanic 的意思是“development panic”）。这对于捕获可恢复但不应该发生的错误很有用
>   - 该 panic 仅在 NewHandler 中创建的处理器会生效

***
#### func Panic(msg string, args ...any)
<span id="Panic"></span>
> 在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
>   - 即使禁用了 PanicLevel 的日志记录，记录器也会出现 panic

***
#### func Fatal(msg string, args ...any)
<span id="Fatal"></span>
> 在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
>   - 然后记录器调用 os.Exit(1)，即使 FatalLevel 的日志记录被禁用

***
#### func Skip(vs ...any) slog.Attr
<span id="Skip"></span>
> 构造一个无操作字段，这在处理其他 Field 构造函数中的无效输入时通常很有用
>   - 该函数还支持将其他字段快捷的转换为 Skip 字段

***
#### func Duration(key string, val time.Duration) slog.Attr
<span id="Duration"></span>
> 使用给定的键和值构造一个字段。编码器控制持续时间的序列化方式

***
#### func DurationP(key string, val *time.Duration) slog.Attr
<span id="DurationP"></span>
> 构造一个带有 time.Duration 的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Bool(key string, val bool) slog.Attr
<span id="Bool"></span>
> 构造一个带有布尔值的字段

***
#### func BoolP(key string, val *bool) slog.Attr
<span id="BoolP"></span>
> 构造一个带有布尔值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func String(key string, val string) slog.Attr
<span id="String"></span>
> 构造一个带有字符串值的字段

***
#### func StringP(key string, val *string) slog.Attr
<span id="StringP"></span>
> 构造一个带有字符串值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Int\[I generic.Integer\](key string, val I) slog.Attr
<span id="Int"></span>
> 构造一个带有整数值的字段

***
#### func IntP\[I generic.Integer\](key string, val *I) slog.Attr
<span id="IntP"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Int8\[I generic.Integer\](key string, val I) slog.Attr
<span id="Int8"></span>
> 构造一个带有整数值的字段

***
#### func Int8P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Int8P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Int16\[I generic.Integer\](key string, val I) slog.Attr
<span id="Int16"></span>
> 构造一个带有整数值的字段

***
#### func Int16P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Int16P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Int32\[I generic.Integer\](key string, val I) slog.Attr
<span id="Int32"></span>
> 构造一个带有整数值的字段

***
#### func Int32P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Int32P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Int64\[I generic.Integer\](key string, val I) slog.Attr
<span id="Int64"></span>
> 构造一个带有整数值的字段

***
#### func Int64P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Int64P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Uint\[I generic.Integer\](key string, val I) slog.Attr
<span id="Uint"></span>
> 构造一个带有整数值的字段

***
#### func UintP\[I generic.Integer\](key string, val *I) slog.Attr
<span id="UintP"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Uint8\[I generic.Integer\](key string, val I) slog.Attr
<span id="Uint8"></span>
> 构造一个带有整数值的字段

***
#### func Uint8P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Uint8P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Uint16\[I generic.Integer\](key string, val I) slog.Attr
<span id="Uint16"></span>
> 构造一个带有整数值的字段

***
#### func Uint16P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Uint16P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Uint32\[I generic.Integer\](key string, val I) slog.Attr
<span id="Uint32"></span>
> 构造一个带有整数值的字段

***
#### func Uint32P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Uint32P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Uint64\[I generic.Integer\](key string, val I) slog.Attr
<span id="Uint64"></span>
> 构造一个带有整数值的字段

***
#### func Uint64P\[I generic.Integer\](key string, val *I) slog.Attr
<span id="Uint64P"></span>
> 构造一个带有整数值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Float\[F generic.Float\](key string, val F) slog.Attr
<span id="Float"></span>
> 构造一个带有浮点值的字段

***
#### func FloatP\[F generic.Float\](key string, val *F) slog.Attr
<span id="FloatP"></span>
> 构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Float32\[F generic.Float\](key string, val F) slog.Attr
<span id="Float32"></span>
> 构造一个带有浮点值的字段

***
#### func Float32P\[F generic.Float\](key string, val *F) slog.Attr
<span id="Float32P"></span>
> 构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Float64\[F generic.Float\](key string, val F) slog.Attr
<span id="Float64"></span>
> 构造一个带有浮点值的字段

***
#### func Float64P\[F generic.Float\](key string, val *F) slog.Attr
<span id="Float64P"></span>
> 构造一个带有浮点值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Time(key string, val time.Time) slog.Attr
<span id="Time"></span>
> 构造一个带有时间值的字段

***
#### func TimeP(key string, val *time.Time) slog.Attr
<span id="TimeP"></span>
> 构造一个带有时间值的字段。返回的 Field 将在适当的时候安全且显式地表示 "null"

***
#### func Any(key string, val any) slog.Attr
<span id="Any"></span>
> 构造一个带有任意值的字段

***
#### func Group(key string, args ...any) slog.Attr
<span id="Group"></span>
> 返回分组字段

***
#### func Stack(key string) slog.Attr
<span id="Stack"></span>
> 返回堆栈字段

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestStack(t *testing.T) {
	var i int
	for {
		time.Sleep(time.Second)
		Debug("TestStack")
		Info("TestStack")
		Warn("TestStack")
		Error("TestStack")
		i++
		if i == 3 {
			Default().Logger.Handler().(*handler).opts.GerRuntimeHandler().ChangeLevel(slog.LevelInfo)
		}
	}
}

```


</details>


***
#### func Err(err error) slog.Attr
<span id="Err"></span>
> 构造一个带有错误值的字段

***
#### func NewHandler(w io.Writer, opts *Options) slog.Handler
<span id="NewHandler"></span>
> 创建一个更偏向于人类可读的处理程序，该处理程序也是默认的处理程序

***
#### func NewLogger(handlers ...slog.Handler) *Logger
<span id="NewLogger"></span>
> 创建一个新的日志记录器

***
#### func NewMultiHandler(handlers ...slog.Handler) slog.Handler
<span id="NewMultiHandler"></span>
> 创建一个新的多处理程序

***
#### func NewOptions() *Options
<span id="NewOptions"></span>
> 创建一个新的日志选项

***
<span id="struct_Field"></span>
### Field `STRUCT`

```go
type Field slog.Attr
```
<span id="struct_Logger"></span>
### Logger `STRUCT`

```go
type Logger struct {
	*slog.Logger
}
```
<span id="struct_MultiHandler"></span>
### MultiHandler `STRUCT`

```go
type MultiHandler struct {
	handlers []slog.Handler
}
```
<span id="struct_MultiHandler_Enabled"></span>

#### func (MultiHandler) Enabled(ctx context.Context, level slog.Level)  bool

***
<span id="struct_MultiHandler_Handle"></span>

#### func (MultiHandler) Handle(ctx context.Context, record slog.Record) (err error)

***
<span id="struct_MultiHandler_WithAttrs"></span>

#### func (MultiHandler) WithAttrs(attrs []slog.Attr)  slog.Handler

***
<span id="struct_MultiHandler_WithGroup"></span>

#### func (MultiHandler) WithGroup(name string)  slog.Handler

***
<span id="struct_Option"></span>
### Option `STRUCT`

```go
type Option func(opts *Options)
```
