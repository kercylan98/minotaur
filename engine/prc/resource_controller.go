package prc

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"github.com/puzpuzpuz/xsync/v3"
)

// NewResourceController 创建一个新的资源控制器
func NewResourceController(configurator ...ResourceControllerConfigurator) *ResourceController {
	rc := &ResourceController{
		config:    newResourceControllerConfiguration(),
		processes: xsync.NewMapOf[LogicalAddress, Process](),
	}
	for _, c := range configurator {
		c.Configure(rc.config)
	}
	return rc
}

// ResourceController 是一个支持分布式、集群架构的资源控制器，它将所有资源视为进程(Process)，进程之间通过 ProcessRef 进行通信。
type ResourceController struct {
	config    *ResourceControllerConfiguration
	par       []PhysicalAddressResolver             // 线程不安全的物理地址解析器列表，将遍历找到首个有效解析器（应在初始化期间便注册完毕）
	processes *xsync.MapOf[LogicalAddress, Process] // 用于存储所有进程的映射表
}

// logger 日志记录器
func (rc *ResourceController) logger() *log.Logger {
	return rc.config.loggerProvider.Provide()
}

// RegisterResolver 注册用于物理地址解析的解析器，解析器应返回一个可用进程。
//   - 解析器需要依赖于外部的进程管理，本身不会涉及进程的注册与反注册
func (rc *ResourceController) RegisterResolver(resolver ...PhysicalAddressResolver) {
	rc.par = append(rc.par, resolver...)
}

// GetPhysicalAddress 获取资源控制器的物理地址
func (rc *ResourceController) GetPhysicalAddress() PhysicalAddress {
	return rc.config.physicalAddress
}

// Register 向资源控制器注册一个进程，如果进程已存在，将会返回已有的 ProcessRef 和一个标识是否已存在的状态信息，这对于进程的重复注册检测是非常有用的
func (rc *ResourceController) Register(id *ProcessId, process Process) (ref *ProcessRef, exist bool) {
	process, exist = rc.processes.LoadOrStore(id.LogicalAddress, process)
	if !exist {
		process.Initialize(rc, id)
		rc.logger().Debug("ResourceController", log.String("register", id.URL().String()))
	}
	ref = NewProcessRef(id)
	return
}

// Unregister 从资源控制器注销一个进程
func (rc *ResourceController) Unregister(killer *ProcessRef, ref *ProcessRef) {
	process, exist := rc.processes.LoadAndDelete(ref.GetId().LogicalAddress)
	if !exist {
		return
	}
	process.Terminate(killer)
	rc.logger().Debug("ResourceController", log.String("unregister", ref.URL().String()))
}

// Belong 检查 ref 是否属于该资源控制器。该函数并不检查进程是否存在，只检查进程的归属关系。
func (rc *ResourceController) Belong(ref *ProcessRef) bool {
	return network.IsSameLocalAddress(rc.config.physicalAddress, ref.GetId().PhysicalAddress)
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

	if !rc.Belong(ref) {
		// 远程进程加载
		for _, resolver := range rc.par {
			if process = resolver.Resolve(ref.GetId()); process != nil {
				ref.cache.Store(&process)
				return
			}
		}
	}

	// 本地进程加载
	var exist bool
	process, exist = rc.processes.Load(ref.GetId().LogicalAddress)
	if exist {
		ref.cache.Store(&process)
		return process
	} else {
		// 找不到进程时返回默认的替代进程，当默认的替代进程也不存在那么将是空指针
		return rc.config.notFoundSubstitute
	}
}
