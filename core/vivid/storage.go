package vivid

import (
	"github.com/alphadose/haxmap"
)

var defaultStorage = &MemoryStorage{
	m: haxmap.New[string, PersistenceStatus](),
}

type Storage interface {
	// Persist 将指定名称的持久化状态存储到存储器中
	Persist(name string, status PersistenceStatus)

	// Load 从存储器中加载指定名称的持久化状态
	Load(name string) PersistenceStatus
}

func NewMemoryStorage() *MemoryStorage {
	return defaultStorage
}

type MemoryStorage struct {
	m *haxmap.Map[string, PersistenceStatus]
}

func (s *MemoryStorage) Persist(name string, status PersistenceStatus) {
	s.m.Set(name, status)
}

func (s *MemoryStorage) Load(name string) PersistenceStatus {
	if v, ok := s.m.Get(name); ok {
		return v
	}
	return nil
}
