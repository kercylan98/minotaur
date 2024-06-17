package balancer

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"hash/crc32"
	"sort"
)

// NewConsistentHashWeight 创建一致性哈希权重负载均衡器
func NewConsistentHashWeight[I constraints.Ordered, T Item[I]](replicas int) *ConsistentHashWeight[I, T] {
	return &ConsistentHashWeight[I, T]{
		hashRing: make(map[uint32]T),
		replicas: replicas,
	}
}

// ConsistentHashWeight 一致性哈希权重负载均衡器
type ConsistentHashWeight[I constraints.Ordered, T Item[I]] struct {
	replicas int // 每个实例的虚拟节点数
	hashRing map[uint32]T
	keys     []uint32
}

func (c *ConsistentHashWeight[I, T]) Get(id I) (t T) {
	for _, i := range c.hashRing {
		if i.GetId() == id {
			return t
		}
	}
	return t
}

func (c *ConsistentHashWeight[I, T]) Select(opts ...*SelectOptions) (t T, err error) {
	opt := NewSelectOptions().Apply(opts...)
	if len(c.keys) == 0 {
		err = ErrNoInstance
		return
	}

	var hash uint32
	if len(opt.ConsistencyKey) > 0 {
		hash = crc32.ChecksumIEEE([]byte(opt.ConsistencyKey))
	}
	idx := c.search(hash)
	return c.hashRing[c.keys[idx]], nil
}

func (c *ConsistentHashWeight[I, T]) Update(instance T) {
	c.Remove(instance.GetId())
	c.Add(instance)
}

func (c *ConsistentHashWeight[I, T]) Add(instance T) {
	for i := 0; i < c.replicas; i++ {
		key := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v-%d", instance.GetId(), i)))
		c.hashRing[key] = instance
		c.keys = append(c.keys, key)
	}
	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
}

func (c *ConsistentHashWeight[I, T]) Remove(instanceId I) {
	for i := 0; i < c.replicas; i++ {
		key := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v-%d", instanceId, i)))
		delete(c.hashRing, key)
	}
	c.keys = c.reorganizeKeys()
}

func (c *ConsistentHashWeight[I, T]) GetInstances() []T {
	instances := make([]T, 0, len(c.hashRing))
	for _, instance := range c.hashRing {
		instances = append(instances, instance)
	}
	return instances
}

// search 在哈希环中找到最接近的节点索引
func (c *ConsistentHashWeight[I, T]) search(hash uint32) int {
	var idx int
	if hash != 0 {
		idx = sort.Search(len(c.keys), func(i int) bool {
			return c.keys[i] >= hash
		})
	}

	if hash == 0 || idx >= len(c.keys) {
		// 获取权重最大的节点
		var maxWeight int
		var maxKey uint32
		for key, t := range c.hashRing {
			if t.GetWeight() > maxWeight {
				maxWeight = t.GetWeight()
				maxKey = key
			}
		}
		// 查找 key 所在索引
		for i, key := range c.keys {
			if key == maxKey {
				idx = i
				break
			}
		}
	}

	return idx
}

// reorganizeKeys 重新整理哈希键列表
func (c *ConsistentHashWeight[I, T]) reorganizeKeys() []uint32 {
	var keys []uint32
	for key := range c.hashRing {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}
