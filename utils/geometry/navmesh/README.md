# Navmesh

navmesh 提供了用于导航网格处理的函数和数据结构。导航网格是一种常用的数据结构，用于在游戏开发和虚拟环境中进行路径规划和导航。该包旨在简化导航网格的创建、查询和操作过程，并提供高效的导航功能。
主要特性：
  - 导航网格表示：navmesh 包支持使用导航网格来表示虚拟环境中的可行走区域和障碍物。您可以定义多边形区域和连接关系，以构建导航网格，并在其中执行路径规划和导航。
  - 导航算法：采用了 A* 算法作为导航算法，用于在导航网格中找到最短路径或最优路径。这些算法使用启发式函数和代价评估来指导路径搜索，并提供高效的路径规划能力。

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/navmesh)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewNavMesh](#NewNavMesh)|创建一个新的导航网格，并返回一个指向该导航网格的指针。


> 结构体定义

|结构体|描述
|:--|:--
|[NavMesh](#navmesh)|暂无描述...

</details>


#### func NewNavMesh(shapes []geometry.Shape[V], meshShrinkAmount V)  *NavMesh[V]
<span id="NewNavMesh"></span>
> 创建一个新的导航网格，并返回一个指向该导航网格的指针。
> 
> 参数：
>   - shapes: 形状切片，类型为 []geometry.Shape[V]，表示导航网格中的形状。
>   - meshShrinkAmount: 网格缩小量，类型为 V，表示导航网格的缩小量。
> 
> 返回值：
>   - *NavMesh[V]: 指向创建的导航网格的指针。
> 
> 注意事项：
>   - 导航网格的形状可以是任何几何形状。
>   - meshShrinkAmount 表示导航网格的缩小量，用于在形状之间创建链接时考虑形状的缩小效果。
>   - 函数内部使用了泛型类型参数 V，可以根据需要指定形状的坐标类型。
>   - 函数返回一个指向创建的导航网格的指针。
> 
> 使用建议：
>   - 确保 NavMesh 计算精度的情况下，V 建议使用 float64 类型
***
### NavMesh

```go
type NavMesh[V generic.SignedNumber] struct {
	meshShapes       []*shape[V]
	meshShrinkAmount V
}
```
#### func (*NavMesh) Neighbours(node *shape[V])  []*shape[V]
> 实现 astar.Graph 的接口，用于向 A* 算法提供相邻图形
***
#### func (*NavMesh) Find(point geometry.Point[V], maxDistance V) (distance V, findPoint geometry.Point[V], findShape geometry.Shape[V])
> 用于在 NavMesh 中查找离给定点最近的形状，并返回距离、找到的点和找到的形状。
> 
> 参数：
>   - point: 给定的点，类型为 geometry.Point[V]，表示一个 V 维度的点坐标。
>   - maxDistance: 最大距离，类型为 V，表示查找的最大距离限制。
> 
> 返回值：
>   - distance: 距离，类型为 V，表示离给定点最近的形状的距离。
>   - findPoint: 找到的点，类型为 geometry.Point[V]，表示离给定点最近的点坐标。
>   - findShape: 找到的形状，类型为 geometry.Shape[V]，表示离给定点最近的形状。
> 
> 注意事项：
>   - 如果给定点在 NavMesh 中的某个形状内部或者在形状的边上，距离为 0，找到的形状为该形状，找到的点为给定点。
>   - 如果给定点不在任何形状内部或者形状的边上，将计算给定点到每个形状的距离，并找到最近的形状和对应的点。
>   - 距离的计算采用几何学中的投影点到形状的距离。
>   - 函数返回离给定点最近的形状的距离、找到的点和找到的形状。
***
#### func (*NavMesh) FindPath(start geometry.Point[V], end geometry.Point[V]) (result []geometry.Point[V])
> 函数用于在 NavMesh 中查找从起点到终点的路径，并返回路径上的点序列。
> 
> 参数：
>   - start: 起点，类型为 geometry.Point[V]，表示路径的起始点。
>   - end: 终点，类型为 geometry.Point[V]，表示路径的终点。
> 
> 返回值：
>   - result: 路径上的点序列，类型为 []geometry.Point[V]。
> 
> 注意事项：
>   - 函数首先根据起点和终点的位置，找到离它们最近的形状作为起点形状和终点形状。
>   - 如果起点或终点不在任何形状内部，且 NavMesh 的 meshShrinkAmount 大于0，则会考虑缩小的形状。
>   - 使用 A* 算法在 NavMesh 上搜索从起点形状到终点形状的最短路径。
>   - 使用漏斗算法对路径进行优化，以得到最终的路径点序列。
***
