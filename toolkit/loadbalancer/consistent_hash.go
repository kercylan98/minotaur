package loadbalancer

import (
	"github.com/kercylan98/minotaur/toolkit/convert"
	"hash/fnv"
	"sort"
	"sync"
)

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		keys:     []int{},
		hashMap:  make(map[int]string),
		mutex:    sync.RWMutex{},
	}
}

type ConsistentHash struct {
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环上的所有节点的哈希值
	hashMap  map[int]string // 哈希值到真实节点的映射
	mutex    sync.RWMutex   // 用于保护数据结构
}

// Add 添加一个节点到哈希环
func (c *ConsistentHash) Add(node string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < c.replicas; i++ {
		hash := c.hash(node + convert.IntToString(i))
		c.keys = append(c.keys, hash)
		c.hashMap[hash] = node
	}
	sort.Ints(c.keys)
}

// Remove 从哈希环中移除一个节点
func (c *ConsistentHash) Remove(node string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < c.replicas; i++ {
		hash := c.hash(node + convert.IntToString(i))
		delete(c.hashMap, hash)
		// 从 keys 中移除节点的哈希值
		for j, k := range c.keys {
			if k == hash {
				c.keys = append(c.keys[:j], c.keys[j+1:]...)
				break
			}
		}
	}
}

// Get 返回给定 key 所在的节点
func (c *ConsistentHash) Get(key string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if len(c.keys) == 0 {
		return ""
	}

	hash := c.hash(key)
	// 顺时针找到第一个比 key 大的哈希值，即对应的节点
	idx := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= hash })
	if idx == len(c.keys) {
		idx = 0 // 如果 key 大于所有哈希值，则返回第一个节点
	}
	return c.hashMap[c.keys[idx]]
}

// hash 计算字符串的哈希值
func (c *ConsistentHash) hash(key string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32())
}
