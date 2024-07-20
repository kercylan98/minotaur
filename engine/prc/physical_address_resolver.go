package prc

// PhysicalAddressResolver 物理地址解析器是用来解析非本地地址的接口，它应返回一个与远程建立连接并能够与之交互的进程。
type PhysicalAddressResolver interface {
	Resolve(id *ProcessId) Process
}

// FunctionalPhysicalAddressResolver 是一个函数类型的物理地址解析器，它可以通过一个函数来实现 Resolve 方法。
type FunctionalPhysicalAddressResolver func(id *ProcessId) Process

func (f FunctionalPhysicalAddressResolver) Resolve(id *ProcessId) Process {
	return f(id)
}
