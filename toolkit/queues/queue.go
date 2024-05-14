package queues

// Queue 队列的接口，该接口用于定义一个队列
type Queue[T any] interface {
	// Start 开始队列的消费逻辑，此刻应保证队列已经允许写入
	Start()

	// Stop 停止队列，队列的停止方式根据不同的实现可能会有所不同
	Stop()

	// Enqueue 用于将一个元素放入队列
	Enqueue(elem T)

	// Dequeue 用于从队列中取出一个元素
	Dequeue() <-chan T
}
