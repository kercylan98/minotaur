package dispatcher

// Action 消息分发器操作器，用于暴露外部可操作的消息分发器函数
type Action[P Producer, M Message[P]] struct {
	unlock bool
	d      *Dispatcher[P, M]
}

// Name 获取消息分发器名称
func (a *Action[P, M]) Name() string {
	return a.d.Name()
}

// UnExpel 取消特定生产者的驱逐计划
func (a *Action[P, M]) UnExpel() {
	if !a.unlock {
		a.d.UnExpel()
	} else {
		a.d.noLockUnExpel()
	}
}

// Expel 设置该消息分发器即将被驱逐，当消息分发器中没有任何消息时，会自动关闭
func (a *Action[P, M]) Expel() {
	if !a.unlock {
		a.d.Expel()
	} else {
		a.d.noLockExpel()
	}
}
