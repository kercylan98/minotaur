# Mappings

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
|[NewSyncMap](#NewSyncMap)|创建一个 SyncMap


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[SyncMap](#syncmap)|是基于 sync.RWMutex 实现的线程安全的 map

</details>


***
## 详情信息
#### func NewSyncMap(source ...map[K]V) *SyncMap[K, V]
<span id="NewSyncMap"></span>
> 创建一个 SyncMap

***
### SyncMap `STRUCT`
是基于 sync.RWMutex 实现的线程安全的 map
  - 适用于要考虑并发读写但是并发读写的频率不高的情况
```go
type SyncMap[K comparable, V any] struct {
	lock sync.RWMutex
	data map[K]V
	atom bool
}
```
