# Super



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/super)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
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


> 结构体定义

|结构体|描述
|:--|:--
|[BitSet](#bitset)|是一个可以动态增长的比特位集合
|[LossCounter](#losscounter)|暂无描述...
|[Matcher](#matcher)|匹配器
|[Permission](#permission)|暂无描述...
|[StackGo](#stackgo)|用于获取上一个协程调用的堆栈信息
|[VerifyHandle](#verifyhandle)|校验句柄

</details>


#### func NewBitSet(bits ...Bit)  *BitSet[Bit]
<span id="NewBitSet"></span>
> 通过指定的 Bit 位创建一个 BitSet
***
#### func TryWriteChannel(ch chan T, data T)  bool
<span id="TryWriteChannel"></span>
> 尝试写入 channel，如果 channel 无法写入则忽略，返回是否写入成功
>   - 无法写入的情况包括：channel 已满、channel 已关闭
***
#### func TryWriteChannelByHandler(ch chan T, data T, handler func ())
<span id="TryWriteChannelByHandler"></span>
> 尝试写入 channel，如果 channel 无法写入则执行 handler
>   - 无法写入的情况包括：channel 已满、channel 已关闭
***
#### func RegError(code int, message string)  error
<span id="RegError"></span>
> 通过错误码注册错误，返回错误的引用
***
#### func RegErrorRef(code int, message string, ref error)  error
<span id="RegErrorRef"></span>
> 通过错误码注册错误，返回错误的引用
>   - 引用将会被重定向到注册的错误信息
***
#### func GetError(err error)  int,  error
<span id="GetError"></span>
> 通过错误引用获取错误码和真实错误信息，如果错误不存在则返回 0，如果错误引用不存在则返回原本的错误
***
#### func RecoverTransform(a any)  error
<span id="RecoverTransform"></span>
> recover 错误转换
***
#### func Handle(f func ())
<span id="Handle"></span>
> 执行 f 函数，如果 f 为 nil，则不执行
***
#### func HandleErr(f func ()  error)  error
<span id="HandleErr"></span>
> 执行 f 函数，如果 f 为 nil，则不执行
***
#### func HandleV(v V, f func (v V))
<span id="HandleV"></span>
> 执行 f 函数，如果 f 为 nil，则不执行
***
#### func GoFormat(filePath string)
<span id="GoFormat"></span>
> go 代码格式化
***
#### func If(expression bool, t T, f T)  T
<span id="If"></span>
***
#### func MarshalJSON(v interface {})  []byte
<span id="MarshalJSON"></span>
> 将对象转换为 json
>   - 当转换失败时，将返回 json 格式的空对象
***
#### func MarshalJSONE(v interface {})  []byte,  error
<span id="MarshalJSONE"></span>
> 将对象转换为 json
>   - 当转换失败时，将返回错误信息
***
#### func UnmarshalJSON(data []byte, v interface {})  error
<span id="UnmarshalJSON"></span>
> 将 json 转换为对象
***
#### func MarshalIndentJSON(v interface {}, prefix string, indent string)  []byte
<span id="MarshalIndentJSON"></span>
> 将对象转换为 json
***
#### func MarshalToTargetWithJSON(src interface {}, dest interface {})  error
<span id="MarshalToTargetWithJSON"></span>
> 将对象转换为目标对象
***
#### func StartLossCounter()  *LossCounter
<span id="StartLossCounter"></span>
> 开始损耗计数
***
#### func Match(value Value)  *Matcher[Value, Result]
<span id="Match"></span>
> 匹配
***
#### func IsNumber(v any)  bool
<span id="IsNumber"></span>
> 判断是否为数字
***
#### func NumberToRome(num int)  string
<span id="NumberToRome"></span>
> 将数字转换为罗马数字
***
#### func StringToInt(value string)  int
<span id="StringToInt"></span>
> 字符串转换为整数
***
#### func StringToFloat64(value string)  float64
<span id="StringToFloat64"></span>
> 字符串转换为 float64
***
#### func StringToBool(value string)  bool
<span id="StringToBool"></span>
> 字符串转换为 bool
***
#### func StringToUint64(value string)  uint64
<span id="StringToUint64"></span>
> 字符串转换为 uint64
***
#### func StringToUint(value string)  uint
<span id="StringToUint"></span>
> 字符串转换为 uint
***
#### func StringToFloat32(value string)  float32
<span id="StringToFloat32"></span>
> 字符串转换为 float32
***
#### func StringToInt64(value string)  int64
<span id="StringToInt64"></span>
> 字符串转换为 int64
***
#### func StringToUint32(value string)  uint32
<span id="StringToUint32"></span>
> 字符串转换为 uint32
***
#### func StringToInt32(value string)  int32
<span id="StringToInt32"></span>
> 字符串转换为 int32
***
#### func StringToUint16(value string)  uint16
<span id="StringToUint16"></span>
> 字符串转换为 uint16
***
#### func StringToInt16(value string)  int16
<span id="StringToInt16"></span>
> 字符串转换为 int16
***
#### func StringToUint8(value string)  uint8
<span id="StringToUint8"></span>
> 字符串转换为 uint8
***
#### func StringToInt8(value string)  int8
<span id="StringToInt8"></span>
> 字符串转换为 int8
***
#### func StringToByte(value string)  byte
<span id="StringToByte"></span>
> 字符串转换为 byte
***
#### func StringToRune(value string)  rune
<span id="StringToRune"></span>
> 字符串转换为 rune
***
#### func IntToString(value int)  string
<span id="IntToString"></span>
> 整数转换为字符串
***
#### func Float64ToString(value float64)  string
<span id="Float64ToString"></span>
> float64 转换为字符串
***
#### func BoolToString(value bool)  string
<span id="BoolToString"></span>
> bool 转换为字符串
***
#### func Uint64ToString(value uint64)  string
<span id="Uint64ToString"></span>
> uint64 转换为字符串
***
#### func UintToString(value uint)  string
<span id="UintToString"></span>
> uint 转换为字符串
***
#### func Float32ToString(value float32)  string
<span id="Float32ToString"></span>
> float32 转换为字符串
***
#### func Int64ToString(value int64)  string
<span id="Int64ToString"></span>
> int64 转换为字符串
***
#### func Uint32ToString(value uint32)  string
<span id="Uint32ToString"></span>
> uint32 转换为字符串
***
#### func Int32ToString(value int32)  string
<span id="Int32ToString"></span>
> int32 转换为字符串
***
#### func Uint16ToString(value uint16)  string
<span id="Uint16ToString"></span>
> uint16 转换为字符串
***
#### func Int16ToString(value int16)  string
<span id="Int16ToString"></span>
> int16 转换为字符串
***
#### func Uint8ToString(value uint8)  string
<span id="Uint8ToString"></span>
> uint8 转换为字符串
***
#### func Int8ToString(value int8)  string
<span id="Int8ToString"></span>
> int8 转换为字符串
***
#### func ByteToString(value byte)  string
<span id="ByteToString"></span>
> byte 转换为字符串
***
#### func RuneToString(value rune)  string
<span id="RuneToString"></span>
> rune 转换为字符串
***
#### func StringToSlice(value string)  []string
<span id="StringToSlice"></span>
> 字符串转换为切片
***
#### func SliceToString(value []string)  string
<span id="SliceToString"></span>
> 切片转换为字符串
***
#### func NewPermission()  *Permission[Code, EntityID]
<span id="NewPermission"></span>
> 创建权限
***
#### func Retry(count int, interval time.Duration, f func ()  error)  error
<span id="Retry"></span>
> 根据提供的 count 次数尝试执行 f 函数，如果 f 函数返回错误，则在 interval 后重试，直到成功或者达到 count 次数
***
#### func RetryByRule(f func ()  error, rule func (count int)  time.Duration)  error
<span id="RetryByRule"></span>
> 根据提供的规则尝试执行 f 函数，如果 f 函数返回错误，则根据 rule 的返回值进行重试
>   - rule 将包含一个入参，表示第几次重试，返回值表示下一次重试的时间间隔，当返回值为 0 时，表示不再重试
>   - rule 的 count 将在 f 首次失败后变为 1，因此 rule 的入参将从 1 开始
***
#### func RetryByExponentialBackoff(f func ()  error, maxRetries int, baseDelay time.Duration, maxDelay time.Duration, multiplier float64, randomization float64, ignoreErrors ...error)  error
<span id="RetryByExponentialBackoff"></span>
> 根据指数退避算法尝试执行 f 函数
>   - maxRetries：最大重试次数
>   - baseDelay：基础延迟
>   - maxDelay：最大延迟
>   - multiplier：延迟时间的乘数，通常为 2
>   - randomization：延迟时间的随机化因子，通常为 0.5
>   - ignoreErrors：忽略的错误，当 f 返回的错误在 ignoreErrors 中时，将不会进行重试
***
#### func ConditionalRetryByExponentialBackoff(f func ()  error, cond func ()  bool, maxRetries int, baseDelay time.Duration, maxDelay time.Duration, multiplier float64, randomization float64, ignoreErrors ...error)  error
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
#### func NewStackGo()  *StackGo
<span id="NewStackGo"></span>
> 返回一个用于获取上一个协程调用的堆栈信息的收集器
***
#### func LaunchTime()  time.Time
<span id="LaunchTime"></span>
> 获取程序启动时间
***
#### func Hostname()  string
<span id="Hostname"></span>
> 获取主机名
***
#### func PID()  int
<span id="PID"></span>
> 获取进程 PID
***
#### func StringToBytes(s string)  []byte
<span id="StringToBytes"></span>
> 以零拷贝的方式将字符串转换为字节切片
***
#### func BytesToString(b []byte)  string
<span id="BytesToString"></span>
> 以零拷贝的方式将字节切片转换为字符串
***
#### func Convert(src A)  B
<span id="Convert"></span>
> 以零拷贝的方式将一个对象转换为另一个对象
>   - 两个对象字段必须完全一致
>   - 该函数可以绕过私有字段的访问限制
***
#### func Verify(handle func ( V))  *VerifyHandle[V]
<span id="Verify"></span>
> 对特定表达式进行校验，当表达式不成立时，将执行 handle
***
#### func OldVersion(version1 string, version2 string)  bool
<span id="OldVersion"></span>
> 检查 version2 对于 version1 来说是不是旧版本
***
#### func CompareVersion(version1 string, version2 string)  int
<span id="CompareVersion"></span>
> 返回一个整数，用于表示两个版本号的比较结果：
>   - 如果 version1 大于 version2，它将返回 1
>   - 如果 version1 小于 version2，它将返回 -1
>   - 如果 version1 和 version2 相等，它将返回 0
***
### BitSet
是一个可以动态增长的比特位集合
  - 默认情况下将使用 64 位无符号整数来表示比特位，当需要表示的比特位超过 64 位时，将自动增长
```go
type BitSet[Bit generic.Integer] struct {
	set []uint64
}
```
#### func (*BitSet) Set(bit Bit)  *BitSet[Bit]
> 将指定的位 bit 设置为 1
***
#### func (*BitSet) Del(bit Bit)  *BitSet[Bit]
> 将指定的位 bit 设置为 0
***
#### func (*BitSet) Shrink()  *BitSet[Bit]
> 将 BitSet 中的比特位集合缩小到最小
>   - 正常情况下当 BitSet 中的比特位超出 64 位时，将自动增长，当 BitSet 中的比特位数量减少时，可以使用该方法将 BitSet 中的比特位集合缩小到最小
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
#### func (*BitSet) MarshalJSON()  []byte,  error
> 实现 json.Marshaler 接口
***
#### func (*BitSet) UnmarshalJSON(data []byte)  error
> 实现 json.Unmarshaler 接口
***
### LossCounter

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
### Matcher
匹配器
```go
type Matcher[Value any, Result any] struct {
	value Value
	r     Result
	d     bool
}
```
### Permission

```go
type Permission[Code generic.Integer, EntityID comparable] struct {
	permissions map[EntityID]Code
	l           sync.RWMutex
}
```
### StackGo
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
### VerifyHandle
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
