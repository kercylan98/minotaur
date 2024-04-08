package queue

type (
	incrementCustomMessageCountHandler func(delta int64)
)

func newController[Id, Q comparable, M Message[Q]](queue *Queue[Id, Q, M], message Message[Q]) Controller {
	return Controller{
		incrementCustomMessageCount: func(delta int64) {
			queueName := message.GetQueue()

			queue.cond.L.Lock()

			currIdent := queue.identifiers[queueName]
			currIdent += delta
			queue.identifiers[queueName] = currIdent
			queue.state.total += delta
			//log.Info("消息总计数", log.Int64("计数", q.state.total))

			queue.cond.Signal()
			queue.cond.L.Unlock()
		},
	}
}

// Controller 队列控制器
type Controller struct {
	incrementCustomMessageCount incrementCustomMessageCountHandler
}

// IncrementCustomMessageCount 增加自定义消息计数，当消息计数不为 > 0 时会导致队列关闭进入等待状态
func (c Controller) IncrementCustomMessageCount(delta int64) {
	c.incrementCustomMessageCount(delta)
}
