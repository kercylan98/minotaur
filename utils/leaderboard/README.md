# Leaderboard

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
|[NewBinarySearch](#NewBinarySearch)|创建一个基于内存的二分查找排行榜
|[WithBinarySearchCount](#WithBinarySearchCount)|通过限制排行榜竞争者数量来创建排行榜
|[WithBinarySearchASC](#WithBinarySearchASC)|通过升序的方式创建排行榜


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[BinarySearch](#struct_BinarySearch)|暂无描述...
|`STRUCT`|[BinarySearchRankChangeEventHandle](#struct_BinarySearchRankChangeEventHandle)|暂无描述...
|`STRUCT`|[BinarySearchOption](#struct_BinarySearchOption)|暂无描述...

</details>


***
## 详情信息
#### func NewBinarySearch\[CompetitorID comparable, Score generic.Ordered\](options ...BinarySearchOption[CompetitorID, Score]) *BinarySearch[CompetitorID, Score]
<span id="NewBinarySearch"></span>
> 创建一个基于内存的二分查找排行榜

示例代码：
```go

func ExampleNewBinarySearch() {
	bs := leaderboard2.NewBinarySearch[string, int](leaderboard2.WithBinarySearchCount[string, int](10))
	fmt.Println(bs != nil)
}

```

***
#### func WithBinarySearchCount\[CompetitorID comparable, Score generic.Ordered\](rankCount int) BinarySearchOption[CompetitorID, Score]
<span id="WithBinarySearchCount"></span>
> 通过限制排行榜竞争者数量来创建排行榜
>   - 默认情况下允许100位竞争者

***
#### func WithBinarySearchASC\[CompetitorID comparable, Score generic.Ordered\]() BinarySearchOption[CompetitorID, Score]
<span id="WithBinarySearchASC"></span>
> 通过升序的方式创建排行榜
>   - 默认情况下为降序

***
<span id="struct_BinarySearch"></span>
### BinarySearch `STRUCT`

```go
type BinarySearch[CompetitorID comparable, Score generic.Ordered] struct {
	*binarySearchEvent[CompetitorID, Score]
	asc                         bool
	rankCount                   int
	competitors                 *mappings.SyncMap[CompetitorID, Score]
	scores                      []*scoreItem[CompetitorID, Score]
	rankChangeEventHandles      []BinarySearchRankChangeEventHandle[CompetitorID, Score]
	rankClearBeforeEventHandles []BinarySearchRankClearBeforeEventHandle[CompetitorID, Score]
}
```
<span id="struct_BinarySearchRankChangeEventHandle"></span>
### BinarySearchRankChangeEventHandle `STRUCT`

```go
type BinarySearchRankChangeEventHandle[CompetitorID comparable, Score generic.Ordered] func(leaderboard *BinarySearch[CompetitorID, Score], competitorId CompetitorID, oldRank int, oldScore Score)
```
<span id="struct_BinarySearchOption"></span>
### BinarySearchOption `STRUCT`

```go
type BinarySearchOption[CompetitorID comparable, Score generic.Ordered] func(list *BinarySearch[CompetitorID, Score])
```
