package core

import (
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/twmb/murmur3"
)

// newProcessRegisterTable 创建一个进程注册表
func newProcessRegisterTable(bucketSize int, defaultProcess ...Process) *processRegisterTable {
	t := &processRegisterTable{
		registerTable: mappings.NewBucket[Address, Process](bucketSize, func(size int, key Address) int {
			hash := murmur3.Sum32([]byte(key))
			index := int(hash) % size
			return index
		}),
	}
	if len(defaultProcess) > 0 {
		t.defaultProcess = defaultProcess[0]
	}
	return t
}

// processRegisterTable 进程注册表
type processRegisterTable struct {
	registerTable  *mappings.Bucket[Address, Process]
	resolver       any // 远程地址解析
	defaultProcess Process
}

// Register 注册一个进程
func (prt *processRegisterTable) Register(process Process) (ref *ProcessRef, exists bool) {
	address := process.GetAddress()
	bucket := prt.registerTable.GetBucket(address)
	_, exists = bucket.GetOrSet(address, process)
	return &ProcessRef{
		address: address,
	}, exists
}

// GetProcess 获取一个进程
func (prt *processRegisterTable) GetProcess(address Address) (Process, bool) {
	bucket := prt.registerTable.GetBucket(address)
	process, exists := bucket.Get(address)
	if exists {
		return process, exists
	}

	return prt.defaultProcess, exists
}

// Unregister 注销一个进程
func (prt *processRegisterTable) Unregister(ref *ProcessRef) {
	bucket := prt.registerTable.GetBucket(ref.address)
	process, exist := bucket.GetAndDel(ref.address)
	if exist {
		process.Dead()
	}
}
