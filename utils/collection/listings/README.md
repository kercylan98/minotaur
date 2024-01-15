# Listings

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
|[NewMatrix](#NewMatrix)|创建一个新的 Matrix 实例。
|[NewPagedSlice](#NewPagedSlice)|创建一个新的 PagedSlice 实例。
|[NewPrioritySlice](#NewPrioritySlice)|创建一个优先级切片
|[NewSyncSlice](#NewSyncSlice)|创建一个 SyncSlice


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Matrix](#struct_Matrix)|暂无描述...
|`STRUCT`|[PagedSlice](#struct_PagedSlice)|是一个高效的动态数组，它通过分页管理内存并减少频繁的内存分配来提高性能。
|`STRUCT`|[PrioritySlice](#struct_PrioritySlice)|是一个优先级切片
|`STRUCT`|[SyncSlice](#struct_SyncSlice)|是基于 sync.RWMutex 实现的线程安全的 slice

</details>


***
## 详情信息
#### func NewMatrix\[V any\](dimensions ...int) *Matrix[V]
<span id="NewMatrix"></span>
> 创建一个新的 Matrix 实例。

***
#### func NewPagedSlice\[T any\](pageSize int) *PagedSlice[T]
<span id="NewPagedSlice"></span>
> 创建一个新的 PagedSlice 实例。

***
#### func NewPrioritySlice\[V any\](lengthAndCap ...int) *PrioritySlice[V]
<span id="NewPrioritySlice"></span>
> 创建一个优先级切片

***
#### func NewSyncSlice\[V any\](length int, cap int) *SyncSlice[V]
<span id="NewSyncSlice"></span>
> 创建一个 SyncSlice

***
<span id="struct_Matrix"></span>
### Matrix `STRUCT`

```go
type Matrix[V any] struct {
	dimensions []int
	data       []V
}
```
<span id="struct_Matrix_Get"></span>

#### func (*Matrix) Get(index ...int)  *V
> 获取矩阵中给定索引的元素。

***
<span id="struct_Matrix_Set"></span>

#### func (*Matrix) Set(index []int, value V)
> 设置矩阵中给定索引的元素。

***
<span id="struct_Matrix_Dimensions"></span>

#### func (*Matrix) Dimensions()  []int
> 返回矩阵的维度大小。

***
<span id="struct_Matrix_Clear"></span>

#### func (*Matrix) Clear()
> 清空矩阵。

***
<span id="struct_PagedSlice"></span>
### PagedSlice `STRUCT`
是一个高效的动态数组，它通过分页管理内存并减少频繁的内存分配来提高性能。
```go
type PagedSlice[T any] struct {
	pages    [][]T
	pageSize int
	len      int
	lenLast  int
}
```
<span id="struct_PagedSlice_Add"></span>

#### func (*PagedSlice) Add(value T)
> 添加一个元素到 PagedSlice 中。

***
<span id="struct_PagedSlice_Get"></span>

#### func (*PagedSlice) Get(index int)  *T
> 获取 PagedSlice 中给定索引的元素。

***
<span id="struct_PagedSlice_Set"></span>

#### func (*PagedSlice) Set(index int, value T)
> 设置 PagedSlice 中给定索引的元素。

***
<span id="struct_PagedSlice_Len"></span>

#### func (*PagedSlice) Len()  int
> 返回 PagedSlice 中元素的数量。

***
<span id="struct_PagedSlice_Clear"></span>

#### func (*PagedSlice) Clear()
> 清空 PagedSlice。

***
<span id="struct_PagedSlice_Range"></span>

#### func (*PagedSlice) Range(f func (index int, value T)  bool)
> 迭代 PagedSlice 中的所有元素。

***
<span id="struct_PrioritySlice"></span>
### PrioritySlice `STRUCT`
是一个优先级切片
```go
type PrioritySlice[V any] struct {
	items []*priorityItem[V]
}
```
<span id="struct_PrioritySlice_Len"></span>

#### func (*PrioritySlice) Len()  int
> 返回切片长度

***
<span id="struct_PrioritySlice_Cap"></span>

#### func (*PrioritySlice) Cap()  int
> 返回切片容量

***
<span id="struct_PrioritySlice_Clear"></span>

#### func (*PrioritySlice) Clear()
> 清空切片

***
<span id="struct_PrioritySlice_Append"></span>

#### func (*PrioritySlice) Append(v V, p int)
> 添加元素

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestPrioritySlice_Append(t *testing.T) {
	var s = listings.NewPrioritySlice[string]()
	s.Append("name_1", 2)
	s.Append("name_2", 1)
	fmt.Println(s)
}

```


</details>


***
<span id="struct_PrioritySlice_Appends"></span>

#### func (*PrioritySlice) Appends(priority int, vs ...V)
> 添加元素

***
<span id="struct_PrioritySlice_Get"></span>

#### func (*PrioritySlice) Get(index int) ( V,  int)
> 获取元素

***
<span id="struct_PrioritySlice_GetValue"></span>

#### func (*PrioritySlice) GetValue(index int)  V
> 获取元素值

***
<span id="struct_PrioritySlice_GetPriority"></span>

#### func (*PrioritySlice) GetPriority(index int)  int
> 获取元素优先级

***
<span id="struct_PrioritySlice_Set"></span>

#### func (*PrioritySlice) Set(index int, value V, priority int)
> 设置元素

***
<span id="struct_PrioritySlice_SetValue"></span>

#### func (*PrioritySlice) SetValue(index int, value V)
> 设置元素值

***
<span id="struct_PrioritySlice_SetPriority"></span>

#### func (*PrioritySlice) SetPriority(index int, priority int)
> 设置元素优先级

***
<span id="struct_PrioritySlice_Action"></span>

#### func (*PrioritySlice) Action(action func (items []*priorityItem[V])  []*priorityItem[V])
> 直接操作切片，如果返回值不为 nil，则替换切片

***
<span id="struct_PrioritySlice_Range"></span>

#### func (*PrioritySlice) Range(action func (index int, item *priorityItem[V])  bool)
> 遍历切片，如果返回值为 false，则停止遍历

***
<span id="struct_PrioritySlice_RangeValue"></span>

#### func (*PrioritySlice) RangeValue(action func (index int, value V)  bool)
> 遍历切片值，如果返回值为 false，则停止遍历

***
<span id="struct_PrioritySlice_RangePriority"></span>

#### func (*PrioritySlice) RangePriority(action func (index int, priority int)  bool)
> 遍历切片优先级，如果返回值为 false，则停止遍历

***
<span id="struct_PrioritySlice_Slice"></span>

#### func (*PrioritySlice) Slice()  []V
> SyncSlice 返回切片

***
<span id="struct_PrioritySlice_String"></span>

#### func (*PrioritySlice) String()  string
> 返回切片字符串

***
<span id="struct_SyncSlice"></span>
### SyncSlice `STRUCT`
是基于 sync.RWMutex 实现的线程安全的 slice
```go
type SyncSlice[V any] struct {
	rw   sync.RWMutex
	data []V
}
```
<span id="struct_SyncSlice_Get"></span>

#### func (*SyncSlice) Get(index int)  V

***
<span id="struct_SyncSlice_GetWithRange"></span>

#### func (*SyncSlice) GetWithRange(start int, end int)  []V

***
<span id="struct_SyncSlice_Set"></span>

#### func (*SyncSlice) Set(index int, value V)

***
<span id="struct_SyncSlice_Append"></span>

#### func (*SyncSlice) Append(values ...V)

***
<span id="struct_SyncSlice_Release"></span>

#### func (*SyncSlice) Release()

***
<span id="struct_SyncSlice_Clear"></span>

#### func (*SyncSlice) Clear()

***
<span id="struct_SyncSlice_GetData"></span>

#### func (*SyncSlice) GetData()  []V

***
