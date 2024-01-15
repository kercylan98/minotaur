# Offset

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
|[NewTime](#NewTime)|新建一个包含偏移的时间
|[SetGlobal](#SetGlobal)|设置全局偏移时间
|[GetGlobal](#GetGlobal)|获取全局偏移时间
|[Now](#Now)|获取当前时间偏移后的时间
|[Since](#Since)|获取当前时间偏移后的时间自从 t 以来经过的时间


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Time](#time)|带有偏移量的时间

</details>


***
## 详情信息
#### func NewTime(offset time.Duration) *Time
<span id="NewTime"></span>
> 新建一个包含偏移的时间

***
#### func SetGlobal(offset time.Duration)
<span id="SetGlobal"></span>
> 设置全局偏移时间

***
#### func GetGlobal() *Time
<span id="GetGlobal"></span>
> 获取全局偏移时间

***
#### func Now() time.Time
<span id="Now"></span>
> 获取当前时间偏移后的时间

***
#### func Since(t time.Time) time.Duration
<span id="Since"></span>
> 获取当前时间偏移后的时间自从 t 以来经过的时间

***
### Time `STRUCT`
带有偏移量的时间
```go
type Time struct {
	offset time.Duration
}
```
#### func (*Time) SetOffset(offset time.Duration)
> 设置时间偏移
***
#### func (*Time) Now()  time.Time
> 获取当前时间偏移后的时间
***
#### func (*Time) Since(t time.Time)  time.Duration
> 获取当前时间偏移后的时间自从 t 以来经过的时间
***
