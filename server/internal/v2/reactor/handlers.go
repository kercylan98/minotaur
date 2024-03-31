package reactor

type queueMessageHandler[M any] func(q *queue[M], msg M)

type MessageHandler[M any] func(msg M)

type ErrorHandler[M any] func(msg M, err error)
