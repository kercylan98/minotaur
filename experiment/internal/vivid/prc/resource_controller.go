package prc

import (
	"github.com/kercylan98/minotaur/toolkit/network"
	"github.com/puzpuzpuz/xsync"
)

// NewResourceController 创建一个新的资源控制器
func NewResourceController(pa PhysicalAddress) *ResourceController {
	return &ResourceController{
		pa:        pa,
		processes: xsync.NewMapOf[Process](),
	}

}

// ResourceController 是一个支持分布式、集群架构的资源控制器，它将所有资源视为进程(Process)，进程之间通过 ProcessRef 进行通信。
type ResourceController struct {
	cn        ClusterName                           // 资源控制器的集群名称，即便是相同物理地址，但是集群名称不同的话，也应该视为非本地
	pa        PhysicalAddress                       // 资源控制器的物理地址，如：:8080
	par       []PhysicalAddressResolver             // 线程不安全的物理地址解析器列表，将遍历找到首个有效解析器（应在初始化期间便注册完毕）
	processes *xsync.MapOf[LogicalAddress, Process] // 用于存储所有进程的映射表
}

// GetPhysicalAddress 获取资源控制器的物理地址
func (rc *ResourceController) GetPhysicalAddress() PhysicalAddress {
	return rc.pa
}

// Register 向资源控制器注册一个进程，如果进程已存在，将会返回已有的 ProcessRef 和一个标识是否已存在的状态信息，这对于进程的重复注册检测是非常有用的
func (rc *ResourceController) Register(id *ProcessId, process Process) (ref *ProcessRef, exist bool) {
	process, exist = rc.processes.LoadOrStore(id.LogicalAddress, process)
	if !exist {
		process.Initialize(rc, id)
	}
	ref = NewProcessRef(id)
	return
}

// Unregister 从资源控制器注销一个进程
func (rc *ResourceController) Unregister(killer *ProcessRef, ref *ProcessRef) {
	process, exist := rc.processes.LoadAndDelete(ref.id.LogicalAddress)
	if !exist {
		return
	}
	process.Terminate(killer)
}

// GetProcess 获取一个进程
func (rc *ResourceController) GetProcess(ref *ProcessRef) (process Process) {
	processPtr := ref.cache.Load()
	if processPtr != nil {
		process = *processPtr
		if !process.IsTerminated() {
			return process
		}

		ref.cache.Store(nil)
	}

	if rc.cn != ref.id.ClusterName || !network.IsSameLocalAddress(rc.pa, ref.id.PhysicalAddress) {
		// 远程进程加载
		for _, resolver := range rc.par {
			if process = resolver.Resolve(ref.id); process != nil {
				ref.cache.Store(&process)
				return
			}
		}
	}

	// 本地进程加载
	var exist bool
	process, exist = rc.processes.Load(ref.id.LogicalAddress)
	if exist {
		ref.cache.Store(&process)
	}
	return process // 如果进程不存在，返回的是一个 Process 的空指针
}
