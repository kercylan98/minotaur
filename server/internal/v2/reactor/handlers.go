package reactor

import "github.com/kercylan98/minotaur/server/internal/v2/queue"

type MessageHandler[M any] func(message queue.MessageWrapper[int, string, M])

type ErrorHandler[M any] func(message queue.MessageWrapper[int, string, M], err error)
