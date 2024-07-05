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
		registerTable: mappings.NewMutexBucket[string, Process](bucketSize, func(size int, key string) int {
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
	registerTable  *mappings.MutexBucket[string, Process] // key = core.PhysicalAddress.LogicPath()
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
	bucket := prt.registerTable.GetBucket(address.LogicPath())
	_, exists = bucket.GetOrSet(address.LogicPath(), process)
	return &ProcessRef{
		address: address,
	}, exists
}

// GetProcess 获取一个进程
func (prt *processRegisterTable) GetProcess(address Address) (Process, bool) {
	// 如果两个地址不为本地，那么尝试使用注册的解析器进行解析
	if !prt.Address().IsLocal(address) {
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

	bucket := prt.registerTable.GetBucket(address.LogicPath())
	process, exists := bucket.Get(address.LogicPath())
	if exists {
		return process, exists
	}

	return prt.defaultProcess, exists
}

// Unregister 注销一个进程
func (prt *processRegisterTable) Unregister(ref *ProcessRef, handler ...func()) {
	bucket := prt.registerTable.GetBucket(ref.address.LogicPath())
	bucket.Lock()
	process, exist := bucket.NoneLockGetAndDel(ref.address.LogicPath())
	for _, f := range handler {
		func() {
			defer func() {
				if err := recover(); err != nil {
					// ignore
				}
			}()
			f()
		}()
	}
	bucket.Unlock()
	if exist {
		if status, ok := process.(ProcessStatus); ok {
			status.Dead()
		}
	}
}
