package gossip

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

// HashRing 是一种哈希环算法，用于实现负载均衡和数据分片
type HashRing struct {
	nodes      []string          // 存储节点列表
	sortedKeys []uint32          // 已排序的哈希值
	nodeMap    map[uint32]string // 哈希值到节点的映射
	replicas   int               // 每个节点对应的虚拟节点数量
}

// NewHashRing 创建一个新的哈希环，并设置虚拟节点数量
func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		nodeMap:  make(map[uint32]string),
		replicas: replicas,
	}
}

// hash 计算一个字符串的哈希值
func (h *HashRing) hash(key string) uint32 {
	hash := sha1.New()
	hash.Write([]byte(key))
	bs := hash.Sum(nil)
	return uint32(bs[0])<<24 | uint32(bs[1])<<16 | uint32(bs[2])<<8 | uint32(bs[3])
}

// AddNode 添加一个节点到哈希环中
func (h *HashRing) AddNode(node string) {
	for i := 0; i < h.replicas; i++ {
		hash := h.hash(fmt.Sprintf("%s:%d", node, i))
		if _, exist := h.nodeMap[hash]; exist {
			return
		}
		h.sortedKeys = append(h.sortedKeys, hash)
		h.nodeMap[hash] = node
	}
	sort.Slice(h.sortedKeys, func(i, j int) bool {
		return h.sortedKeys[i] < h.sortedKeys[j]
	})
	h.nodes = append(h.nodes, node)
}

// GetNeighbours 根据节点名称获取最近的 numNeighbours 个节点
func (h *HashRing) GetNeighbours(node string, numNeighbours int) []string {
	if numNeighbours > len(h.nodes)-1 {
		numNeighbours = len(h.nodes) - 1
	}

	// 找到所有与此物理节点相关的虚拟节点的哈希值
	var nodeHashes []uint32
	for i := 0; i < h.replicas; i++ {
		hash := h.hash(fmt.Sprintf("%s:%d", node, i))
		if _, ok := h.nodeMap[hash]; ok {
			nodeHashes = append(nodeHashes, hash)
		}
	}

	if len(nodeHashes) == 0 {
		return nil
	}

	// 在哈希环中找到物理节点对应的其中一个虚拟节点位置
	var nodeHash uint32 = nodeHashes[0] // 使用第一个找到的虚拟节点哈希值
	idx := sort.Search(len(h.sortedKeys), func(i int) bool {
		return h.sortedKeys[i] >= nodeHash
	})

	var neighbours []string
	seen := make(map[string]bool)

	// 顺时针方向查找最近的 numNeighbours 个不同节点
	for i := 1; len(neighbours) < numNeighbours; i++ {
		neighbourIdx := (idx + i) % len(h.sortedKeys)
		neighbour := h.nodeMap[h.sortedKeys[neighbourIdx]]

		// 跳过当前节点本身
		if neighbour == node {
			continue
		}

		if !seen[neighbour] { // 确保不重复添加相同的物理节点
			neighbours = append(neighbours, neighbour)
			seen[neighbour] = true
		}
	}

	return neighbours
}

// RemoveNode 删除一个节点及其虚拟节点
func (h *HashRing) RemoveNode(node string) {
	// 删除与该节点相关的所有虚拟节点
	for i := 0; i < h.replicas; i++ {
		hash := h.hash(fmt.Sprintf("%s:%d", node, i))
		// 从 nodeMap 中删除该虚拟节点
		delete(h.nodeMap, hash)

		// 从 sortedKeys 中删除该哈希值
		idx := sort.Search(len(h.sortedKeys), func(i int) bool {
			return h.sortedKeys[i] >= hash
		})
		// 确保哈希值存在，避免删除错误的索引
		if idx < len(h.sortedKeys) && h.sortedKeys[idx] == hash {
			h.sortedKeys = append(h.sortedKeys[:idx], h.sortedKeys[idx+1:]...)
		}
	}

	// 从 nodes 列表中删除该物理节点
	for i, n := range h.nodes {
		if n == node {
			h.nodes = append(h.nodes[:i], h.nodes[i+1:]...)
			break
		}
	}
}
