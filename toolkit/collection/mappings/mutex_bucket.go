package mappings

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sync"
)

func NewMutexBucket[K constraints.Hash, V any](bucketSize int, hashFunc HashFunc[K]) *MutexBucket[K, V] {
	bucket := &MutexBucket[K, V]{
		hashFunc: hashFunc,
		buckets:  make([]*MutexBucketItem[K, V], bucketSize),
	}

	for i := 0; i < bucketSize; i++ {
		bucket.buckets[i] = &MutexBucketItem[K, V]{
			kv: make(map[K]V),
		}
	}

	return bucket
}

// MutexBucket 基于 sync 包实现的并发安全的哈希桶
type MutexBucket[K constraints.Hash, V any] struct {
	hashFunc HashFunc[K]
	buckets  []*MutexBucketItem[K, V]
}

func (b *MutexBucket[K, V]) GetBucket(key K) *MutexBucketItem[K, V] {
	return b.buckets[b.hashFunc(len(b.buckets), key)]
}

func (b *MutexBucket[K, V]) Get(key K) (V, bool) {
	bucket := b.GetBucket(key)
	bucket.RLock()
	v, exists := bucket.kv[key]
	bucket.RUnlock()
	return v, exists
}

func (b *MutexBucket[K, V]) Set(key K, value V) {
	bucket := b.GetBucket(key)
	bucket.Lock()
	bucket.kv[key] = value
	bucket.Unlock()
}

func (b *MutexBucket[K, V]) Del(key K) {
	bucket := b.GetBucket(key)
	bucket.Lock()
	delete(bucket.kv, key)
	bucket.Unlock()
}

func (b *MutexBucket[K, V]) Len() int {
	var length int
	for _, bucket := range b.buckets {
		bucket.RLock()
		length += len(bucket.kv)
		bucket.RUnlock()
	}
	return length
}

func (b *MutexBucket[K, V]) Clear() {
	for _, bucket := range b.buckets {
		bucket.Lock()
		bucket.kv = make(map[K]V)
		bucket.Unlock()
	}
}

type MutexBucketItem[K constraints.Hash, V any] struct {
	sync.RWMutex
	kv map[K]V
}

func (i *MutexBucketItem[K, V]) Get(key K) (value V, exists bool) {
	i.RLock()
	v, exist := i.kv[key]
	i.RUnlock()
	return v, exist
}

func (i *MutexBucketItem[K, V]) GetOrSet(key K, val V) (value V, exists bool) {
	i.RLock()
	v, exist := i.kv[key]
	i.RUnlock()
	if exist {
		return v, true
	}
	i.Lock()
	defer i.Unlock()
	v, exist = i.kv[key]
	if exist {
		return v, true
	}
	i.kv[key] = val
	return val, false
}

func (i *MutexBucketItem[K, V]) GetAndDel(key K) (value V, exists bool) {
	i.Lock()
	defer i.Unlock()
	v, exist := i.kv[key]
	if exist {
		delete(i.kv, key)
	}
	return v, exist
}

func (i *MutexBucketItem[K, V]) NoneLockGetAndDel(key K) (value V, exists bool) {
	v, exist := i.kv[key]
	if exist {
		delete(i.kv, key)
	}
	return v, exist
}
