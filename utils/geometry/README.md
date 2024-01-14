# Geometry

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

geometry 旨在提供一组用于处理几何形状和计算几何属性的函数和数据结构。该包旨在简化几何计算的过程，并提供一致的接口和易于使用的功能。
主要特性：
  - 几何形状："geometry"包支持处理各种几何形状，如点、线、多边形和圆等。您可以使用这些形状来表示和操作实际世界中的几何对象。
  - 几何计算：该包提供了一系列函数，用于执行常见的几何计算，如计算两点之间的距离、计算线段的长度、计算多边形的面积等。这些函数旨在提供高效和准确的计算结果。
  - 坐标转换："geometry"包还提供了一些函数，用于在不同坐标系之间进行转换。您可以将点从笛卡尔坐标系转换为极坐标系，或者从二维坐标系转换为三维坐标系等。
  - 简化接口：该包的设计目标之一是提供简化的接口，使几何计算变得更加直观和易于使用。您可以轻松地创建和操作几何对象，而无需处理繁琐的底层细节。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewCircle](#NewCircle)|通过传入圆的半径和需要的点数量，生成一个圆
|[CalcCircleCentroidDistance](#CalcCircleCentroidDistance)|计算两个圆质心距离
|[GetOppositionDirection](#GetOppositionDirection)|获取特定方向的对立方向
|[GetDirectionNextWithCoordinate](#GetDirectionNextWithCoordinate)|获取特定方向上的下一个坐标
|[GetDirectionNextWithPoint](#GetDirectionNextWithPoint)|获取特定方向上的下一个坐标
|[GetDirectionNextWithPos](#GetDirectionNextWithPos)|获取位置在特定宽度和特定方向上的下一个位置
|[CalcDirection](#CalcDirection)|计算点2位于点1的方向
|[CalcDistanceWithCoordinate](#CalcDistanceWithCoordinate)|计算两点之间的距离
|[CalcDistanceWithPoint](#CalcDistanceWithPoint)|计算两点之间的距离
|[CalcDistanceSquared](#CalcDistanceSquared)|计算两点之间的平方距离
|[CalcAngle](#CalcAngle)|计算点2位于点1之间的角度
|[CalcNewCoordinate](#CalcNewCoordinate)|根据给定的x、y坐标、角度和距离计算新的坐标
|[CalcRadianWithAngle](#CalcRadianWithAngle)|根据角度 angle 计算弧度
|[CalcAngleDifference](#CalcAngleDifference)|计算两个角度之间的最小角度差
|[CalcRayIsIntersect](#CalcRayIsIntersect)|根据给定的位置和角度生成射线，检测射线是否与多边形发生碰撞
|[NewLineSegment](#NewLineSegment)|创建一根线段
|[NewLineSegmentCap](#NewLineSegmentCap)|创建一根包含数据的线段
|[NewLineSegmentCapWithLine](#NewLineSegmentCapWithLine)|通过已有线段创建一根包含数据的线段
|[ConvertLineSegmentGeneric](#ConvertLineSegmentGeneric)|转换线段的泛型类型为特定类型
|[PointOnLineSegmentWithCoordinate](#PointOnLineSegmentWithCoordinate)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[PointOnLineSegmentWithPos](#PointOnLineSegmentWithPos)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[PointOnLineSegmentWithPoint](#PointOnLineSegmentWithPoint)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[PointOnLineSegmentWithCoordinateInBounds](#PointOnLineSegmentWithCoordinateInBounds)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[PointOnLineSegmentWithPosInBounds](#PointOnLineSegmentWithPosInBounds)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[PointOnLineSegmentWithPointInBounds](#PointOnLineSegmentWithPointInBounds)|通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
|[CalcLineSegmentIsCollinear](#CalcLineSegmentIsCollinear)|检查两条线段在一个误差内是否共线
|[CalcLineSegmentIsOverlap](#CalcLineSegmentIsOverlap)|通过对点进行排序来检查两条共线线段是否重叠，返回重叠线段
|[CalcLineSegmentIsIntersect](#CalcLineSegmentIsIntersect)|计算两条线段是否相交
|[CalcLineSegmentSlope](#CalcLineSegmentSlope)|计算线段的斜率
|[CalcLineSegmentIntercept](#CalcLineSegmentIntercept)|计算线段的截距
|[NewPoint](#NewPoint)|创建一个由 x、y 坐标组成的点
|[NewPointCap](#NewPointCap)|创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
|[NewPointCapWithData](#NewPointCapWithData)|通过设置数据的方式创建一个由 x、y 坐标组成的点，这个点具有一个数据容量
|[NewPointCapWithPoint](#NewPointCapWithPoint)|通过设置数据的方式创建一个由已有坐标组成的点，这个点具有一个数据容量
|[CoordinateToPoint](#CoordinateToPoint)|将坐标转换为x、y的坐标数组
|[CoordinateToPos](#CoordinateToPos)|将坐标转换为二维数组的顺序位置坐标
|[PointToCoordinate](#PointToCoordinate)|将坐标数组转换为x和y坐标
|[PointToPos](#PointToPos)|将坐标转换为二维数组的顺序位置
|[PosToCoordinate](#PosToCoordinate)|通过宽度将一个二维数组的顺序位置转换为xy坐标
|[PosToPoint](#PosToPoint)|通过宽度将一个二维数组的顺序位置转换为x、y的坐标数组
|[PosToCoordinateX](#PosToCoordinateX)|通过宽度将一个二维数组的顺序位置转换为X坐标
|[PosToCoordinateY](#PosToCoordinateY)|通过宽度将一个二维数组的顺序位置转换为Y坐标
|[PointCopy](#PointCopy)|复制一个坐标数组
|[PointToPosWithMulti](#PointToPosWithMulti)|将一组坐标转换为二维数组的顺序位置
|[PosToPointWithMulti](#PosToPointWithMulti)|将一组二维数组的顺序位置转换为一组数组坐标
|[PosSameRow](#PosSameRow)|返回两个顺序位置在同一宽度是否位于同一行
|[DoublePointToCoordinate](#DoublePointToCoordinate)|将两个位置转换为 x1, y1, x2, y2 的坐标进行返回
|[CalcProjectionPoint](#CalcProjectionPoint)|计算一个点到一条线段的最近点（即投影点）的。这个函数接收一个点和一条线段作为输入，线段由两个端点组成。
|[GetAdjacentTranslatePos](#GetAdjacentTranslatePos)|获取一个连续位置的矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
|[GetAdjacentTranslateCoordinateXY](#GetAdjacentTranslateCoordinateXY)|获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
|[GetAdjacentTranslateCoordinateYX](#GetAdjacentTranslateCoordinateYX)|获取一个基于 y、x 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置
|[GetAdjacentDiagonalsPos](#GetAdjacentDiagonalsPos)|获取一个连续位置的矩阵中，特定位置相邻的对角线最多四个方向的位置
|[GetAdjacentDiagonalsCoordinateXY](#GetAdjacentDiagonalsCoordinateXY)|获取一个基于 x、y 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置
|[GetAdjacentDiagonalsCoordinateYX](#GetAdjacentDiagonalsCoordinateYX)|获取一个基于 tx 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置
|[GetAdjacentPos](#GetAdjacentPos)|获取一个连续位置的矩阵中，特定位置相邻的最多八个方向的位置
|[GetAdjacentCoordinateXY](#GetAdjacentCoordinateXY)|获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多八个方向的位置
|[GetAdjacentCoordinateYX](#GetAdjacentCoordinateYX)|获取一个基于 yx 的二维矩阵中，特定位置相邻的最多八个方向的位置
|[CoordinateMatrixToPosMatrix](#CoordinateMatrixToPosMatrix)|将二维矩阵转换为顺序的二维矩阵
|[GetShapeCoverageAreaWithPoint](#GetShapeCoverageAreaWithPoint)|通过传入的一组坐标 points 计算一个图形覆盖的矩形范围
|[GetShapeCoverageAreaWithPos](#GetShapeCoverageAreaWithPos)|通过传入的一组坐标 positions 计算一个图形覆盖的矩形范围
|[CoverageAreaBoundless](#CoverageAreaBoundless)|将一个图形覆盖矩形范围设置为无边的
|[GenerateShapeOnRectangle](#GenerateShapeOnRectangle)|生成一组二维坐标的形状
|[GenerateShapeOnRectangleWithCoordinate](#GenerateShapeOnRectangleWithCoordinate)|生成一组二维坐标的形状
|[GetExpressibleRectangleBySize](#GetExpressibleRectangleBySize)|获取一个宽高可表达的所有特定尺寸以上的矩形形状
|[GetExpressibleRectangle](#GetExpressibleRectangle)|获取一个宽高可表达的所有矩形形状
|[GetRectangleFullPointsByXY](#GetRectangleFullPointsByXY)|通过开始结束坐标获取一个矩形包含的所有点
|[GetRectangleFullPoints](#GetRectangleFullPoints)|获取一个矩形填充满后包含的所有点
|[GetRectangleFullPos](#GetRectangleFullPos)|获取一个矩形填充满后包含的所有位置
|[CalcRectangleCentroid](#CalcRectangleCentroid)|计算矩形质心
|[SetShapeStringHasBorder](#SetShapeStringHasBorder)|设置 Shape.String 是拥有边界的
|[SetShapeStringNotHasBorder](#SetShapeStringNotHasBorder)|设置 Shape.String 是没有边界的
|[NewShape](#NewShape)|通过多个点生成一个形状进行返回
|[NewShapeWithString](#NewShapeWithString)|通过字符串将指定 rune 转换为点位置生成形状进行返回
|[CalcBoundingRadius](#CalcBoundingRadius)|计算多边形转换为圆的半径
|[CalcBoundingRadiusWithCentroid](#CalcBoundingRadiusWithCentroid)|计算多边形在特定质心下圆的半径
|[CalcTriangleTwiceArea](#CalcTriangleTwiceArea)|计算由 a、b、c 三个点组成的三角形的面积的两倍
|[IsPointOnEdge](#IsPointOnEdge)|检查点是否在 edges 的任意一条边上
|[ProjectionPointToShape](#ProjectionPointToShape)|将一个点投影到一个多边形上，找到离该点最近的投影点，并返回投影点和距离
|[WithShapeSearchRectangleLowerLimit](#WithShapeSearchRectangleLowerLimit)|通过矩形宽高下限的方式搜索
|[WithShapeSearchRectangleUpperLimit](#WithShapeSearchRectangleUpperLimit)|通过矩形宽高上限的方式搜索
|[WithShapeSearchRightAngle](#WithShapeSearchRightAngle)|通过直角的方式进行搜索
|[WithShapeSearchOppositionDirection](#WithShapeSearchOppositionDirection)|通过限制对立方向的方式搜索
|[WithShapeSearchDirectionCount](#WithShapeSearchDirectionCount)|通过限制方向数量的方式搜索
|[WithShapeSearchDirectionCountLowerLimit](#WithShapeSearchDirectionCountLowerLimit)|通过限制特定方向数量下限的方式搜索
|[WithShapeSearchDirectionCountUpperLimit](#WithShapeSearchDirectionCountUpperLimit)|通过限制特定方向数量上限的方式搜索
|[WithShapeSearchDeduplication](#WithShapeSearchDeduplication)|通过去重的方式进行搜索
|[WithShapeSearchPointCountLowerLimit](#WithShapeSearchPointCountLowerLimit)|通过限制图形构成的最小点数进行搜索
|[WithShapeSearchPointCountUpperLimit](#WithShapeSearchPointCountUpperLimit)|通过限制图形构成的最大点数进行搜索
|[WithShapeSearchAsc](#WithShapeSearchAsc)|通过升序的方式进行搜索
|[WithShapeSearchDesc](#WithShapeSearchDesc)|通过降序的方式进行搜索


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Circle](#circle)|圆形
|`STRUCT`|[FloorPlan](#floorplan)|平面图
|`STRUCT`|[Direction](#direction)|方向
|`STRUCT`|[LineSegment](#linesegment)|通过两个点表示一根线段
|`STRUCT`|[LineSegmentCap](#linesegmentcap)|可以包含一份额外数据的线段
|`STRUCT`|[Point](#point)|表示了一个由 x、y 坐标组成的点
|`STRUCT`|[PointCap](#pointcap)|表示了一个由 x、y 坐标组成的点，这个点具有一个数据容量
|`STRUCT`|[Shape](#shape)|通过多个点表示了一个形状
|`STRUCT`|[ShapeSearchOption](#shapesearchoption)|图形搜索可选项，用于 Shape.ShapeSearch 搜索支持

</details>


***
## 详情信息
#### func NewCircle(radius V, points int)  Circle[V]
<span id="NewCircle"></span>
> 通过传入圆的半径和需要的点数量，生成一个圆

示例代码：
```go

func ExampleNewCircle() {
	fmt.Println(geometry.NewCircle[float64](7, 12))
}

```

***
#### func CalcCircleCentroidDistance(circle1 Circle[V], circle2 Circle[V])  V
<span id="CalcCircleCentroidDistance"></span>
> 计算两个圆质心距离

***
#### func GetOppositionDirection(direction Direction)  Direction
<span id="GetOppositionDirection"></span>
> 获取特定方向的对立方向

***
#### func GetDirectionNextWithCoordinate(direction Direction, x V, y V) (nx V, ny V)
<span id="GetDirectionNextWithCoordinate"></span>
> 获取特定方向上的下一个坐标

***
#### func GetDirectionNextWithPoint(direction Direction, point Point[V])  Point[V]
<span id="GetDirectionNextWithPoint"></span>
> 获取特定方向上的下一个坐标

***
#### func GetDirectionNextWithPos(direction Direction, width V, pos V)  V
<span id="GetDirectionNextWithPos"></span>
> 获取位置在特定宽度和特定方向上的下一个位置
>   - 需要注意的是，在左右方向时，当下一个位置不在矩形区域内时，将会返回上一行的末位置或下一行的首位置

***
#### func CalcDirection(x1 V, y1 V, x2 V, y2 V)  Direction
<span id="CalcDirection"></span>
> 计算点2位于点1的方向

***
#### func CalcDistanceWithCoordinate(x1 V, y1 V, x2 V, y2 V)  V
<span id="CalcDistanceWithCoordinate"></span>
> 计算两点之间的距离

***
#### func CalcDistanceWithPoint(point1 Point[V], point2 Point[V])  V
<span id="CalcDistanceWithPoint"></span>
> 计算两点之间的距离

***
#### func CalcDistanceSquared(x1 V, y1 V, x2 V, y2 V)  V
<span id="CalcDistanceSquared"></span>
> 计算两点之间的平方距离
>   - 这个函数的主要用途是在需要计算两点之间距离的情况下，但不需要得到实际的距离值，而只需要比较距离大小。因为平方根运算相对较为耗时，所以在只需要比较大小的情况下，通常会使用平方距离。

***
#### func CalcAngle(x1 V, y1 V, x2 V, y2 V)  V
<span id="CalcAngle"></span>
> 计算点2位于点1之间的角度

***
#### func CalcNewCoordinate(x V, y V, angle V, distance V) (newX V, newY V)
<span id="CalcNewCoordinate"></span>
> 根据给定的x、y坐标、角度和距离计算新的坐标

***
#### func CalcRadianWithAngle(angle V)  V
<span id="CalcRadianWithAngle"></span>
> 根据角度 angle 计算弧度

***
#### func CalcAngleDifference(angleA V, angleB V)  V
<span id="CalcAngleDifference"></span>
> 计算两个角度之间的最小角度差

***
#### func CalcRayIsIntersect(x V, y V, angle V, shape Shape[V])  bool
<span id="CalcRayIsIntersect"></span>
> 根据给定的位置和角度生成射线，检测射线是否与多边形发生碰撞

***
#### func NewLineSegment(start Point[V], end Point[V])  LineSegment[V]
<span id="NewLineSegment"></span>
> 创建一根线段

***
#### func NewLineSegmentCap(start Point[V], end Point[V], data Data)  LineSegmentCap[V, Data]
<span id="NewLineSegmentCap"></span>
> 创建一根包含数据的线段

***
#### func NewLineSegmentCapWithLine(line LineSegment[V], data Data)  LineSegmentCap[V, Data]
<span id="NewLineSegmentCapWithLine"></span>
> 通过已有线段创建一根包含数据的线段

***
#### func ConvertLineSegmentGeneric(line LineSegment[V])  LineSegment[TO]
<span id="ConvertLineSegmentGeneric"></span>
> 转换线段的泛型类型为特定类型

***
#### func PointOnLineSegmentWithCoordinate(x1 V, y1 V, x2 V, y2 V, x V, y V)  bool
<span id="PointOnLineSegmentWithCoordinate"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上

***
#### func PointOnLineSegmentWithPos(width V, pos1 V, pos2 V, pos V)  bool
<span id="PointOnLineSegmentWithPos"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上

***
#### func PointOnLineSegmentWithPoint(point1 Point[V], point2 Point[V], point Point[V])  bool
<span id="PointOnLineSegmentWithPoint"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上

***
#### func PointOnLineSegmentWithCoordinateInBounds(x1 V, y1 V, x2 V, y2 V, x V, y V)  bool
<span id="PointOnLineSegmentWithCoordinateInBounds"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
>   - 与 PointOnLineSegmentWithCoordinate 不同的是， PointOnLineSegmentWithCoordinateInBounds 中会判断线段及点的位置是否正确

***
#### func PointOnLineSegmentWithPosInBounds(width V, pos1 V, pos2 V, pos V)  bool
<span id="PointOnLineSegmentWithPosInBounds"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
>   - 与 PointOnLineSegmentWithPos 不同的是， PointOnLineSegmentWithPosInBounds 中会判断线段及点的位置是否正确

***
#### func PointOnLineSegmentWithPointInBounds(point1 Point[V], point2 Point[V], point Point[V])  bool
<span id="PointOnLineSegmentWithPointInBounds"></span>
> 通过一个线段两个点的位置和一个点的坐标，判断这个点是否在一条线段上
>   - 与 PointOnLineSegmentWithPoint 不同的是， PointOnLineSegmentWithPointInBounds 中会判断线段及点的位置是否正确

***
#### func CalcLineSegmentIsCollinear(line1 LineSegment[V], line2 LineSegment[V], tolerance V)  bool
<span id="CalcLineSegmentIsCollinear"></span>
> 检查两条线段在一个误差内是否共线
>   - 共线是指两条线段在同一直线上，即它们的延长线可以重合

***
#### func CalcLineSegmentIsOverlap(line1 LineSegment[V], line2 LineSegment[V]) (line LineSegment[V], overlap bool)
<span id="CalcLineSegmentIsOverlap"></span>
> 通过对点进行排序来检查两条共线线段是否重叠，返回重叠线段

***
#### func CalcLineSegmentIsIntersect(line1 LineSegment[V], line2 LineSegment[V])  bool
<span id="CalcLineSegmentIsIntersect"></span>
> 计算两条线段是否相交

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCalcLineSegmentIsIntersect(t *testing.T) {
	line1 := geometry.NewLineSegment(geometry.NewPoint(1, 1), geometry.NewPoint(3, 5))
	line2 := geometry.NewLineSegment(geometry.NewPoint(0, 5), geometry.NewPoint(3, 6))
	fmt.Println(geometry.CalcLineSegmentIsIntersect(line1, line2))
}

```


</details>


***
#### func CalcLineSegmentSlope(line LineSegment[V])  V
<span id="CalcLineSegmentSlope"></span>
> 计算线段的斜率

***
#### func CalcLineSegmentIntercept(line LineSegment[V])  V
<span id="CalcLineSegmentIntercept"></span>
> 计算线段的截距

***
#### func NewPoint(x V, y V)  Point[V]
<span id="NewPoint"></span>
> 创建一个由 x、y 坐标组成的点

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewPoint(t *testing.T) {
	p := [2]int{1, 1}
	fmt.Println(PointToPos(9, p))
}

```


</details>


***
#### func NewPointCap(x V, y V)  PointCap[V, D]
<span id="NewPointCap"></span>
> 创建一个由 x、y 坐标组成的点，这个点具有一个数据容量

***
#### func NewPointCapWithData(x V, y V, data D)  PointCap[V, D]
<span id="NewPointCapWithData"></span>
> 通过设置数据的方式创建一个由 x、y 坐标组成的点，这个点具有一个数据容量

***
#### func NewPointCapWithPoint(point Point[V], data D)  PointCap[V, D]
<span id="NewPointCapWithPoint"></span>
> 通过设置数据的方式创建一个由已有坐标组成的点，这个点具有一个数据容量

***
#### func CoordinateToPoint(x V, y V)  Point[V]
<span id="CoordinateToPoint"></span>
> 将坐标转换为x、y的坐标数组

***
#### func CoordinateToPos(width V, x V, y V)  V
<span id="CoordinateToPos"></span>
> 将坐标转换为二维数组的顺序位置坐标
>   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值

***
#### func PointToCoordinate(position Point[V]) (x V, y V)
<span id="PointToCoordinate"></span>
> 将坐标数组转换为x和y坐标

***
#### func PointToPos(width V, xy Point[V])  V
<span id="PointToPos"></span>
> 将坐标转换为二维数组的顺序位置
>   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值

***
#### func PosToCoordinate(width V, pos V) (x V, y V)
<span id="PosToCoordinate"></span>
> 通过宽度将一个二维数组的顺序位置转换为xy坐标

***
#### func PosToPoint(width V, pos V)  Point[V]
<span id="PosToPoint"></span>
> 通过宽度将一个二维数组的顺序位置转换为x、y的坐标数组

***
#### func PosToCoordinateX(width V, pos V)  V
<span id="PosToCoordinateX"></span>
> 通过宽度将一个二维数组的顺序位置转换为X坐标

***
#### func PosToCoordinateY(width V, pos V)  V
<span id="PosToCoordinateY"></span>
> 通过宽度将一个二维数组的顺序位置转换为Y坐标

***
#### func PointCopy(point Point[V])  Point[V]
<span id="PointCopy"></span>
> 复制一个坐标数组

***
#### func PointToPosWithMulti(width V, points ...Point[V])  []V
<span id="PointToPosWithMulti"></span>
> 将一组坐标转换为二维数组的顺序位置
>   - 需要确保x的取值范围必须小于width，或者将会得到不正确的值

***
#### func PosToPointWithMulti(width V, positions ...V)  []Point[V]
<span id="PosToPointWithMulti"></span>
> 将一组二维数组的顺序位置转换为一组数组坐标

***
#### func PosSameRow(width V, pos1 V, pos2 V)  bool
<span id="PosSameRow"></span>
> 返回两个顺序位置在同一宽度是否位于同一行

***
#### func DoublePointToCoordinate(point1 Point[V], point2 Point[V]) (x1 V, y1 V, x2 V, y2 V)
<span id="DoublePointToCoordinate"></span>
> 将两个位置转换为 x1, y1, x2, y2 的坐标进行返回

***
#### func CalcProjectionPoint(line LineSegment[V], point Point[V])  Point[V]
<span id="CalcProjectionPoint"></span>
> 计算一个点到一条线段的最近点（即投影点）的。这个函数接收一个点和一条线段作为输入，线段由两个端点组成。
>   - 该函数的主要用于需要计算一个点到一条线段的最近点的情况下

***
#### func GetAdjacentTranslatePos(matrix []T, width P, pos P) (result []P)
<span id="GetAdjacentTranslatePos"></span>
> 获取一个连续位置的矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置

***
#### func GetAdjacentTranslateCoordinateXY(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentTranslateCoordinateXY"></span>
> 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置

***
#### func GetAdjacentTranslateCoordinateYX(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentTranslateCoordinateYX"></span>
> 获取一个基于 y、x 的二维矩阵中，特定位置相邻的最多四个平移方向（上下左右）的位置

***
#### func GetAdjacentDiagonalsPos(matrix []T, width P, pos P) (result []P)
<span id="GetAdjacentDiagonalsPos"></span>
> 获取一个连续位置的矩阵中，特定位置相邻的对角线最多四个方向的位置

***
#### func GetAdjacentDiagonalsCoordinateXY(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentDiagonalsCoordinateXY"></span>
> 获取一个基于 x、y 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置

***
#### func GetAdjacentDiagonalsCoordinateYX(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentDiagonalsCoordinateYX"></span>
> 获取一个基于 tx 的二维矩阵中，特定位置相邻的对角线最多四个方向的位置

***
#### func GetAdjacentPos(matrix []T, width P, pos P) (result []P)
<span id="GetAdjacentPos"></span>
> 获取一个连续位置的矩阵中，特定位置相邻的最多八个方向的位置

***
#### func GetAdjacentCoordinateXY(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentCoordinateXY"></span>
> 获取一个基于 x、y 的二维矩阵中，特定位置相邻的最多八个方向的位置

***
#### func GetAdjacentCoordinateYX(matrix [][]T, x P, y P) (result []Point[P])
<span id="GetAdjacentCoordinateYX"></span>
> 获取一个基于 yx 的二维矩阵中，特定位置相邻的最多八个方向的位置

***
#### func CoordinateMatrixToPosMatrix(matrix [][]V) (width int, posMatrix []V)
<span id="CoordinateMatrixToPosMatrix"></span>
> 将二维矩阵转换为顺序的二维矩阵

***
#### func GetShapeCoverageAreaWithPoint(points ...Point[V]) (left V, right V, top V, bottom V)
<span id="GetShapeCoverageAreaWithPoint"></span>
> 通过传入的一组坐标 points 计算一个图形覆盖的矩形范围

示例代码：
```go

func ExampleGetShapeCoverageAreaWithPoint() {
	var points []geometry.Point[int]
	points = append(points, geometry.NewPoint(1, 1))
	points = append(points, geometry.NewPoint(2, 1))
	points = append(points, geometry.NewPoint(2, 2))
	left, right, top, bottom := geometry.GetShapeCoverageAreaWithPoint(points...)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGetShapeCoverageAreaWithPoint(t *testing.T) {
	Convey("TestGetShapeCoverageAreaWithPoint", t, func() {
		var points []geometry.Point[int]
		points = append(points, geometry.NewPoint(1, 1))
		points = append(points, geometry.NewPoint(2, 1))
		points = append(points, geometry.NewPoint(2, 2))
		left, right, top, bottom := geometry.GetShapeCoverageAreaWithPoint(points...)
		So(left, ShouldEqual, 1)
		So(right, ShouldEqual, 2)
		So(top, ShouldEqual, 1)
		So(bottom, ShouldEqual, 2)
	})
}

```


</details>


***
#### func GetShapeCoverageAreaWithPos(width V, positions ...V) (left V, right V, top V, bottom V)
<span id="GetShapeCoverageAreaWithPos"></span>
> 通过传入的一组坐标 positions 计算一个图形覆盖的矩形范围

示例代码：
```go

func ExampleGetShapeCoverageAreaWithPos() {
	left, right, top, bottom := geometry.GetShapeCoverageAreaWithPos(3, 4, 7, 8)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGetShapeCoverageAreaWithPos(t *testing.T) {
	Convey("TestGetShapeCoverageAreaWithPos", t, func() {
		left, right, top, bottom := geometry.GetShapeCoverageAreaWithPos(3, 4, 7, 8)
		So(left, ShouldEqual, 1)
		So(right, ShouldEqual, 2)
		So(top, ShouldEqual, 1)
		So(bottom, ShouldEqual, 2)
	})
}

```


</details>


***
#### func CoverageAreaBoundless(l V, r V, t V, b V) (left V, right V, top V, bottom V)
<span id="CoverageAreaBoundless"></span>
> 将一个图形覆盖矩形范围设置为无边的
>   - 无边化表示会将多余的部分进行裁剪，例如图形左边从 2 开始的时候，那么左边将会被裁剪到从 0 开始

示例代码：
```go

func ExampleCoverageAreaBoundless() {
	left, right, top, bottom := geometry.CoverageAreaBoundless(1, 2, 1, 2)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCoverageAreaBoundless(t *testing.T) {
	Convey("TestCoverageAreaBoundless", t, func() {
		left, right, top, bottom := geometry.CoverageAreaBoundless(1, 2, 1, 2)
		So(left, ShouldEqual, 0)
		So(right, ShouldEqual, 1)
		So(top, ShouldEqual, 0)
		So(bottom, ShouldEqual, 1)
	})
}

```


</details>


***
#### func GenerateShapeOnRectangle(points ...Point[V]) (result []PointCap[V, bool])
<span id="GenerateShapeOnRectangle"></span>
> 生成一组二维坐标的形状
>   - 这个形状将被在一个刚好能容纳形状的矩形中表示
>   - 为 true 的位置表示了形状的每一个点

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGenerateShapeOnRectangle(t *testing.T) {
	Convey("TestGenerateShapeOnRectangle", t, func() {
		var points geometry.Shape[int]
		points = append(points, geometry.NewPoint(1, 1))
		points = append(points, geometry.NewPoint(2, 1))
		points = append(points, geometry.NewPoint(2, 2))
		fmt.Println(points)
		ps := geometry.GenerateShapeOnRectangle(points.Points()...)
		So(ps[0].GetX(), ShouldEqual, 0)
		So(ps[0].GetY(), ShouldEqual, 0)
		So(ps[0].GetData(), ShouldEqual, true)
		So(ps[1].GetX(), ShouldEqual, 1)
		So(ps[1].GetY(), ShouldEqual, 0)
		So(ps[1].GetData(), ShouldEqual, true)
		So(ps[2].GetX(), ShouldEqual, 0)
		So(ps[2].GetY(), ShouldEqual, 1)
		So(ps[2].GetData(), ShouldEqual, false)
		So(ps[3].GetX(), ShouldEqual, 1)
		So(ps[3].GetY(), ShouldEqual, 1)
		So(ps[3].GetData(), ShouldEqual, true)
	})
}

```


</details>


***
#### func GenerateShapeOnRectangleWithCoordinate(points ...Point[V]) (result [][]bool)
<span id="GenerateShapeOnRectangleWithCoordinate"></span>
> 生成一组二维坐标的形状
>   - 这个形状将被在一个刚好能容纳形状的矩形中表示
>   - 为 true 的位置表示了形状的每一个点

***
#### func GetExpressibleRectangleBySize(width V, height V, minWidth V, minHeight V) (result []Point[V])
<span id="GetExpressibleRectangleBySize"></span>
> 获取一个宽高可表达的所有特定尺寸以上的矩形形状
>   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
>   - 矩形尺寸由大到小

***
#### func GetExpressibleRectangle(width V, height V) (result []Point[V])
<span id="GetExpressibleRectangle"></span>
> 获取一个宽高可表达的所有矩形形状
>   - 返回值表示了每一个矩形右下角的x,y位置（左上角始终为0, 0）
>   - 矩形尺寸由大到小

***
#### func GetRectangleFullPointsByXY(startX V, startY V, endX V, endY V) (result []Point[V])
<span id="GetRectangleFullPointsByXY"></span>
> 通过开始结束坐标获取一个矩形包含的所有点
>   - 例如 1,1 到 2,2 的矩形结果为 1,1 2,1 1,2 2,2

***
#### func GetRectangleFullPoints(width V, height V) (result []Point[V])
<span id="GetRectangleFullPoints"></span>
> 获取一个矩形填充满后包含的所有点

***
#### func GetRectangleFullPos(width V, height V) (result []V)
<span id="GetRectangleFullPos"></span>
> 获取一个矩形填充满后包含的所有位置

***
#### func CalcRectangleCentroid(shape Shape[V])  Point[V]
<span id="CalcRectangleCentroid"></span>
> 计算矩形质心
>   - 非多边形质心计算，仅为顶点的平均值 - 该区域中多边形因子的适当质心

***
#### func SetShapeStringHasBorder()
<span id="SetShapeStringHasBorder"></span>
> 设置 Shape.String 是拥有边界的

***
#### func SetShapeStringNotHasBorder()
<span id="SetShapeStringNotHasBorder"></span>
> 设置 Shape.String 是没有边界的

***
#### func NewShape(points ...Point[V])  Shape[V]
<span id="NewShape"></span>
> 通过多个点生成一个形状进行返回

示例代码：
```go

func ExampleNewShape() {
	shape := geometry.NewShape[int](geometry.NewPoint(3, 0), geometry.NewPoint(3, 1), geometry.NewPoint(3, 2), geometry.NewPoint(3, 3), geometry.NewPoint(4, 3))
	fmt.Println(shape)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewShape(t *testing.T) {
	Convey("TestNewShape", t, func() {
		shape := geometry.NewShape[int](geometry.NewPoint(3, 0), geometry.NewPoint(3, 1), geometry.NewPoint(3, 2), geometry.NewPoint(3, 3), geometry.NewPoint(4, 3))
		fmt.Println(shape)
		points := shape.Points()
		count := shape.PointCount()
		So(count, ShouldEqual, 5)
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 2))
		So(points[3], ShouldEqual, geometry.NewPoint(3, 3))
		So(points[4], ShouldEqual, geometry.NewPoint(4, 3))
	})
}

```


</details>


***
#### func NewShapeWithString(rows []string, point rune) (shape Shape[V])
<span id="NewShapeWithString"></span>
> 通过字符串将指定 rune 转换为点位置生成形状进行返回
>   - 每个点的顺序从上到下，从左到右

示例代码：
```go

func ExampleNewShapeWithString() {
	shape := geometry.NewShapeWithString[int]([]string{"###X###", "###X###", "###X###", "###XX##"}, 'X')
	fmt.Println(shape)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewShapeWithString(t *testing.T) {
	Convey("TestNewShapeWithString", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{"###X###", "###X###", "###X###", "###XX##"}, 'X')
		points := shape.Points()
		count := shape.PointCount()
		So(count, ShouldEqual, 5)
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 2))
		So(points[3], ShouldEqual, geometry.NewPoint(3, 3))
		So(points[4], ShouldEqual, geometry.NewPoint(4, 3))
	})
}

```


</details>


***
#### func CalcBoundingRadius(shape Shape[V])  V
<span id="CalcBoundingRadius"></span>
> 计算多边形转换为圆的半径

***
#### func CalcBoundingRadiusWithCentroid(shape Shape[V], centroid Point[V])  V
<span id="CalcBoundingRadiusWithCentroid"></span>
> 计算多边形在特定质心下圆的半径

***
#### func CalcTriangleTwiceArea(a Point[V], b Point[V], c Point[V])  V
<span id="CalcTriangleTwiceArea"></span>
> 计算由 a、b、c 三个点组成的三角形的面积的两倍

***
#### func IsPointOnEdge(edges []LineSegment[V], point Point[V])  bool
<span id="IsPointOnEdge"></span>
> 检查点是否在 edges 的任意一条边上

***
#### func ProjectionPointToShape(point Point[V], shape Shape[V])  Point[V],  V
<span id="ProjectionPointToShape"></span>
> 将一个点投影到一个多边形上，找到离该点最近的投影点，并返回投影点和距离

***
#### func WithShapeSearchRectangleLowerLimit(minWidth int, minHeight int)  ShapeSearchOption
<span id="WithShapeSearchRectangleLowerLimit"></span>
> 通过矩形宽高下限的方式搜索

***
#### func WithShapeSearchRectangleUpperLimit(maxWidth int, maxHeight int)  ShapeSearchOption
<span id="WithShapeSearchRectangleUpperLimit"></span>
> 通过矩形宽高上限的方式搜索

***
#### func WithShapeSearchRightAngle()  ShapeSearchOption
<span id="WithShapeSearchRightAngle"></span>
> 通过直角的方式进行搜索

***
#### func WithShapeSearchOppositionDirection(direction Direction)  ShapeSearchOption
<span id="WithShapeSearchOppositionDirection"></span>
> 通过限制对立方向的方式搜索
>   - 对立方向例如上不能与下共存

***
#### func WithShapeSearchDirectionCount(count int)  ShapeSearchOption
<span id="WithShapeSearchDirectionCount"></span>
> 通过限制方向数量的方式搜索

***
#### func WithShapeSearchDirectionCountLowerLimit(direction Direction, count int)  ShapeSearchOption
<span id="WithShapeSearchDirectionCountLowerLimit"></span>
> 通过限制特定方向数量下限的方式搜索

***
#### func WithShapeSearchDirectionCountUpperLimit(direction Direction, count int)  ShapeSearchOption
<span id="WithShapeSearchDirectionCountUpperLimit"></span>
> 通过限制特定方向数量上限的方式搜索

***
#### func WithShapeSearchDeduplication()  ShapeSearchOption
<span id="WithShapeSearchDeduplication"></span>
> 通过去重的方式进行搜索
>   - 去重方式中每个点仅会被使用一次

***
#### func WithShapeSearchPointCountLowerLimit(lowerLimit int)  ShapeSearchOption
<span id="WithShapeSearchPointCountLowerLimit"></span>
> 通过限制图形构成的最小点数进行搜索
>   - 当搜索到的图形的点数量低于 lowerLimit 时，将被忽略

***
#### func WithShapeSearchPointCountUpperLimit(upperLimit int)  ShapeSearchOption
<span id="WithShapeSearchPointCountUpperLimit"></span>
> 通过限制图形构成的最大点数进行搜索
>   - 当搜索到的图形的点数量大于 upperLimit 时，将被忽略

***
#### func WithShapeSearchAsc()  ShapeSearchOption
<span id="WithShapeSearchAsc"></span>
> 通过升序的方式进行搜索

***
#### func WithShapeSearchDesc()  ShapeSearchOption
<span id="WithShapeSearchDesc"></span>
> 通过降序的方式进行搜索

***
### Circle `STRUCT`
圆形
```go
type Circle[V generic.SignedNumber] struct {
	Shape[V]
}
```
#### func (Circle) Radius()  V
> 获取圆形半径
***
#### func (Circle) Centroid()  Point[V]
> 获取圆形质心位置
***
#### func (Circle) Overlap(circle Circle[V])  bool
> 与另一个圆是否发生重叠
***
#### func (Circle) Area()  V
> 获取圆形面积
***
#### func (Circle) Length()  V
> 获取圆的周长
***
#### func (Circle) CentroidDistance(circle Circle[V])  V
> 计算与另一个圆的质心距离
***
### FloorPlan `STRUCT`
平面图
```go
type FloorPlan []string
```
#### func (FloorPlan) IsFree(point Point[int])  bool
> 检查位置是否为空格
***
#### func (FloorPlan) IsInBounds(point Point[int])  bool
> 检查位置是否在边界内
***
#### func (FloorPlan) Put(point Point[int], c rune)
> 设置平面图特定位置的字符
***
#### func (FloorPlan) String()  string
> 获取平面图结果
***
### Direction `STRUCT`
方向
```go
type Direction uint8
```
### LineSegment `STRUCT`
通过两个点表示一根线段
```go
type LineSegment[V generic.SignedNumber] [2]Point[V]
```
#### func (LineSegment) GetPoints()  [2]Point[V]
> 获取该线段的两个点
***
#### func (LineSegment) GetStart()  Point[V]
> 获取该线段的开始位置
***
#### func (LineSegment) GetEnd()  Point[V]
> 获取该线段的结束位置
***
#### func (LineSegment) GetLength()  V
> 获取该线段的长度
***
### LineSegmentCap `STRUCT`
可以包含一份额外数据的线段
```go
type LineSegmentCap[V generic.SignedNumber, Data any] struct {
	LineSegment[V]
	Data Data
}
```
### Point `STRUCT`
表示了一个由 x、y 坐标组成的点
```go
type Point[V generic.SignedNumber] [2]V
```
#### func (Point) GetX()  V
> 返回该点的 x 坐标
***
#### func (Point) GetY()  V
> 返回该点的 y 坐标
***
#### func (Point) GetXY() (x V, y V)
> 返回该点的 x、y 坐标
***
#### func (Point) GetPos(width V)  V
> 返回该点位于特定宽度的二维数组的顺序位置
***
#### func (Point) GetOffset(x V, y V)  Point[V]
> 获取偏移后的新坐标
***
#### func (Point) Negative()  bool
> 返回该点是否是一个负数坐标
***
#### func (Point) OutOf(minWidth V, minHeight V, maxWidth V, maxHeight V)  bool
> 返回该点在特定宽高下是否越界f
***
#### func (Point) Equal(point Point[V])  bool
> 返回两个点是否相等
***
#### func (Point) Copy()  Point[V]
> 复制一个点位置
***
#### func (Point) Add(point Point[V])  Point[V]
> 得到加上 point 后的点
***
#### func (Point) Sub(point Point[V])  Point[V]
> 得到减去 point 后的点
***
#### func (Point) Mul(point Point[V])  Point[V]
> 得到乘以 point 后的点
***
#### func (Point) Div(point Point[V])  Point[V]
> 得到除以 point 后的点
***
#### func (Point) Abs()  Point[V]
> 返回位置的绝对值
***
#### func (Point) Max(point Point[V])  Point[V]
> 返回两个位置中每个维度的最大值组成的新的位置
***
#### func (Point) Min(point Point[V])  Point[V]
> 返回两个位置中每个维度的最小值组成的新的位置
***
### PointCap `STRUCT`
表示了一个由 x、y 坐标组成的点，这个点具有一个数据容量
```go
type PointCap[V generic.SignedNumber, D any] struct {
	Point[V]
	Data D
}
```
### Shape `STRUCT`
通过多个点表示了一个形状
```go
type Shape[V generic.SignedNumber] []Point[V]
```
#### func (Shape) Points()  []Point[V]
> 获取这个形状的所有点
示例代码：
```go

func ExampleShape_Points() {
	shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
	points := shape.Points()
	fmt.Println(points)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestShape_Points(t *testing.T) {
	Convey("TestShape_Points", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
		points := shape.Points()
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(2, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[3], ShouldEqual, geometry.NewPoint(4, 1))
	})
}

```


</details>


***
#### func (Shape) PointCount()  int
> 获取这个形状的点数量
示例代码：
```go

func ExampleShape_PointCount() {
	shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
	fmt.Println(shape.PointCount())
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestShape_PointCount(t *testing.T) {
	Convey("TestShape_PointCount", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
		So(shape.PointCount(), ShouldEqual, 4)
	})
}

```


</details>


***
#### func (Shape) Contains(point Point[V])  bool
> 返回该形状中是否包含点
***
#### func (Shape) ToCircle()  Circle[V]
> 将形状转换为圆形进行处理
>   - 当形状非圆形时将会产生意外情况
***
#### func (Shape) String()  string
> 将该形状转换为可视化的字符串进行返回
示例代码：
```go

func ExampleShape_String() {
	shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
	fmt.Println(shape)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestShape_String(t *testing.T) {
	Convey("TestShape_String", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##"}, 'X')
		str := shape.String()
		So(str, ShouldEqual, "[[3 0] [2 1] [3 1] [4 1]]\n# X #\nX X X")
	})
}

```


</details>


***
#### func (Shape) ShapeSearch(options ...ShapeSearchOption) (result []Shape[V])
> 获取该形状中包含的所有图形组合及其位置
>   - 需要注意的是，即便图形最终表示为相同的，但是只要位置组合顺序不同，那么也将被认定为一种图形组合
>   - [[1 0] [1 1] [1 2]] 和 [[1 1] [1 0] [1 2]] 可以被视为两个图形组合
>   - 返回的坐标为原始形状的坐标
> 
> 可通过可选项对搜索结果进行过滤
示例代码：
```go

func ExampleShape_ShapeSearch() {
	shape := geometry.NewShapeWithString[int]([]string{"###X###", "##XXX##", "###X###"}, 'X')
	shapes := shape.ShapeSearch(geometry.WithShapeSearchDeduplication(), geometry.WithShapeSearchDesc())
	for _, shape := range shapes {
		fmt.Println(shape)
	}
}

```

***
#### func (Shape) Edges() (edges []LineSegment[V])
> 获取该形状每一条边
>   - 该形状需要最少由3个点组成，否则将不会返回任意一边
***
#### func (Shape) IsPointOnEdge(point Point[V])  bool
> 检查点是否在该形状的一条边上
***
### ShapeSearchOption `STRUCT`
图形搜索可选项，用于 Shape.ShapeSearch 搜索支持
```go
type ShapeSearchOption func(options *shapeSearchOptions)
```
