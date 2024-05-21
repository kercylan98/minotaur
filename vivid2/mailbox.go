package vivid

const DefaultMailboxFactoryId MailboxFactoryId = 1

type MailboxFactoryId = uint64

type MailboxFactory interface {
	Get() Mailbox
	Put(Mailbox)
}

type Mailbox interface {
	// Start 开始队列的消费逻辑，此刻应保证队列已经允许写入
	Start()

	// Stop 停止队列，队列的停止方式根据不同的实现可能会有所不同
	Stop()

	// Enqueue 将一个消息放入队列
	Enqueue(message MessageContext)
}
