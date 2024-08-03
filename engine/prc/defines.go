package prc

type (
	// PhysicalAddress 物理地址是用于标识内容的网络地址
	PhysicalAddress = string
	// LogicalAddress 逻辑地址是用于标识内容的本地内部地址
	LogicalAddress = string
	// Message 消息是用于传递的数据
	Message = any
)

const (
	LocalhostPhysicalAddress = "localhost" // 无网络本地
)
