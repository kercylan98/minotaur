# Dp

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

dp (DistributionPattern) 提供用于在二维数组中根据不同的特征标记为数组成员建立分布链接的函数和数据结构。该包的目标是实现快速查找与给定位置成员具有相同特征且位置紧邻的其他成员。
主要特性：
  - 分布链接机制：dp 包提供了一种分布链接的机制，可以根据成员的特征将它们链接在一起。这样，可以快速查找与给定成员具有相同特征且位置紧邻的其他成员。
  - 二维数组支持：该包支持在二维数组中建立分布链接。可以将二维数组中的成员视为节点，并根据其特征进行链接。
  - 快速查找功能：使用 dp 包提供的函数，可以快速查找与给定位置成员具有相同特征且位置紧邻的其他成员。这有助于在二维数组中进行相关性分析或查找相邻成员。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewDistributionPattern](#NewDistributionPattern)|构建一个分布图实例


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[DistributionPattern](#struct_DistributionPattern)|分布图
|`STRUCT`|[Link](#struct_Link)|暂无描述...

</details>


***
## 详情信息
#### func NewDistributionPattern\[Item any\](sameKindVerifyHandle func (itemA Item)  bool) *DistributionPattern[Item]
<span id="NewDistributionPattern"></span>
> 构建一个分布图实例

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewDistributionPattern(t *testing.T) {
	dp := NewDistributionPattern[int](func(itemA, itemB int) bool {
		return itemA == itemB
	})
	matrix := []int{1, 1, 2, 2, 2, 2, 1, 2, 2}
	dp.LoadMatrixWithPos(3, matrix)
	for pos, link := range dp.links {
		fmt.Println(pos, link, fmt.Sprintf("%p", link))
	}
	fmt.Println()
	matrix[6] = 2
	dp.Refresh(6)
	for pos, link := range dp.links {
		fmt.Println(pos, link, fmt.Sprintf("%p", link))
	}
}

```


</details>


***
<span id="struct_DistributionPattern"></span>
### DistributionPattern `STRUCT`
分布图
```go
type DistributionPattern[Item any] struct {
	matrix               []Item
	links                map[int]map[int]Item
	sameKindVerifyHandle func(itemA Item) bool
	width                int
	usePos               bool
}
```
#### func (*DistributionPattern) GetLinks(pos int) (result []Link[Item])
> 获取关联的成员
>   - 其中包含传入的 pos 成员
***
#### func (*DistributionPattern) HasLink(pos int)  bool
> 检查一个位置是否包含除它本身外的其他关联成员
***
#### func (*DistributionPattern) LoadMatrix(matrix [][]Item)
> 通过二维矩阵加载分布图
>   - 通过该函数加载的分布图使用的矩阵是复制后的矩阵，因此无法直接通过刷新(Refresh)来更新分布关系
>   - 需要通过直接刷新的方式请使用 LoadMatrixWithPos
***
#### func (*DistributionPattern) LoadMatrixWithPos(width int, matrix []Item)
> 通过二维矩阵加载分布图
***
#### func (*DistributionPattern) Refresh(pos int)
> 刷新特定位置的分布关系
>   - 由于 LoadMatrix 的矩阵是复制后的矩阵，所以任何外部的改动都不会影响到分布图的变化，在这种情况下，刷新将没有任何意义
>   - 需要通过直接刷新的方式请使用 LoadMatrixWithPos 加载矩阵，或者通过 RefreshWithItem 函数进行刷新
***
#### func (*DistributionPattern) RefreshWithItem(pos int, item Item)
> 通过特定的成员刷新特定位置的分布关系
>   - 如果矩阵通过 LoadMatrixWithPos 加载，将会重定向至 Refresh
***
<span id="struct_Link"></span>
### Link `STRUCT`

```go
type Link[V any] struct {
	Pos  int
	Item V
}
```
