package queue

import (
	"errors"
)

var (
	ErrorQueueClosed  = errors.New("queue closed")  // 队列已关闭
	ErrorQueueInvalid = errors.New("queue invalid") // 无效的队列
)
