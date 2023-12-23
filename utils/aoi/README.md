# AOI (Area of Interest)

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/aoi)

AOI (Area of Interest) 是一种常见的游戏服务器技术，用于处理大量玩家在同一空间内的交互问题。在 `Minotaur` 中，我们提供了一个基于 Go 语言的 AOI 实现。

## TwoDimensional [`二维AOI`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/aoi#TwoDimensional)

[`TwoDimensional`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/aoi#TwoDimensional)
是一个二维 AOI 的实现。它提供了添加实体、删除实体、刷新实体、获取焦点实体等方法。每个实体需要实现 [`TwoDimensionalEntity`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/aoi#TwoDimensionalEntity) 接口，该接口包含了获取实体 ID、获取实体坐标、获取实体视野半径等方法。