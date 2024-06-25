package core

import (
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/twmb/murmur3"
)

type AddressResolver func(address Address) Process

// newProcessRegisterTable 创建一个进程注册表
func newProcessRegisterTable(address Address, bucketSize int, defaultProcess ...Process) *processRegisterTable {
	t := &processRegisterTable{
		address: NewAddress(address.Network(), address.System(), address.Host(), address.Port(), ""),
		registerTable: mappings.NewMutexBucket[Address, Process](bucketSize, func(size int, key Address) int {
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
	address        Address
	registerTable  *mappings.MutexBucket[Address, Process]
	resolver       []AddressResolver
	defaultProcess Process
	onlyAddress    bool
}

// SetAddressResolverOnlyAddress 设置地址解析器不包含额外的 path 信息
func (prt *processRegisterTable) SetAddressResolverOnlyAddress() {
	prt.onlyAddress = true
}

// Address 获取地址
func (prt *processRegisterTable) Address() Address {
	return prt.address
}

// RegisterAddressResolver 注册一个地址解析器
func (prt *processRegisterTable) RegisterAddressResolver(resolver AddressResolver) {
	prt.resolver = append(prt.resolver, resolver)
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
	if address.Address() != prt.Address().Address() {
		for _, resolver := range prt.resolver {
			var process Process
			if prt.onlyAddress {
				process = resolver(address.ParseToRoot())
			} else {
				process = resolver(address)
			}
			if process != nil {
				return process, true
			}
		}
		return prt.defaultProcess, false
	}

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
		if status, ok := process.(ProcessStatus); ok {
			status.Dead()
		}
	}
}
