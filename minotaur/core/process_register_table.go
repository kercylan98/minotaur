package core

import (
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/twmb/murmur3"
)

// newProcessRegisterTable 创建一个进程注册表
func newProcessRegisterTable(bucketSize int) *processRegisterTable {
	return &processRegisterTable{
		registerTable: mappings.NewBucket[Address, Process](bucketSize, func(size int, key Address) int {
			hash := murmur3.Sum32([]byte(key))
			index := int(hash) % size
			return index
		}),
	}
}

// processRegisterTable 进程注册表
type processRegisterTable struct {
	registerTable *mappings.Bucket[Address, Process]
	resolver      any // 远程地址解析
}

// Register 注册一个进程
func (prt *processRegisterTable) Register(process Process) (ref *ProcessRef, exists bool) {
	bucket := prt.registerTable.GetBucket(process.GetAddress())
	_, exists = bucket.GetOrSet(process.GetAddress(), process)
	return &ProcessRef{
		address: process.GetAddress(),
	}, exists
}

// GetProcess 获取一个进程
func (prt *processRegisterTable) GetProcess(address Address) (Process, bool) {
	bucket := prt.registerTable.GetBucket(address)
	return bucket.Get(address)
}

// Unregister 注销一个进程
func (prt *processRegisterTable) Unregister(ref *ProcessRef) {
	bucket := prt.registerTable.GetBucket(ref.address)
	process, exist := bucket.GetAndDel(ref.address)
	if exist {
		process.Dead()
	}
}
