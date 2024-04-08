package reactor

import "github.com/kercylan98/minotaur/server/internal/v2/queue"

type MessageHandler[Q comparable, M queue.Message[Q]] func(message M)

type ErrorHandler[Q comparable, M queue.Message[Q]] func(message M, err error)
