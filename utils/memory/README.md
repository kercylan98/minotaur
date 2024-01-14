# Memory



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/memory)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[Run](#Run)|运行持久化缓存程序
|[BindPersistCacheProgram](#BindPersistCacheProgram)|绑定持久化缓存程序
|[BindAction](#BindAction)|绑定需要缓存的操作函数
|[NewOption](#NewOption)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[Option](#option)|暂无描述...

</details>


#### func Run()
<span id="Run"></span>
> 运行持久化缓存程序
***
#### func BindPersistCacheProgram(name string, handler OutputParamHandlerFunc, option ...*Option)  func ()
<span id="BindPersistCacheProgram"></span>
> 绑定持久化缓存程序
>   - name 持久化缓存程序名称
>   - handler 持久化缓存程序处理函数
>   - option 持久化缓存程序选项
> 
> 注意事项：
>   - 持久化程序建议声明为全局变量进行使用
>   - 持久化程序处理函数参数类型必须与绑定的缓存程序输出参数类型一致，并且相同 name 的持久化程序必须在 BindAction 之后进行绑定
>   - 默认情况下只有执行该函数返回的函数才会进行持久化，如果需要持久化策略，可以设置 option 参数或者自行实现策略调用返回的函数
>   - 所有持久化程序绑定完成后，应该主动调用 Run 函数运行
***
#### func BindAction(name string, handler Func)  Func
<span id="BindAction"></span>
> 绑定需要缓存的操作函数
>   - name 缓存操作名称
>   - handler 缓存操作处理函数
> 
> 注意事项：
>   - 关于持久化缓存程序的绑定请参考 BindPersistCacheProgram
>   - handler 函数的返回值将被作为缓存目标，如果返回值为非指针类型，将可能会发生意外的情况
>   - 当传入的 handler 没有任何返回值时，将不会被缓存，并且不会占用缓存操作名称
> 
> 使用场景：
>   - 例如在游戏中，需要根据玩家 ID 查询玩家信息，可以使用该函数进行绑定，当查询玩家信息时，如果缓存中存在该玩家信息，将直接返回缓存中的数据，否则将执行 handler 函数进行查询并缓存
***
#### func NewOption()  *Option
<span id="NewOption"></span>
***
### Option

```go
type Option struct {
	ticker     *timer.Ticker
	firstDelay time.Duration
	interval   time.Duration
	delay      time.Duration
}
```
#### func (*Option) WithPeriodicity(ticker *timer.Ticker, firstDelay time.Duration, interval time.Duration, delay time.Duration)  *Option
> 设置持久化周期
>   - ticker 定时器，通常建议使用服务器的定时器，这样可以降低多线程的程序复杂性
>   - firstDelay 首次持久化延迟，当首次持久化为 0 时，将会在下一个持久化周期开始时持久化
>   - interval 持久化间隔
>   - delay 每条数据持久化间隔，适当的设置该值可以使持久化期间尽量降低对用户体验的影响，如果为0，将会一次性持久化所有数据
***
