# Astar

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

astar 提供用于实现 A* 算法的函数和数据结构。A* 算法是一种常用的路径搜索算法，用于在图形或网络中找到最短路径。该包旨在简化 A* 算法的实现过程，并提供一致的接口和易于使用的功能。
主要特性：
  - 图形表示：astar 包支持使用图形或网络来表示路径搜索的环境。您可以定义节点和边，以构建图形，并在其中执行路径搜索。
  - A* 算法：该包提供了 A* 算法的实现，用于在图形中找到最短路径。A* 算法使用启发式函数来评估节点的优先级，并选择最有希望的节点进行扩展，以达到最短路径的目标。
  - 自定义启发式函数：您可以根据具体问题定义自己的启发式函数，以指导 A* 算法的搜索过程。启发式函数用于估计从当前节点到目标节点的代价，以帮助算法选择最佳路径。
  - 可定制性：astar 包提供了一些可定制的选项，以满足不同场景下的需求。您可以设置节点的代价、边的权重等参数，以调整算法的行为。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[Find](#Find)|使用 A* 算法在导航网格上查找从起点到终点的最短路径，并返回路径上的节点序列。


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`INTERFACE`|[Graph](#struct_Graph)|适用于 A* 算法的图数据结构接口定义，表示导航网格，其中包含了节点和连接节点的边。

</details>


***
## 详情信息
#### func Find\[Node comparable, V generic.SignedNumber\](graph Graph[Node], start Node, end Node, cost func (a Node)  V, heuristic func (a Node)  V) []Node
<span id="Find"></span>
> 使用 A* 算法在导航网格上查找从起点到终点的最短路径，并返回路径上的节点序列。
> 
> 参数：
>   - graph: 图对象，类型为 Graph[Node]，表示导航网格。
>   - start: 起点节点，类型为 Node，表示路径的起点。
>   - end: 终点节点，类型为 Node，表示路径的终点。
>   - cost: 路径代价函数，类型为 func(a, b Node) V，用于计算两个节点之间的代价。
>   - heuristic: 启发函数，类型为 func(a, b Node) V，用于估计从当前节点到目标节点的启发式代价。
> 
> 返回值：
>   - []Node: 节点序列，表示从起点到终点的最短路径。如果找不到路径，则返回空序列。
> 
> 注意事项：
>   - graph 对象表示导航网格，其中包含了节点和连接节点的边。
>   - start 和 end 分别表示路径的起点和终点。
>   - cost 函数用于计算两个节点之间的代价，可以根据实际情况自定义实现。
>   - heuristic 函数用于估计从当前节点到目标节点的启发式代价，可以根据实际情况自定义实现。
>   - 函数使用了 A* 算法来搜索最短路径。
>   - 函数内部使用了堆数据结构来管理待处理的节点。
>   - 函数返回一个节点序列，表示从起点到终点的最短路径。如果找不到路径，则返回空序列。

**示例代码：**

```go

func ExampleFind() {
	graph := Graph{FloorPlan: geometry.FloorPlan{"===========", "X XX  X   X", "X  X   XX X", "X XX      X", "X     XXX X", "X XX  X   X", "X XX  X   X", "==========="}}
	paths := astar.Find[geometry.Point[int], int](graph, geometry.NewPoint(1, 1), geometry.NewPoint(8, 6), func(a, b geometry.Point[int]) int {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a, b))
	}, func(a, b geometry.Point[int]) int {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a, b))
	})
	for _, path := range paths {
		graph.Put(path, '.')
	}
	fmt.Println(graph)
}

```

***
<span id="struct_Graph"></span>
### Graph `INTERFACE`
适用于 A* 算法的图数据结构接口定义，表示导航网格，其中包含了节点和连接节点的边。
```go
type Graph[Node comparable] interface {
	Neighbours(node Node) []Node
}
```
#### func (Graph) Neighbours(point geometry.Point[int])  []geometry.Point[int]
***
