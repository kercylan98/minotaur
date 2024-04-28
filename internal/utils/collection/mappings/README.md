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
|`STRUCT`|[SyncMap](#struct_SyncMap)|是基于 sync.RWMutex 实现的线程安全的 map

</details>


***
## 详情信息
#### func NewSyncMap\[K comparable, V any\](source ...map[K]V) *SyncMap[K, V]
<span id="NewSyncMap"></span>
> 创建一个 SyncMap

***
<span id="struct_SyncMap"></span>
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
<span id="struct_SyncMap_Set"></span>

#### func (*SyncMap) Set(key K, value V)
> 设置一个值

***
<span id="struct_SyncMap_Get"></span>

#### func (*SyncMap) Get(key K)  V
> 获取一个值

***
<span id="struct_SyncMap_Atom"></span>

#### func (*SyncMap) Atom(handle func (m map[K]V))
> 原子操作

***
<span id="struct_SyncMap_Exist"></span>

#### func (*SyncMap) Exist(key K)  bool
> 判断是否存在

***
<span id="struct_SyncMap_GetExist"></span>

#### func (*SyncMap) GetExist(key K) ( V,  bool)
> 获取一个值并判断是否存在

***
<span id="struct_SyncMap_Delete"></span>

#### func (*SyncMap) Delete(key K)
> 删除一个值

***
<span id="struct_SyncMap_DeleteGet"></span>

#### func (*SyncMap) DeleteGet(key K)  V
> 删除一个值并返回

***
<span id="struct_SyncMap_DeleteGetExist"></span>

#### func (*SyncMap) DeleteGetExist(key K) ( V,  bool)
> 删除一个值并返回是否存在

***
<span id="struct_SyncMap_DeleteExist"></span>

#### func (*SyncMap) DeleteExist(key K)  bool
> 删除一个值并返回是否存在

***
<span id="struct_SyncMap_Clear"></span>

#### func (*SyncMap) Clear()
> 清空

***
<span id="struct_SyncMap_ClearHandle"></span>

#### func (*SyncMap) ClearHandle(handle func (key K, value V))
> 清空并处理

***
<span id="struct_SyncMap_Range"></span>

#### func (*SyncMap) Range(handle func (key K, value V)  bool)
> 遍历所有值，如果 handle 返回 true 则停止遍历

***
<span id="struct_SyncMap_Keys"></span>

#### func (*SyncMap) Keys()  []K
> 获取所有的键

***
<span id="struct_SyncMap_Slice"></span>

#### func (*SyncMap) Slice()  []V
> 获取所有的值

***
<span id="struct_SyncMap_Map"></span>

#### func (*SyncMap) Map()  map[K]V
> 转换为普通 map

***
<span id="struct_SyncMap_Size"></span>

#### func (*SyncMap) Size()  int
> 获取数量

***
<span id="struct_SyncMap_MarshalJSON"></span>

#### func (*SyncMap) MarshalJSON() ( []byte,  error)

***
<span id="struct_SyncMap_UnmarshalJSON"></span>

#### func (*SyncMap) UnmarshalJSON(bytes []byte)  error

***
