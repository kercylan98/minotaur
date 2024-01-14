# Aoi

aoi 提供了一种有效的方法来处理 AOI（Area of Interest）问题。

AOI 问题是在大规模多人在线游戏中常见的问题，它涉及到确定哪些对象对玩家来说是“感兴趣的”，
也就是说，哪些对象在玩家的视野范围内。

这个包提供了一种数据结构和一些方法来有效地解决这个问题。

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/aoi)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewTwoDimensional](#NewTwoDimensional)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[TwoDimensional](#twodimensional)|暂无描述...
|[TwoDimensionalEntity](#twodimensionalentity)|基于2D定义的AOI对象功能接口
|[EntityJoinVisionEventHandle](#entityjoinvisioneventhandle)|暂无描述...

</details>


#### func NewTwoDimensional(width int, height int, areaWidth int, areaHeight int)  *TwoDimensional[EID, PosType, E]
<span id="NewTwoDimensional"></span>
***
### TwoDimensional

```go
type TwoDimensional[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] struct {
	*event[EID, PosType, E]
	rw               sync.RWMutex
	width            float64
	height           float64
	areaWidth        float64
	areaHeight       float64
	areaWidthLimit   int
	areaHeightLimit  int
	areas            [][]map[EID]E
	focus            map[EID]map[EID]E
	repartitionQueue []func()
}
```
### TwoDimensionalEntity
基于2D定义的AOI对象功能接口
  - AOI 对象提供了 AOI 系统中常用的属性，诸如位置坐标和视野范围等
```go
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] struct{}
```
### EntityJoinVisionEventHandle

```go
type EntityJoinVisionEventHandle[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] struct{}
```
