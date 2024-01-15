# Aoi

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

aoi 提供了一种有效的方法来处理 AOI（Area of Interest）问题。

AOI 问题是在大规模多人在线游戏中常见的问题，它涉及到确定哪些对象对玩家来说是“感兴趣的”，
也就是说，哪些对象在玩家的视野范围内。

这个包提供了一种数据结构和一些方法来有效地解决这个问题。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewTwoDimensional](#NewTwoDimensional)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[TwoDimensional](#twodimensional)|暂无描述...
|`INTERFACE`|[TwoDimensionalEntity](#twodimensionalentity)|基于2D定义的AOI对象功能接口
|`STRUCT`|[EntityJoinVisionEventHandle](#entityjoinvisioneventhandle)|暂无描述...

</details>


***
## 详情信息
#### func NewTwoDimensional(width int, height int, areaWidth int, areaHeight int) *TwoDimensional[EID, PosType, E]
<span id="NewTwoDimensional"></span>

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewTwoDimensional(t *testing.T) {
	aoiTW := aoi.NewTwoDimensional[int64, float64, *Ent](10000, 10000, 100, 100)
	start := time.Now()
	for i := 0; i < 50000; i++ {
		aoiTW.AddEntity(&Ent{guid: int64(i), pos: geometry.NewPoint[float64](float64(random.Int64(0, 10000)), float64(random.Int64(0, 10000))), vision: 200})
	}
	fmt.Println("添加耗时：", time.Since(start))
	start = time.Now()
	aoiTW.SetSize(10100, 10100)
	fmt.Println("重设大小耗时：", time.Since(start))
}

```


</details>


***
### TwoDimensional `STRUCT`

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
### TwoDimensionalEntity `INTERFACE`
基于2D定义的AOI对象功能接口
  - AOI 对象提供了 AOI 系统中常用的属性，诸如位置坐标和视野范围等
```go
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] interface {
	GetTwoDimensionalEntityID() EID
	GetVision() float64
	GetPosition() geometry.Point[PosType]
}
```
### EntityJoinVisionEventHandle `STRUCT`

```go
type EntityJoinVisionEventHandle[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] func(entity E)
```
