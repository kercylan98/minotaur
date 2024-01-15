# Sole

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
|[RegNameSpace](#RegNameSpace)|注册特定命名空间的唯一标识符
|[UnRegNameSpace](#UnRegNameSpace)|解除注销特定命名空间的唯一标识符
|[Get](#Get)|获取全局唯一标识符
|[Reset](#Reset)|重置全局唯一标识符
|[GetWith](#GetWith)|获取特定命名空间的唯一标识符
|[ResetWith](#ResetWith)|重置特定命名空间的唯一标识符
|[NewOnce](#NewOnce)|创建一个用于数据取值去重的结构实例
|[SonyflakeIDE](#SonyflakeIDE)|获取一个雪花id
|[SonyflakeID](#SonyflakeID)|获取一个雪花id
|[SonyflakeSetting](#SonyflakeSetting)|配置雪花id生成策略
|[AutoIncrementUint32](#AutoIncrementUint32)|获取一个自增的 uint32 值
|[AutoIncrementUint64](#AutoIncrementUint64)|获取一个自增的 uint64 值
|[AutoIncrementInt32](#AutoIncrementInt32)|获取一个自增的 int32 值
|[AutoIncrementInt64](#AutoIncrementInt64)|获取一个自增的 int64 值
|[AutoIncrementInt](#AutoIncrementInt)|获取一个自增的 int 值
|[AutoIncrementString](#AutoIncrementString)|获取一个自增的字符串


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Once](#struct_Once)|用于数据取值去重的结构体

</details>


***
## 详情信息
#### func RegNameSpace(name any)
<span id="RegNameSpace"></span>
> 注册特定命名空间的唯一标识符

***
#### func UnRegNameSpace(name any)
<span id="UnRegNameSpace"></span>
> 解除注销特定命名空间的唯一标识符

***
#### func Get() int64
<span id="Get"></span>
> 获取全局唯一标识符

***
#### func Reset()
<span id="Reset"></span>
> 重置全局唯一标识符

***
#### func GetWith(name any) int64
<span id="GetWith"></span>
> 获取特定命名空间的唯一标识符

***
#### func ResetWith(name any)
<span id="ResetWith"></span>
> 重置特定命名空间的唯一标识符

***
#### func NewOnce\[V any\]() *Once[V]
<span id="NewOnce"></span>
> 创建一个用于数据取值去重的结构实例

***
#### func SonyflakeIDE() (int64,  error)
<span id="SonyflakeIDE"></span>
> 获取一个雪花id

***
#### func SonyflakeID() int64
<span id="SonyflakeID"></span>
> 获取一个雪花id

***
#### func SonyflakeSetting(settings sonyflake.Settings)
<span id="SonyflakeSetting"></span>
> 配置雪花id生成策略

***
#### func AutoIncrementUint32() uint32
<span id="AutoIncrementUint32"></span>
> 获取一个自增的 uint32 值

***
#### func AutoIncrementUint64() uint64
<span id="AutoIncrementUint64"></span>
> 获取一个自增的 uint64 值

***
#### func AutoIncrementInt32() int32
<span id="AutoIncrementInt32"></span>
> 获取一个自增的 int32 值

***
#### func AutoIncrementInt64() int64
<span id="AutoIncrementInt64"></span>
> 获取一个自增的 int64 值

***
#### func AutoIncrementInt() int
<span id="AutoIncrementInt"></span>
> 获取一个自增的 int 值

***
#### func AutoIncrementString() string
<span id="AutoIncrementString"></span>
> 获取一个自增的字符串

***
<span id="struct_Once"></span>
### Once `STRUCT`
用于数据取值去重的结构体
```go
type Once[V any] struct {
	r map[any]struct{}
}
```
<span id="struct_Once_Get"></span>

#### func (*Once) Get(key any, value V, defaultValue V)  V
> 获取一个值，当该值已经被获取过的时候，返回 defaultValue，否则返回 value

***
<span id="struct_Once_Reset"></span>

#### func (*Once) Reset(key ...any)
> 当 key 数量大于 0 时，将会重置对应 key 的记录，否则重置所有记录

***
