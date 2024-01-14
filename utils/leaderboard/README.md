# Leaderboard



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/leaderboard)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewBinarySearch](#NewBinarySearch)|创建一个基于内存的二分查找排行榜
|[WithBinarySearchCount](#WithBinarySearchCount)|通过限制排行榜竞争者数量来创建排行榜
|[WithBinarySearchASC](#WithBinarySearchASC)|通过升序的方式创建排行榜


> 结构体定义

|结构体|描述
|:--|:--
|[BinarySearch](#binarysearch)|暂无描述...
|[BinarySearchRankChangeEventHandle](#binarysearchrankchangeeventhandle)|暂无描述...
|[BinarySearchOption](#binarysearchoption)|暂无描述...

</details>


#### func NewBinarySearch(options ...BinarySearchOption[CompetitorID, Score])  *BinarySearch[CompetitorID, Score]
<span id="NewBinarySearch"></span>
> 创建一个基于内存的二分查找排行榜
***
#### func WithBinarySearchCount(rankCount int)  BinarySearchOption[CompetitorID, Score]
<span id="WithBinarySearchCount"></span>
> 通过限制排行榜竞争者数量来创建排行榜
>   - 默认情况下允许100位竞争者
***
#### func WithBinarySearchASC()  BinarySearchOption[CompetitorID, Score]
<span id="WithBinarySearchASC"></span>
> 通过升序的方式创建排行榜
>   - 默认情况下为降序
***
### BinarySearch

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
### BinarySearchRankChangeEventHandle

```go
type BinarySearchRankChangeEventHandle[CompetitorID comparable, Score generic.Ordered] struct{}
```
### BinarySearchOption

```go
type BinarySearchOption[CompetitorID comparable, Score generic.Ordered] struct{}
```
