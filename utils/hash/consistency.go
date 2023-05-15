package hash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"strings"
)

func NewConsistency(replicas int) *Consistency {
	return &Consistency{
		replicas: replicas,
	}
}

// Consistency 一致性哈希生成
//
// https://blog.csdn.net/zhpCSDN921011/article/details/126845397
type Consistency struct {
	replicas int         // 虚拟节点的数量
	keys     []int       // 所有虚拟节点的哈希值
	hashMap  map[int]int // 虚拟节点的哈希值: 节点（虚拟节点映射到真实节点）
}

// AddNode 添加节点
func (slf *Consistency) AddNode(keys ...int) {
	if slf.hashMap == nil {
		slf.hashMap = map[int]int{}
	}
	if slf.replicas == 0 {
		slf.replicas = 3
	}
	for _, key := range keys {
		for i := 0; i < slf.replicas; i++ {
			// 计算虚拟节点哈希值
			hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + strconv.Itoa(key))))

			// 存储虚拟节点哈希值
			slf.keys = append(slf.keys, hash)

			// 存入map做映射
			slf.hashMap[hash] = key
		}
	}

	//排序哈希值，下面匹配的时候要二分搜索
	sort.Ints(slf.keys)
}

// PickNode 获取与 key 最接近的节点
func (slf *Consistency) PickNode(key any) int {
	if len(slf.keys) == 0 {
		return 0
	}

	partitionKey := fmt.Sprintf("%#v", key)
	beg := strings.Index(partitionKey, "{")
	end := strings.Index(partitionKey, "}")
	if beg != -1 && !(end == -1 || end == beg+1) {
		partitionKey = partitionKey[beg+1 : end]
	}

	// 计算传入key的哈希值
	hash := int(crc32.ChecksumIEEE([]byte(partitionKey)))

	// sort.Search使用二分查找满足m.Keys[i]>=hash的最小哈希值
	idx := sort.Search(len(slf.keys), func(i int) bool { return slf.keys[i] >= hash })

	// 若key的hash值大于最后一个虚拟节点的hash值，则选择第一个虚拟节点
	if idx == len(slf.keys) {
		idx = 0
	}

	return slf.hashMap[slf.keys[idx]]

}
