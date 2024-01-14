# Configuration

configuration 基于配置导表功能实现的配置加载及刷新功能

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/configuration)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[Init](#Init)|配置初始化
|[Load](#Load)|加载配置
|[Refresh](#Refresh)|刷新配置
|[WithTickerLoad](#WithTickerLoad)|通过定时器加载配置
|[StopTickerLoad](#StopTickerLoad)|停止通过定时器加载配置
|[RegConfigRefreshEvent](#RegConfigRefreshEvent)|当配置刷新时将立即执行被注册的事件处理函数
|[OnConfigRefreshEvent](#OnConfigRefreshEvent)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[RefreshEventHandle](#refresheventhandle)|配置刷新事件处理函数
|[Loader](#loader)|配置加载器

</details>


#### func Init(loader ...Loader)
<span id="Init"></span>
> 配置初始化
>   - 在初始化后会立即加载配置
***
#### func Load()
<span id="Load"></span>
> 加载配置
>   - 加载后并不会刷新线上配置，需要执行 Refresh 函数对线上配置进行刷新
***
#### func Refresh()
<span id="Refresh"></span>
> 刷新配置
***
#### func WithTickerLoad(ticker *timer.Ticker, interval time.Duration)
<span id="WithTickerLoad"></span>
> 通过定时器加载配置
>   - 通过定时器加载配置后，会自动刷新线上配置
>   - 调用该函数后不会立即刷新，而是在 interval 后加载并刷新一次配置，之后每隔 interval 加载并刷新一次配置
***
#### func StopTickerLoad()
<span id="StopTickerLoad"></span>
> 停止通过定时器加载配置
***
#### func RegConfigRefreshEvent(handle RefreshEventHandle)
<span id="RegConfigRefreshEvent"></span>
> 当配置刷新时将立即执行被注册的事件处理函数
***
#### func OnConfigRefreshEvent()
<span id="OnConfigRefreshEvent"></span>
***
### RefreshEventHandle
配置刷新事件处理函数
```go
type RefreshEventHandle struct{}
```
### Loader
配置加载器
```go
type Loader struct{}
```
