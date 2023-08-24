package storage

import (
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
	"time"
)

// New 创建一个新的存储器
func New[K comparable, D any, W Warehouse[K, D]](warehouse W) *Storage[K, D, W] {
	s := &Storage[K, D, W]{
		w: warehouse,
	}

	if cache, err := warehouse.Init(); err != nil {
		panic(err)
	} else {
		s.cache = cache
	}

	if s.cache == nil {
		s.cache = make(map[K][]byte)
	}
	return s
}

// Storage 用于缓存数据的存储器
type Storage[K comparable, D any, W Warehouse[K, D]] struct {
	cl    sync.RWMutex
	sl    sync.Mutex
	w     W
	cache map[K][]byte
}

// Query 查询特定 key 的数据
func (slf *Storage[Key, D, W]) Query(key Key) (v D, err error) {
	slf.cl.RLock()
	cache, exist := slf.cache[key]
	slf.cl.RUnlock()

	if !exist {
		cache, err = slf.w.Query(key)
		if err != nil {
			panic(err)
		}
	}

	v = slf.w.GenerateZero()
	if err = super.UnmarshalJSON(cache, v); err != nil {
		panic(err)
	}
	return v, err
}

// Create 创建特定 key 的数据
func (slf *Storage[K, D, W]) Create(key K, data D) error {
	d, err := super.MarshalJSONE(data)
	if err != nil {
		return err
	}
	return slf.w.Create(key, d)
}

// Save 保存特定 key 的数据
func (slf *Storage[K, D, W]) Save(key K, data D) error {
	d, err := super.MarshalJSONE(data)
	if err != nil {
		return err
	}
	slf.cl.Lock()
	slf.cache[key] = d
	slf.cl.Unlock()
	return nil
}

// Flush 将缓存中的数据全部保存到数据库中，如果保存失败，会调用 errHandle 处理错误，如果 errHandle 返回 >= 0 的时长，则继续尝试保存，否则将跳过本条数据
//   - 当 errHandle 为 nil 时，将会无限重试保存，间隔为 100ms
func (slf *Storage[K, D, W]) Flush(errHandle func(data []byte, err error) time.Duration) {
	slf.cl.Lock()
	if len(slf.cache) == 0 {
		slf.cl.Unlock()
		return
	}
	cache := slf.cache
	slf.cache = make(map[K][]byte)
	slf.cl.Unlock()

	slf.sl.Lock()
	defer slf.sl.Unlock()
	for key, data := range cache {
		for {
			if err := slf.w.Save(key, data); err != nil {
				if errHandle == nil {
					time.Sleep(time.Millisecond * 100)
					continue
				} else if d := errHandle(data, err); d >= 0 {
					time.Sleep(d)
					continue
				} else {
					break
				}
			} else {
				break
			}
		}
	}
}

// Migrate 迁移数据，如果 keys 为空，则迁移全部数据
func (slf *Storage[K, D, W]) Migrate(keys ...K) (data []byte, err error) {
	slf.cl.RLock()
	defer slf.cl.RUnlock()
	if len(keys) > 0 {
		var m = make(map[K][]byte)
		for _, key := range keys {
			d, exist := slf.cache[key]
			if !exist {
				d, err = slf.w.Query(key)
				if err != nil {
					return nil, ErrDataNotExist
				}
			}
			m[key] = d
		}
		return super.MarshalJSONE(m)
	}
	return super.MarshalJSONE(slf.cache)
}

// LoadMigrationData 加载迁移数据
func (slf *Storage[K, D, W]) LoadMigrationData(data []byte) error {
	m := make(map[K][]byte)
	if err := super.UnmarshalJSON(data, &m); err != nil {
		return err
	}
	slf.cl.Lock()
	defer slf.cl.Unlock()
	for key, d := range m {
		slf.cache[key] = d
	}
	return nil
}
