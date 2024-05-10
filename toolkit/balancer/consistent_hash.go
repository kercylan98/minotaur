package balancer

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"hash/crc32"
	"sort"
)

// NewConsistentHash 创建一个一致性哈希负载均衡器
func NewConsistentHash[I constraints.Ordered, T Item[I]]() *ConsistentHash[I, T] {
	return &ConsistentHash[I, T]{}
}

// ConsistentHash 一致性哈希负载均衡器
type ConsistentHash[I constraints.Ordered, T Item[I]] struct {
	replicas int // 每个实例的虚拟节点数
	hashRing map[uint32]T
	keys     []uint32
}

func (c *ConsistentHash[I, T]) Select(opts ...*SelectOptions) (t T, err error) {
	opt := NewSelectOptions().Apply(opts...)
	if len(c.keys) == 0 {
		err = ErrNoInstance
		return
	}

	hash := crc32.ChecksumIEEE([]byte(opt.ConsistencyKey))
	idx := c.search(hash)
	return c.hashRing[c.keys[idx]], nil
}

func (c *ConsistentHash[I, T]) Add(instance T) {
	for i := 0; i < c.replicas; i++ {
		key := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v-%d", instance.GetId(), i)))
		c.hashRing[key] = instance
		c.keys = append(c.keys, key)
	}
	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
}

func (c *ConsistentHash[I, T]) Remove(instance T) {
	for i := 0; i < c.replicas; i++ {
		key := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v-%d", instance.GetId(), i)))
		delete(c.hashRing, key)
	}
	c.keys = c.reorganizeKeys()
}

func (c *ConsistentHash[I, T]) GetInstances() []T {
	instances := make([]T, 0, len(c.hashRing))
	for _, instance := range c.hashRing {
		instances = append(instances, instance)
	}
	return instances
}

// search 在哈希环中找到最接近的节点索引
func (c *ConsistentHash[I, T]) search(hash uint32) int {
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})

	if idx >= len(c.keys) {
		return 0
	}

	return idx
}

// reorganizeKeys 重新整理哈希键列表
func (c *ConsistentHash[I, T]) reorganizeKeys() []uint32 {
	var keys []uint32
	for key := range c.hashRing {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}
