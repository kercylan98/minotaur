package super

import (
	"runtime/debug"
)

// NewStackGo 返回一个用于获取上一个协程调用的堆栈信息的收集器
func NewStackGo() *StackGo {
	return new(StackGo)
}

// StackGo 用于获取上一个协程调用的堆栈信息
//   - 应当最先运行 Wait 函数，然后在其他协程中调用 Stack 函数或者 GiveUp 函数
//   - 适用于跨协程同步通讯，例如单线程的消息处理统计耗时打印堆栈信息
type StackGo struct {
	stack   chan *struct{} // 消息堆栈
	collect chan []byte    // 消息堆栈收集
}

// Wait 等待收集消息堆栈
//   - 在调用 Wait 函数后，当前协程将会被挂起，直到调用 Stack 或 GiveUp 函数
func (slf *StackGo) Wait() {
	slf.stack = make(chan *struct{}, 0)
	if s := <-slf.stack; s != nil {
		slf.collect <- debug.Stack()
	}
	close(slf.stack)
	slf.stack = nil
}

// Stack 获取消息堆栈
//   - 在调用 Wait 函数后调用该函数，将会返回上一个协程的堆栈信息
//   - 在调用 GiveUp 函数后调用该函数，将会 panic
func (slf *StackGo) Stack() []byte {
	if slf.stack == nil {
		return nil
	}
	slf.collect = make(chan []byte, 1)
	slf.stack <- &struct{}{}
	stack := <-slf.collect
	close(slf.collect)
	slf.collect = nil
	return stack
}

// GiveUp 放弃收集消息堆栈
//   - 在调用 Wait 函数后调用该函数，将会放弃收集消息堆栈并且释放资源
//   - 在调用 GiveUp 函数后调用 Stack 函数，将会 panic
func (slf *StackGo) GiveUp() {
	if slf.stack != nil {
		slf.stack <- nil
	}
}
