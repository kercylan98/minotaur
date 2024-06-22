package mappings

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/twmb/murmur3"
)

type HashFunc[K constraints.Hash] func(size int, key K) int

func NewStringBucket[V any](bucketSize int) *Bucket[string, V] {
	return NewBucket[string, V](bucketSize, func(size int, key string) int {
		hash := murmur3.Sum32([]byte(key))
		index := int(hash) % size
		return index
	})
}

func NewBucket[K constraints.Hash, V any](bucketSize int, hashFunc HashFunc[K]) *Bucket[K, V] {
	bucket := &Bucket[K, V]{
		hashFunc: hashFunc,
		buckets:  make([]*haxmap.Map[K, V], bucketSize),
	}

	for i := 0; i < bucketSize; i++ {
		bucket.buckets[i] = haxmap.New[K, V]()
	}

	return bucket
}

// Bucket 并发安全的基于哈希桶实现的字典结构
type Bucket[K constraints.Hash, V any] struct {
	hashFunc HashFunc[K]
	buckets  []*haxmap.Map[K, V]
}

func (b *Bucket[K, V]) GetBucket(key K) *haxmap.Map[K, V] {
	return b.buckets[b.hashFunc(len(b.buckets), key)]
}

func (b *Bucket[K, V]) Get(key K) (V, bool) {
	bucket := b.GetBucket(key)
	return bucket.Get(key)
}

func (b *Bucket[K, V]) Set(key K, value V) {
	bucket := b.GetBucket(key)
	bucket.Set(key, value)
}

func (b *Bucket[K, V]) Del(key K) {
	bucket := b.GetBucket(key)
	bucket.Del(key)
}

func (b *Bucket[K, V]) Len() int {
	var length uintptr
	for _, bucket := range b.buckets {
		length += bucket.Len()
	}
	return int(length)
}

func (b *Bucket[K, V]) Clear() {
	for _, bucket := range b.buckets {
		bucket.Clear()
	}
}
