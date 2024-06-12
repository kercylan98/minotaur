package vivid

import (
	"time"
)

type subscribeInfo struct {
	subscribeId     SubscribeId    // 订阅 ID
	subscriber      Subscriber     // 事件订阅者
	priority        *int64         // 优先级，值越小优先级越高
	priorityTimeout *time.Duration // 优先级事件超时时间
	producer        *Producer      // 事件生产者
}
