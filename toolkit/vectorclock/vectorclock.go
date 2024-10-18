package vectorclock

import (
	"fmt"
	"sort"
)

type Node = string

// Ordering 定义事件的顺序类型
type Ordering int

const (
	Same       Ordering = iota // 表示两个向量时钟相同
	Before                     // 当前时钟发生在另一个时钟之前
	After                      // 当前时钟发生在另一个时钟之后
	Concurrent                 // 两个时钟是并发的（无法确定先后顺序）
)

// VectorClock 表示向量时钟结构
type VectorClock struct {
	versions map[Node]int64 // 存储每个节点的版本
}

// NewVectorClock 创建一个新的向量时钟
func NewVectorClock() *VectorClock {
	return &VectorClock{
		versions: make(map[Node]int64),
	}
}

// Version 返回当前时钟的版本信息
func (vc *VectorClock) Version() map[Node]int64 {
	return vc.versions
}

// Increment 对特定节点的版本进行递增，并返回更新后的向量时钟
func (vc *VectorClock) Increment(node Node, v int64) *VectorClock {
	vc.versions[node] = vc.versions[node] + v
	return vc
}

// CompareTo 比较两个向量时钟，返回事件顺序
func (vc *VectorClock) CompareTo(that *VectorClock) Ordering {
	hasBefore := false
	hasAfter := false

	// 遍历当前时钟的所有节点版本，比较另一个时钟中的版本
	for node, time := range vc.versions {
		thatTime, exists := that.versions[node]
		if !exists {
			// 如果在另一个时钟中不存在这个节点，则当前时钟是"After"
			hasAfter = true
		} else if time > thatTime {
			hasAfter = true
		} else if time < thatTime {
			hasBefore = true
		}
	}

	// 检查那些仅存在于另一个时钟中的节点
	for node, time := range that.versions {
		if _, exists := vc.versions[node]; !exists {
			hasBefore = true
		} else if time > vc.versions[node] {
			hasBefore = true
		}
	}

	// 根据比较结果返回顺序
	switch {
	case hasAfter && hasBefore:
		return Concurrent
	case hasAfter:
		return After
	case hasBefore:
		return Before
	default:
		return Same
	}
}

// Merge 合并两个向量时钟，取每个节点的最大时间戳
func (vc *VectorClock) Merge(that *VectorClock) *VectorClock {
	merged := NewVectorClock()

	// 将当前时钟的所有版本复制到新时钟中
	for node, time := range vc.versions {
		merged.versions[node] = time
	}

	// 将另一个时钟中的版本合并到新时钟中
	for node, time := range that.versions {
		if merged.versions[node] < time {
			merged.versions[node] = time
		}
	}

	return merged
}

// Prune 移除某个节点的版本信息
func (vc *VectorClock) Prune(node Node) *VectorClock {
	delete(vc.versions, node)
	return vc
}

// String 提供向量时钟的字符串表示，方便调试
func (vc *VectorClock) String() string {
	nodes := make([]Node, 0, len(vc.versions))
	for node := range vc.versions {
		nodes = append(nodes, node)
	}
	// 对节点进行排序，保证输出的一致性
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i] < nodes[j]
	})

	str := "VectorClock("
	for i, node := range nodes {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%s -> %d", node, vc.versions[node])
	}
	str += ")"
	return str
}
