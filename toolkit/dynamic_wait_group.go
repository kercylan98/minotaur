package toolkit

import (
	"fmt"
	"sync"
)

// NewDynamicWaitGroup 创建一个新的 DynamicWaitGroup
func NewDynamicWaitGroup() *DynamicWaitGroup {
	return &DynamicWaitGroup{
		wait: sync.NewCond(new(sync.Mutex)),
	}
}

// DynamicWaitGroup 是一个动态的 WaitGroup，允许在等待的过程中动态地添加或减少等待的计数
type DynamicWaitGroup struct {
	c          int64
	wait       *sync.Cond
	ChangeHook func(curr int64, delta int64)
}

// Add 增加等待的计数
func (d *DynamicWaitGroup) Add(delta int64) {
	d.wait.L.Lock()
	defer d.wait.L.Unlock()
	d.c += delta
	if d.ChangeHook != nil {
		d.ChangeHook(d.c, delta)
	}
	if d.c < 0 {
		panic(fmt.Errorf("negative DynamicWaitGroup counter: %d", d.c))
	}
	if d.c == 0 { // 如果计数变为0，唤醒所有等待的 goroutine
		d.wait.Broadcast()
	}
}

// DoneAll 减少等待的计数到0
func (d *DynamicWaitGroup) DoneAll() {
	d.wait.L.Lock()
	defer d.wait.L.Unlock()
	d.c = 0
	d.wait.Broadcast()
}

// Done 减少等待的计数
func (d *DynamicWaitGroup) Done() {
	d.Add(-1)
}

// Wait 等待所有的计数完成
//   - 当传入 handler 时将会在计数完成后执行 handler，执行时会阻止计数器的变化
func (d *DynamicWaitGroup) Wait(handler ...func()) {
	d.wait.L.Lock()
	defer d.wait.L.Unlock()

	for d.c != 0 {
		d.wait.Wait()
	}

	for _, f := range handler {
		f()
	}
}
