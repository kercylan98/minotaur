package super

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"sync"
)

// NewPermission 创建权限
func NewPermission[Code generic.Integer, EntityID comparable]() *Permission[Code, EntityID] {
	return &Permission[Code, EntityID]{
		permissions: make(map[EntityID]Code),
	}
}

type Permission[Code generic.Integer, EntityID comparable] struct {
	permissions map[EntityID]Code
	l           sync.RWMutex
}

// HasPermission 是否有权限
func (slf *Permission[Code, EntityID]) HasPermission(entityId EntityID, permission Code) bool {
	slf.l.RLock()
	c, exist := slf.permissions[entityId]
	slf.l.RUnlock()
	if !exist {
		return false
	}

	return c&permission != 0
}

// AddPermission 添加权限
func (slf *Permission[Code, EntityID]) AddPermission(entityId EntityID, permission ...Code) {

	slf.l.Lock()
	defer slf.l.Unlock()

	userPermission, exist := slf.permissions[entityId]
	if !exist {
		userPermission = 0
		slf.permissions[entityId] = userPermission
	}

	for _, p := range permission {
		userPermission |= p
	}

	slf.permissions[entityId] = userPermission
}

// RemovePermission 移除权限
func (slf *Permission[Code, EntityID]) RemovePermission(entityId EntityID, permission ...Code) {
	slf.l.Lock()
	defer slf.l.Unlock()

	userPermission, exist := slf.permissions[entityId]
	if !exist {
		return
	}

	for _, p := range permission {
		userPermission &= ^p
	}

	slf.permissions[entityId] = userPermission
}

// SetPermission 设置权限
func (slf *Permission[Code, EntityID]) SetPermission(entityId EntityID, permission ...Code) {
	slf.l.Lock()
	defer slf.l.Unlock()

	var userPermission Code
	for _, p := range permission {
		userPermission |= p
	}
	slf.permissions[entityId] = userPermission
}
