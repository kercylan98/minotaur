package core

import (
	"sync/atomic"
)

// Process 进程是一个抽象概念，它包含了消息发送
type Process interface {
	// GetAddress 获取进程地址
	GetAddress() Address

	// SendUserMessage 发送用户消息
	SendUserMessage(sender *ProcessRef, message Message)

	// SendSystemMessage 发送系统消息
	SendSystemMessage(sender *ProcessRef, message Message)

	// Terminate 终止进程
	Terminate(*ProcessRef)
}

type ProcessStatus interface {
	Process

	// Deaden 判断进程是否已经死亡
	Deaden() bool

	// Dead 设置进程死亡
	Dead()
}

// NewProcessRef 创建一个进程引用
func NewProcessRef(address Address) *ProcessRef {
	return &ProcessRef{
		address: address,
	}
}

// ProcessRef 进程外部引用
type ProcessRef struct {
	address Address
	cache   atomic.Pointer[Process]
}

// Address 获取进程地址
func (r *ProcessRef) Address() Address {
	return r.address
}
