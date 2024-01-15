# Super

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
|[NewBitSet](#NewBitSet)|通过指定的 Bit 位创建一个 BitSet
|[TryWriteChannel](#TryWriteChannel)|尝试写入 channel，如果 channel 无法写入则忽略，返回是否写入成功
|[TryWriteChannelByHandler](#TryWriteChannelByHandler)|尝试写入 channel，如果 channel 无法写入则执行 handler
|[RegError](#RegError)|通过错误码注册错误，返回错误的引用
|[RegErrorRef](#RegErrorRef)|通过错误码注册错误，返回错误的引用
|[GetError](#GetError)|通过错误引用获取错误码和真实错误信息，如果错误不存在则返回 0，如果错误引用不存在则返回原本的错误
|[RecoverTransform](#RecoverTransform)|recover 错误转换
|[Handle](#Handle)|执行 f 函数，如果 f 为 nil，则不执行
|[HandleErr](#HandleErr)|执行 f 函数，如果 f 为 nil，则不执行
|[HandleV](#HandleV)|执行 f 函数，如果 f 为 nil，则不执行
|[GoFormat](#GoFormat)|go 代码格式化
|[If](#If)|暂无描述...
|[MarshalJSON](#MarshalJSON)|将对象转换为 json
|[MarshalJSONE](#MarshalJSONE)|将对象转换为 json
|[UnmarshalJSON](#UnmarshalJSON)|将 json 转换为对象
|[MarshalIndentJSON](#MarshalIndentJSON)|将对象转换为 json
|[MarshalToTargetWithJSON](#MarshalToTargetWithJSON)|将对象转换为目标对象
|[StartLossCounter](#StartLossCounter)|开始损耗计数
|[Match](#Match)|匹配
|[IsNumber](#IsNumber)|判断是否为数字
|[NumberToRome](#NumberToRome)|将数字转换为罗马数字
|[StringToInt](#StringToInt)|字符串转换为整数
|[StringToFloat64](#StringToFloat64)|字符串转换为 float64
|[StringToBool](#StringToBool)|字符串转换为 bool
|[StringToUint64](#StringToUint64)|字符串转换为 uint64
|[StringToUint](#StringToUint)|字符串转换为 uint
|[StringToFloat32](#StringToFloat32)|字符串转换为 float32
|[StringToInt64](#StringToInt64)|字符串转换为 int64
|[StringToUint32](#StringToUint32)|字符串转换为 uint32
|[StringToInt32](#StringToInt32)|字符串转换为 int32
|[StringToUint16](#StringToUint16)|字符串转换为 uint16
|[StringToInt16](#StringToInt16)|字符串转换为 int16
|[StringToUint8](#StringToUint8)|字符串转换为 uint8
|[StringToInt8](#StringToInt8)|字符串转换为 int8
|[StringToByte](#StringToByte)|字符串转换为 byte
|[StringToRune](#StringToRune)|字符串转换为 rune
|[IntToString](#IntToString)|整数转换为字符串
|[Float64ToString](#Float64ToString)|float64 转换为字符串
|[BoolToString](#BoolToString)|bool 转换为字符串
|[Uint64ToString](#Uint64ToString)|uint64 转换为字符串
|[UintToString](#UintToString)|uint 转换为字符串
|[Float32ToString](#Float32ToString)|float32 转换为字符串
|[Int64ToString](#Int64ToString)|int64 转换为字符串
|[Uint32ToString](#Uint32ToString)|uint32 转换为字符串
|[Int32ToString](#Int32ToString)|int32 转换为字符串
|[Uint16ToString](#Uint16ToString)|uint16 转换为字符串
|[Int16ToString](#Int16ToString)|int16 转换为字符串
|[Uint8ToString](#Uint8ToString)|uint8 转换为字符串
|[Int8ToString](#Int8ToString)|int8 转换为字符串
|[ByteToString](#ByteToString)|byte 转换为字符串
|[RuneToString](#RuneToString)|rune 转换为字符串
|[StringToSlice](#StringToSlice)|字符串转换为切片
|[SliceToString](#SliceToString)|切片转换为字符串
|[NewPermission](#NewPermission)|创建权限
|[Retry](#Retry)|根据提供的 count 次数尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功或者达到 count 次数
|[RetryByRule](#RetryByRule)|根据提供的规则尝试执行 f 函数，如果 f 函数返回错误，则根据 rule 的返回值进行重试
|[RetryByExponentialBackoff](#RetryByExponentialBackoff)|根据指数退避算法尝试执行 f 函数
|[ConditionalRetryByExponentialBackoff](#ConditionalRetryByExponentialBackoff)|该函数与 RetryByExponentialBackoff 类似，但是可以被中断
|[RetryAsync](#RetryAsync)|与 Retry 类似，但是是异步执行
|[RetryForever](#RetryForever)|根据提供的 interval 时间间隔尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功
|[NewStackGo](#NewStackGo)|返回一个用于获取上一个协程调用的堆栈信息的收集器
|[LaunchTime](#LaunchTime)|获取程序启动时间
|[Hostname](#Hostname)|获取主机名
|[PID](#PID)|获取进程 PID
|[StringToBytes](#StringToBytes)|以零拷贝的方式将字符串转换为字节切片
|[BytesToString](#BytesToString)|以零拷贝的方式将字节切片转换为字符串
|[Convert](#Convert)|以零拷贝的方式将一个对象转换为另一个对象
|[Verify](#Verify)|对特定表达式进行校验，当表达式不成立时，将执行 handle
|[OldVersion](#OldVersion)|检查 version2 对于 version1 来说是不是旧版本
|[CompareVersion](#CompareVersion)|返回一个整数，用于表示两个版本号的比较结果：


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[BitSet](#struct_BitSet)|是一个可以动态增长的比特位集合
|`STRUCT`|[LossCounter](#struct_LossCounter)|暂无描述...
|`STRUCT`|[Matcher](#struct_Matcher)|匹配器
|`STRUCT`|[Permission](#struct_Permission)|暂无描述...
|`STRUCT`|[StackGo](#struct_StackGo)|用于获取上一个协程调用的堆栈信息
|`STRUCT`|[VerifyHandle](#struct_VerifyHandle)|校验句柄

</details>


***
## 详情信息
#### func NewBitSet\[Bit generic.Integer\](bits ...Bit) *BitSet[Bit]
<span id="NewBitSet"></span>
> 通过指定的 Bit 位创建一个 BitSet

***
#### func TryWriteChannel\[T any\](ch chan T, data T) bool
<span id="TryWriteChannel"></span>
> 尝试写入 channel，如果 channel 无法写入则忽略，返回是否写入成功
>   - 无法写入的情况包括：channel 已满、channel 已关闭

***
#### func TryWriteChannelByHandler\[T any\](ch chan T, data T, handler func ())
<span id="TryWriteChannelByHandler"></span>
> 尝试写入 channel，如果 channel 无法写入则执行 handler
>   - 无法写入的情况包括：channel 已满、channel 已关闭

***
#### func RegError(code int, message string) error
<span id="RegError"></span>
> 通过错误码注册错误，返回错误的引用

***
#### func RegErrorRef(code int, message string, ref error) error
<span id="RegErrorRef"></span>
> 通过错误码注册错误，返回错误的引用
>   - 引用将会被重定向到注册的错误信息

***
#### func GetError(err error) (int,  error)
<span id="GetError"></span>
> 通过错误引用获取错误码和真实错误信息，如果错误不存在则返回 0，如果错误引用不存在则返回原本的错误

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGetError(t *testing.T) {
	var ErrNotFound = errors.New("not found")
	var _ = super.RegErrorRef(100, "test error", ErrNotFound)
	t.Log(super.GetError(ErrNotFound))
}

```


</details>


***
#### func RecoverTransform(a any) error
<span id="RecoverTransform"></span>
> recover 错误转换

**示例代码：**

```go

func ExampleRecoverTransform() {
	defer func() {
		if err := super.RecoverTransform(recover()); err != nil {
			fmt.Println(err)
		}
	}()
	panic("test")
}

```

***
#### func Handle(f func ())
<span id="Handle"></span>
> 执行 f 函数，如果 f 为 nil，则不执行

***
#### func HandleErr(f func ()  error) error
<span id="HandleErr"></span>
> 执行 f 函数，如果 f 为 nil，则不执行

***
#### func HandleV\[V any\](v V, f func (v V))
<span id="HandleV"></span>
> 执行 f 函数，如果 f 为 nil，则不执行

***
#### func GoFormat(filePath string)
<span id="GoFormat"></span>
> go 代码格式化

***
#### func If\[T any\](expression bool, t T, f T) T
<span id="If"></span>

***
#### func MarshalJSON(v interface {}) []byte
<span id="MarshalJSON"></span>
> 将对象转换为 json
>   - 当转换失败时，将返回 json 格式的空对象

***
#### func MarshalJSONE(v interface {}) ([]byte,  error)
<span id="MarshalJSONE"></span>
> 将对象转换为 json
>   - 当转换失败时，将返回错误信息

***
#### func UnmarshalJSON(data []byte, v interface {}) error
<span id="UnmarshalJSON"></span>
> 将 json 转换为对象

***
#### func MarshalIndentJSON(v interface {}, prefix string, indent string) []byte
<span id="MarshalIndentJSON"></span>
> 将对象转换为 json

***
#### func MarshalToTargetWithJSON(src interface {}, dest interface {}) error
<span id="MarshalToTargetWithJSON"></span>
> 将对象转换为目标对象

***
#### func StartLossCounter() *LossCounter
<span id="StartLossCounter"></span>
> 开始损耗计数

***
#### func Match\[Value any, Result any\](value Value) *Matcher[Value, Result]
<span id="Match"></span>
> 匹配

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestMatch(t *testing.T) {
	Convey("TestMatch", t, func() {
		So(super.Match[int, string](1).Case(1, "a").Case(2, "b").Default("c"), ShouldEqual, "a")
		So(super.Match[int, string](2).Case(1, "a").Case(2, "b").Default("c"), ShouldEqual, "b")
		So(super.Match[int, string](3).Case(1, "a").Case(2, "b").Default("c"), ShouldEqual, "c")
	})
}

```


</details>


***
#### func IsNumber(v any) bool
<span id="IsNumber"></span>
> 判断是否为数字

***
#### func NumberToRome(num int) string
<span id="NumberToRome"></span>
> 将数字转换为罗马数字

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNumberToRome(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{{input: 1, output: "I"}, {input: 5, output: "V"}, {input: 10, output: "X"}, {input: 50, output: "L"}, {input: 100, output: "C"}, {input: 500, output: "D"}, {input: 1000, output: "M"}}
	for _, test := range tests {
		result := super.NumberToRome(test.input)
		if result != test.output {
			t.Errorf("NumberToRome(%d) = %s; want %s", test.input, result, test.output)
		}
	}
}

```


</details>


***
#### func StringToInt(value string) int
<span id="StringToInt"></span>
> 字符串转换为整数

***
#### func StringToFloat64(value string) float64
<span id="StringToFloat64"></span>
> 字符串转换为 float64

***
#### func StringToBool(value string) bool
<span id="StringToBool"></span>
> 字符串转换为 bool

***
#### func StringToUint64(value string) uint64
<span id="StringToUint64"></span>
> 字符串转换为 uint64

***
#### func StringToUint(value string) uint
<span id="StringToUint"></span>
> 字符串转换为 uint

***
#### func StringToFloat32(value string) float32
<span id="StringToFloat32"></span>
> 字符串转换为 float32

***
#### func StringToInt64(value string) int64
<span id="StringToInt64"></span>
> 字符串转换为 int64

***
#### func StringToUint32(value string) uint32
<span id="StringToUint32"></span>
> 字符串转换为 uint32

***
#### func StringToInt32(value string) int32
<span id="StringToInt32"></span>
> 字符串转换为 int32

***
#### func StringToUint16(value string) uint16
<span id="StringToUint16"></span>
> 字符串转换为 uint16

***
#### func StringToInt16(value string) int16
<span id="StringToInt16"></span>
> 字符串转换为 int16

***
#### func StringToUint8(value string) uint8
<span id="StringToUint8"></span>
> 字符串转换为 uint8

***
#### func StringToInt8(value string) int8
<span id="StringToInt8"></span>
> 字符串转换为 int8

***
#### func StringToByte(value string) byte
<span id="StringToByte"></span>
> 字符串转换为 byte

***
#### func StringToRune(value string) rune
<span id="StringToRune"></span>
> 字符串转换为 rune

***
#### func IntToString(value int) string
<span id="IntToString"></span>
> 整数转换为字符串

***
#### func Float64ToString(value float64) string
<span id="Float64ToString"></span>
> float64 转换为字符串

***
#### func BoolToString(value bool) string
<span id="BoolToString"></span>
> bool 转换为字符串

***
#### func Uint64ToString(value uint64) string
<span id="Uint64ToString"></span>
> uint64 转换为字符串

***
#### func UintToString(value uint) string
<span id="UintToString"></span>
> uint 转换为字符串

***
#### func Float32ToString(value float32) string
<span id="Float32ToString"></span>
> float32 转换为字符串

***
#### func Int64ToString(value int64) string
<span id="Int64ToString"></span>
> int64 转换为字符串

***
#### func Uint32ToString(value uint32) string
<span id="Uint32ToString"></span>
> uint32 转换为字符串

***
#### func Int32ToString(value int32) string
<span id="Int32ToString"></span>
> int32 转换为字符串

***
#### func Uint16ToString(value uint16) string
<span id="Uint16ToString"></span>
> uint16 转换为字符串

***
#### func Int16ToString(value int16) string
<span id="Int16ToString"></span>
> int16 转换为字符串

***
#### func Uint8ToString(value uint8) string
<span id="Uint8ToString"></span>
> uint8 转换为字符串

***
#### func Int8ToString(value int8) string
<span id="Int8ToString"></span>
> int8 转换为字符串

***
#### func ByteToString(value byte) string
<span id="ByteToString"></span>
> byte 转换为字符串

***
#### func RuneToString(value rune) string
<span id="RuneToString"></span>
> rune 转换为字符串

***
#### func StringToSlice(value string) []string
<span id="StringToSlice"></span>
> 字符串转换为切片

***
#### func SliceToString(value []string) string
<span id="SliceToString"></span>
> 切片转换为字符串

***
#### func NewPermission\[Code generic.Integer, EntityID comparable\]() *Permission[Code, EntityID]
<span id="NewPermission"></span>
> 创建权限

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewPermission(t *testing.T) {
	const (
		Read = 1 << iota
		Write
		Execute
	)
	p := super.NewPermission[int, int]()
	p.AddPermission(1, Read, Write)
	t.Log(p.HasPermission(1, Read))
	t.Log(p.HasPermission(1, Write))
	p.SetPermission(2, Read|Write)
	t.Log(p.HasPermission(2, Read))
	t.Log(p.HasPermission(2, Execute))
	p.SetPermission(2, Execute)
	t.Log(p.HasPermission(2, Execute))
	t.Log(p.HasPermission(2, Read))
	t.Log(p.HasPermission(2, Write))
	p.RemovePermission(2, Execute)
	t.Log(p.HasPermission(2, Execute))
}

```


</details>


***
#### func Retry(count int, interval time.Duration, f func ()  error) error
<span id="Retry"></span>
> 根据提供的 count 次数尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功或者达到 count 次数

***
#### func RetryByRule(f func ()  error, rule func (count int)  time.Duration) error
<span id="RetryByRule"></span>
> 根据提供的规则尝试执行 f 函数，如果 f 函数返回错误，则根据 rule 的返回值进行重试
>   - rule 将包含一个入参，表示第几次重试，返回值表示下一次重试的时间间隔，当返回值为 0 时，表示不再重试
>   - rule 的 count 将在 f 首次失败后变为 1，因此 rule 的入参将从 1 开始

***
#### func RetryByExponentialBackoff(f func ()  error, maxRetries int, baseDelay time.Duration, maxDelay time.Duration, multiplier float64, randomization float64, ignoreErrors ...error) error
<span id="RetryByExponentialBackoff"></span>
> 根据指数退避算法尝试执行 f 函数
>   - maxRetries：最大重试次数
>   - baseDelay：基础延迟
>   - maxDelay：最大延迟
>   - multiplier：延迟时间的乘数，通常为 2
>   - randomization：延迟时间的随机化因子，通常为 0.5
>   - ignoreErrors：忽略的错误，当 f 返回的错误在 ignoreErrors 中时，将不会进行重试

***
#### func ConditionalRetryByExponentialBackoff(f func ()  error, cond func ()  bool, maxRetries int, baseDelay time.Duration, maxDelay time.Duration, multiplier float64, randomization float64, ignoreErrors ...error) error
<span id="ConditionalRetryByExponentialBackoff"></span>
> 该函数与 RetryByExponentialBackoff 类似，但是可以被中断
>   - cond 为中断条件，当 cond 返回 false 时，将会中断重试
> 
> 该函数通常用于在重试过程中，需要中断重试的场景，例如：
>   - 用户请求开始游戏，由于网络等情况，进入重试状态。此时用户再次发送开始游戏请求，此时需要中断之前的重试，避免重复进入游戏

***
#### func RetryAsync(count int, interval time.Duration, f func ()  error, callback func (err error))
<span id="RetryAsync"></span>
> 与 Retry 类似，但是是异步执行
>   - 传入的 callback 函数会在执行完毕后被调用，如果执行成功，则 err 为 nil，否则为错误信息
>   - 如果 callback 为 nil，则不会在执行完毕后被调用

***
#### func RetryForever(interval time.Duration, f func ()  error)
<span id="RetryForever"></span>
> 根据提供的 interval 时间间隔尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功

***
#### func NewStackGo() *StackGo
<span id="NewStackGo"></span>
> 返回一个用于获取上一个协程调用的堆栈信息的收集器

***
#### func LaunchTime() time.Time
<span id="LaunchTime"></span>
> 获取程序启动时间

***
#### func Hostname() string
<span id="Hostname"></span>
> 获取主机名

***
#### func PID() int
<span id="PID"></span>
> 获取进程 PID

***
#### func StringToBytes(s string) []byte
<span id="StringToBytes"></span>
> 以零拷贝的方式将字符串转换为字节切片

***
#### func BytesToString(b []byte) string
<span id="BytesToString"></span>
> 以零拷贝的方式将字节切片转换为字符串

***
#### func Convert\[A any, B any\](src A) B
<span id="Convert"></span>
> 以零拷贝的方式将一个对象转换为另一个对象
>   - 两个对象字段必须完全一致
>   - 该函数可以绕过私有字段的访问限制

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestConvert(t *testing.T) {
	type B struct {
		nocmp [0]func()
		v     atomic.Value
	}
	var a = atomic.NewString("hello")
	var b = super.Convert[*atomic.String, *B](a)
	fmt.Println(b.v.Load())
}

```


</details>


***
#### func Verify\[V any\](handle func ( V)) *VerifyHandle[V]
<span id="Verify"></span>
> 对特定表达式进行校验，当表达式不成立时，将执行 handle

**示例代码：**

```go

func ExampleVerify() {
	var getId = func() int {
		return 1
	}
	var n *super.VerifyHandle[int]
	super.Verify(func(err error) {
		fmt.Println(err)
	}).Case(getId() == 1, errors.New("id can't be 1")).Do()
	super.Verify(func(err error) {
		fmt.Println(err)
	}).PreCase(func() bool {
		return n == nil
	}, errors.New("n can't be nil"), func(verify *super.VerifyHandle[error]) bool {
		return verify.Do()
	})
}

```

***
#### func OldVersion(version1 string, version2 string) bool
<span id="OldVersion"></span>
> 检查 version2 对于 version1 来说是不是旧版本

**示例代码：**

```go

func ExampleOldVersion() {
	result := super.OldVersion("1.2.3", "1.2.2")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestOldVersion(t *testing.T) {
	testCases := []struct {
		version1 string
		version2 string
		want     bool
	}{{"1.2.3", "1.2.2", true}, {"1.2.1", "1.2.2", false}, {"1.2.3", "1.2.3", false}, {"v1.2.3", "v1.2.2", true}, {"v1.2.3", "v1.2.4", false}, {"v1.2.3", "1.2.3", false}, {"vxx2faf.d2ad5.dd3", "gga2faf.d2ad5.dd2", true}, {"awd2faf.d2ad4.dd3", "vsd2faf.d2ad5.dd3", false}, {"vxd2faf.d2ad5.dd3", "qdq2faf.d2ad5.dd3", false}, {"1.2.3", "vdafe2faf.d2ad5.dd3", false}, {"v1.2.3", "vdafe2faf.d2ad5.dd3", false}}
	for _, tc := range testCases {
		got := super.OldVersion(tc.version1, tc.version2)
		if got != tc.want {
			t.Errorf("OldVersion(%q, %q) = %v; want %v", tc.version1, tc.version2, got, tc.want)
		}
	}
}

```


</details>


<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkOldVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		super.OldVersion("vfe2faf.d2ad5.dd3", "vda2faf.d2ad5.dd2")
	}
}

```


</details>


***
#### func CompareVersion(version1 string, version2 string) int
<span id="CompareVersion"></span>
> 返回一个整数，用于表示两个版本号的比较结果：
>   - 如果 version1 大于 version2，它将返回 1
>   - 如果 version1 小于 version2，它将返回 -1
>   - 如果 version1 和 version2 相等，它将返回 0

**示例代码：**

```go

func ExampleCompareVersion() {
	result := super.CompareVersion("1.2.3", "1.2.2")
	fmt.Println(result)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		version1 string
		version2 string
		want     int
	}{{"1.2.3", "1.2.2", 1}, {"1.2.1", "1.2.2", -1}, {"1.2.3", "1.2.3", 0}, {"v1.2.3", "v1.2.2", 1}, {"v1.2.3", "v1.2.4", -1}, {"v1.2.3", "1.2.3", 0}, {"vde2faf.d2ad5.dd3", "e2faf.d2ad5.dd2", 1}, {"vde2faf.d2ad4.dd3", "vde2faf.d2ad5.dd3", -1}, {"vfe2faf.d2ad5.dd3", "ve2faf.d2ad5.dd3", 0}, {"1.2.3", "vdafe2faf.d2ad5.dd3", -1}, {"v1.2.3", "vdafe2faf.d2ad5.dd3", -1}}
	for _, tc := range testCases {
		got := super.CompareVersion(tc.version1, tc.version2)
		if got != tc.want {
			t.Errorf("CompareVersion(%q, %q) = %v; want %v", tc.version1, tc.version2, got, tc.want)
		}
	}
}

```


</details>


<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkCompareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		super.CompareVersion("vfe2faf.d2ad5.dd3", "afe2faf.d2ad5.dd2")
	}
}

```


</details>


***
<span id="struct_BitSet"></span>
### BitSet `STRUCT`
是一个可以动态增长的比特位集合
  - 默认情况下将使用 64 位无符号整数来表示比特位，当需要表示的比特位超过 64 位时，将自动增长
```go
type BitSet[Bit generic.Integer] struct {
	set []uint64
}
```
#### func (*BitSet) Set(bit Bit)  *BitSet[Bit]
> 将指定的位 bit 设置为 1
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestBitSet_Set(t *testing.T) {
	bs := super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	bs.Set(11)
	bs.Set(12)
	bs.Set(13)
	t.Log(bs)
}

```


</details>


***
#### func (*BitSet) Del(bit Bit)  *BitSet[Bit]
> 将指定的位 bit 设置为 0
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestBitSet_Del(t *testing.T) {
	bs := super.NewBitSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	bs.Del(11)
	bs.Del(12)
	bs.Del(13)
	bs.Del(10)
	t.Log(bs)
}

```


</details>


***
#### func (*BitSet) Shrink()  *BitSet[Bit]
> 将 BitSet 中的比特位集合缩小到最小
>   - 正常情况下当 BitSet 中的比特位超出 64 位时，将自动增长，当 BitSet 中的比特位数量减少时，可以使用该方法将 BitSet 中的比特位集合缩小到最小
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestBitSet_Shrink(t *testing.T) {
	bs := super.NewBitSet(63)
	t.Log(bs.Cap())
	bs.Set(200)
	t.Log(bs.Cap())
	bs.Del(200)
	bs.Shrink()
	t.Log(bs.Cap())
}

```


</details>


***
#### func (*BitSet) Cap()  int
> 返回当前 BitSet 中可以表示的最大比特位数量
***
#### func (*BitSet) Has(bit Bit)  bool
> 检查指定的位 bit 是否被设置为 1
***
#### func (*BitSet) Clear()  *BitSet[Bit]
> 清空所有的比特位
***
#### func (*BitSet) Len()  int
> 返回当前 BitSet 中被设置的比特位数量
***
#### func (*BitSet) Bits()  []Bit
> 返回当前 BitSet 中被设置的比特位
***
#### func (*BitSet) Reverse()  *BitSet[Bit]
> 反转当前 BitSet 中的所有比特位
***
#### func (*BitSet) Not()  *BitSet[Bit]
> 返回当前 BitSet 中所有比特位的反转
***
#### func (*BitSet) And(other *BitSet[Bit])  *BitSet[Bit]
> 将当前 BitSet 与另一个 BitSet 进行按位与运算
***
#### func (*BitSet) Or(other *BitSet[Bit])  *BitSet[Bit]
> 将当前 BitSet 与另一个 BitSet 进行按位或运算
***
#### func (*BitSet) Xor(other *BitSet[Bit])  *BitSet[Bit]
> 将当前 BitSet 与另一个 BitSet 进行按位异或运算
***
#### func (*BitSet) Sub(other *BitSet[Bit])  *BitSet[Bit]
> 将当前 BitSet 与另一个 BitSet 进行按位减运算
***
#### func (*BitSet) IsZero()  bool
> 检查当前 BitSet 是否为空
***
#### func (*BitSet) Clone()  *BitSet[Bit]
> 返回当前 BitSet 的副本
***
#### func (*BitSet) Equal(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否与另一个 BitSet 相等
***
#### func (*BitSet) Contains(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否包含另一个 BitSet
***
#### func (*BitSet) ContainsAny(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否包含另一个 BitSet 中的任意比特位
***
#### func (*BitSet) ContainsAll(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否包含另一个 BitSet 中的所有比特位
***
#### func (*BitSet) Intersect(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否与另一个 BitSet 有交集
***
#### func (*BitSet) Union(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否与另一个 BitSet 有并集
***
#### func (*BitSet) Difference(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否与另一个 BitSet 有差集
***
#### func (*BitSet) SymmetricDifference(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否与另一个 BitSet 有对称差集
***
#### func (*BitSet) Subset(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否为另一个 BitSet 的子集
***
#### func (*BitSet) Superset(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否为另一个 BitSet 的超集
***
#### func (*BitSet) Complement(other *BitSet[Bit])  bool
> 检查当前 BitSet 是否为另一个 BitSet 的补集
***
#### func (*BitSet) Max()  Bit
> 返回当前 BitSet 中最大的比特位
***
#### func (*BitSet) Min()  Bit
> 返回当前 BitSet 中最小的比特位
***
#### func (*BitSet) String()  string
> 返回当前 BitSet 的字符串表示
***
#### func (*BitSet) MarshalJSON() ( []byte,  error)
> 实现 json.Marshaler 接口
***
#### func (*BitSet) UnmarshalJSON(data []byte)  error
> 实现 json.Unmarshaler 接口
***
<span id="struct_LossCounter"></span>
### LossCounter `STRUCT`

```go
type LossCounter struct {
	curr    time.Time
	loss    []time.Duration
	lossKey []string
}
```
#### func (*LossCounter) Record(name string)
> 记录一次损耗
***
#### func (*LossCounter) GetLoss(handler func (step int, name string, loss time.Duration))
> 获取损耗
***
#### func (*LossCounter) String()  string
***
<span id="struct_Matcher"></span>
### Matcher `STRUCT`
匹配器
```go
type Matcher[Value any, Result any] struct {
	value Value
	r     Result
	d     bool
}
```
#### func (*Matcher) Case(value Value, result Result)  *Matcher[Value, Result]
> 匹配
***
#### func (*Matcher) Default(value Result)  Result
> 默认
***
<span id="struct_Permission"></span>
### Permission `STRUCT`

```go
type Permission[Code generic.Integer, EntityID comparable] struct {
	permissions map[EntityID]Code
	l           sync.RWMutex
}
```
#### func (*Permission) HasPermission(entityId EntityID, permission Code)  bool
> 是否有权限
***
#### func (*Permission) AddPermission(entityId EntityID, permission ...Code)
> 添加权限
***
#### func (*Permission) RemovePermission(entityId EntityID, permission ...Code)
> 移除权限
***
#### func (*Permission) SetPermission(entityId EntityID, permission ...Code)
> 设置权限
***
<span id="struct_StackGo"></span>
### StackGo `STRUCT`
用于获取上一个协程调用的堆栈信息
  - 应当最先运行 Wait 函数，然后在其他协程中调用 Stack 函数或者 GiveUp 函数
  - 适用于跨协程同步通讯，例如单线程的消息处理统计耗时打印堆栈信息
```go
type StackGo struct {
	stack   chan *struct{}
	collect chan []byte
}
```
#### func (*StackGo) Wait()
> 等待收集消息堆栈
>   - 在调用 Wait 函数后，当前协程将会被挂起，直到调用 Stack 或 GiveUp 函数
***
#### func (*StackGo) Stack()  []byte
> 获取消息堆栈
>   - 在调用 Wait 函数后调用该函数，将会返回上一个协程的堆栈信息
>   - 在调用 GiveUp 函数后调用该函数，将会 panic
***
#### func (*StackGo) GiveUp()
> 放弃收集消息堆栈
>   - 在调用 Wait 函数后调用该函数，将会放弃收集消息堆栈并且释放资源
>   - 在调用 GiveUp 函数后调用 Stack 函数，将会 panic
***
<span id="struct_VerifyHandle"></span>
### VerifyHandle `STRUCT`
校验句柄
```go
type VerifyHandle[V any] struct {
	handle func(V)
	v      V
	hit    bool
}
```
#### func (*VerifyHandle) PreCase(expression func ()  bool, value V, caseHandle func (verify *VerifyHandle[V])  bool)  bool
> 先决校验用例，当 expression 成立时，将跳过 caseHandle 的执行，直接执行 handle 并返回 false
>   - 常用于对前置参数的空指针校验，例如当 a 为 nil 时，不执行 a.B()，而是直接返回 false
***
#### func (*VerifyHandle) Case(expression bool, value V)  *VerifyHandle[V]
> 校验用例，当 expression 成立时，将忽略后续 Case，并将在 Do 时执行 handle，返回 false
***
#### func (*VerifyHandle) Do()  bool
> 执行校验，当校验失败时，将执行 handle，并返回 false
***
