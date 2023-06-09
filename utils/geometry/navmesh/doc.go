// Package navmesh 提供了用于导航网格处理的函数和数据结构。导航网格是一种常用的数据结构，用于在游戏开发和虚拟环境中进行路径规划和导航。该包旨在简化导航网格的创建、查询和操作过程，并提供高效的导航功能。
// 主要特性：
//   - 导航网格表示：navmesh 包支持使用导航网格来表示虚拟环境中的可行走区域和障碍物。您可以定义多边形区域和连接关系，以构建导航网格，并在其中执行路径规划和导航。
//   - 导航算法：采用了 A* 算法作为导航算法，用于在导航网格中找到最短路径或最优路径。这些算法使用启发式函数和代价评估来指导路径搜索，并提供高效的路径规划能力。
package navmesh
