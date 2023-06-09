// Package dp (DistributionPattern) 提供用于在二维数组中根据不同的特征标记为数组成员建立分布链接的函数和数据结构。该包的目标是实现快速查找与给定位置成员具有相同特征且位置紧邻的其他成员。
// 主要特性：
//   - 分布链接机制：dp 包提供了一种分布链接的机制，可以根据成员的特征将它们链接在一起。这样，可以快速查找与给定成员具有相同特征且位置紧邻的其他成员。
//   - 二维数组支持：该包支持在二维数组中建立分布链接。可以将二维数组中的成员视为节点，并根据其特征进行链接。
//   - 快速查找功能：使用 dp 包提供的函数，可以快速查找与给定位置成员具有相同特征且位置紧邻的其他成员。这有助于在二维数组中进行相关性分析或查找相邻成员。
package dp
