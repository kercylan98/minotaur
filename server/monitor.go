package server

import (
	"github.com/kercylan98/minotaur/utils/timer"
	"sync"
	"time"
)

func newMonitor() *monitor {
	m := &monitor{
		ticker: timer.GetTicker(10),
	}
	m.ticker.Loop("tick", timer.Instantly, time.Second, timer.Forever, m.tick)
	return m
}

type Monitor interface {
	MessageTotal() int64
	PacketMessageTotal() int64
	ErrorMessageTotal() int64
	CrossMessageTotal() int64
	TickerMessageTotal() int64
	MessageSecond() int64
	PacketMessageSecond() int64
	ErrorMessageSecond() int64
	CrossMessageSecond() int64
	TickerMessageSecond() int64
	MessageCost() time.Duration
	PacketMessageCost() time.Duration
	ErrorMessageCost() time.Duration
	CrossMessageCost() time.Duration
	TickerMessageCost() time.Duration
	MessageDoneAvg() time.Duration
	PacketMessageDoneAvg() time.Duration
	ErrorMessageDoneAvg() time.Duration
	CrossMessageDoneAvg() time.Duration
	TickerMessageDoneAvg() time.Duration
	MessageQPS() int64
	PacketMessageQPS() int64
	ErrorMessageQPS() int64
	CrossMessageQPS() int64
	TickerMessageQPS() int64
	MessageTopQPS() int64
	PacketMessageTopQPS() int64
	ErrorMessageTopQPS() int64
	CrossMessageTopQPS() int64
	TickerMessageTopQPS() int64
}

type monitor struct {
	rwMutex sync.RWMutex
	ticker  *timer.Ticker

	messageTotal         int64         // 正在执行的消息总数
	packetMessageTotal   int64         // 正在执行的玩家消息总数
	errorMessageTotal    int64         // 正在执行的错误消息总数
	crossMessageTotal    int64         // 正在执行的跨服消息总数
	tickerMessageTotal   int64         // 正在执行的定时器消息总数
	messageSecond        int64         // 一秒内执行的消息并发量
	packetMessageSecond  int64         // 一秒内玩家消息并发量
	errorMessageSecond   int64         // 一秒内错误消息并发量
	crossMessageSecond   int64         // 一秒内跨服消息并发量
	tickerMessageSecond  int64         // 一秒内定时器消息并发量
	messageCost          time.Duration // 一秒内执行的消息总耗时
	packetMessageCost    time.Duration // 一秒内玩家消息总耗时
	errorMessageCost     time.Duration // 一秒内错误消息总耗时
	crossMessageCost     time.Duration // 一秒内跨服消息总耗时
	tickerMessageCost    time.Duration // 一秒内定时器消息总耗时
	messageDoneAvg       time.Duration // 一秒内执行的消息平均响应时间
	packetMessageDoneAvg time.Duration // 一秒内玩家消息平均响应时间
	errorMessageDoneAvg  time.Duration // 一秒内错误消息平均响应时间
	crossMessageDoneAvg  time.Duration // 一秒内跨服消息平均响应时间
	tickerMessageDoneAvg time.Duration // 一秒内定时器消息平均响应时间
	messageQPS           int64         // 一秒内执行的消息QPS
	packetMessageQPS     int64         // 一秒内玩家消息QPS
	errorMessageQPS      int64         // 一秒内错误消息QPS
	crossMessageQPS      int64         // 一秒内跨服消息QPS
	tickerMessageQPS     int64         // 一秒内定时器消息QPS
	messageTopQPS        int64         // 执行的消息最高QPS
	packetMessageTopQPS  int64         // 玩家消息最高QPS
	errorMessageTopQPS   int64         // 错误消息最高QPS
	crossMessageTopQPS   int64         // 跨服消息最高QPS
	tickerMessageTopQPS  int64         // 定时器消息最高QPS
}

func (slf *monitor) tick() {
	slf.rwMutex.Lock()

	// 秒平均响应时间
	if slf.messageSecond == 0 {
		slf.messageDoneAvg = 0
	} else {
		slf.messageDoneAvg = time.Duration(slf.messageCost.Nanoseconds() / slf.messageSecond)
	}
	if slf.packetMessageSecond == 0 {
		slf.packetMessageDoneAvg = 0
	} else {
		slf.packetMessageDoneAvg = time.Duration(slf.packetMessageCost.Nanoseconds() / slf.packetMessageSecond)
	}
	if slf.errorMessageSecond == 0 {
		slf.errorMessageDoneAvg = 0
	} else {
		slf.errorMessageDoneAvg = time.Duration(slf.errorMessageCost.Nanoseconds() / slf.errorMessageSecond)
	}
	if slf.crossMessageSecond == 0 {
		slf.crossMessageDoneAvg = 0
	} else {
		slf.crossMessageDoneAvg = time.Duration(slf.crossMessageCost.Nanoseconds() / slf.crossMessageSecond)
	}
	if slf.tickerMessageSecond == 0 {
		slf.tickerMessageDoneAvg = 0
	} else {
		slf.tickerMessageDoneAvg = time.Duration(slf.tickerMessageCost.Nanoseconds() / slf.tickerMessageSecond)
	}

	// 秒 QPS
	if nanoseconds := slf.messageDoneAvg.Nanoseconds(); nanoseconds == 0 {
		slf.messageQPS = 0
	} else {
		slf.messageQPS = slf.messageSecond / nanoseconds
	}
	if nanoseconds := slf.packetMessageDoneAvg.Nanoseconds(); nanoseconds == 0 {
		slf.packetMessageQPS = 0
	} else {
		slf.packetMessageQPS = slf.packetMessageSecond / nanoseconds
	}
	if nanoseconds := slf.errorMessageDoneAvg.Nanoseconds(); nanoseconds == 0 {
		slf.errorMessageQPS = 0
	} else {
		slf.errorMessageQPS = slf.errorMessageSecond / nanoseconds
	}
	if nanoseconds := slf.crossMessageDoneAvg.Nanoseconds(); nanoseconds == 0 {
		slf.crossMessageQPS = 0
	} else {
		slf.crossMessageQPS = slf.crossMessageSecond / nanoseconds
	}
	if nanoseconds := slf.tickerMessageDoneAvg.Nanoseconds(); nanoseconds == 0 {
		slf.tickerMessageQPS = 0
	} else {
		slf.tickerMessageQPS = slf.tickerMessageSecond / nanoseconds
	}

	// Top QPS
	if slf.messageQPS > slf.messageTopQPS {
		slf.messageTopQPS = slf.messageQPS
	}
	if slf.packetMessageQPS > slf.packetMessageTopQPS {
		slf.packetMessageTopQPS = slf.packetMessageQPS
	}
	if slf.errorMessageQPS > slf.errorMessageTopQPS {
		slf.errorMessageTopQPS = slf.errorMessageQPS
	}
	if slf.crossMessageQPS > slf.crossMessageTopQPS {
		slf.crossMessageTopQPS = slf.crossMessageQPS
	}
	if slf.tickerMessageQPS > slf.tickerMessageTopQPS {
		slf.tickerMessageTopQPS = slf.tickerMessageQPS
	}

	slf.messageSecond = 0
	slf.packetMessageSecond = 0
	slf.errorMessageSecond = 0
	slf.crossMessageSecond = 0
	slf.tickerMessageSecond = 0
	slf.messageCost = 0
	slf.packetMessageCost = 0
	slf.errorMessageCost = 0
	slf.crossMessageCost = 0
	slf.tickerMessageCost = 0
	slf.rwMutex.Unlock()
}

func (slf *monitor) messageRun(msg *Message) {
	slf.rwMutex.Lock()
	defer slf.rwMutex.Unlock()
	switch msg.t {
	case MessageTypePacket:
		slf.packetMessageTotal++
		slf.packetMessageSecond++
	case MessageTypeError:
		slf.errorMessageTotal++
		slf.errorMessageSecond++
	case MessageTypeCross:
		slf.crossMessageTotal++
		slf.crossMessageSecond++
	case MessageTypeTicker:
		slf.tickerMessageTotal++
		slf.tickerMessageSecond++
	default:
		return
	}
	slf.messageTotal++
	slf.messageSecond++
}

func (slf *monitor) messageDone(msg *Message, cost time.Duration) {
	slf.rwMutex.Lock()
	defer slf.rwMutex.Unlock()
	switch msg.t {
	case MessageTypePacket:
		slf.packetMessageTotal--
		slf.packetMessageCost += cost
	case MessageTypeError:
		slf.errorMessageTotal--
		slf.errorMessageCost += cost
	case MessageTypeCross:
		slf.crossMessageTotal--
		slf.crossMessageCost += cost
	case MessageTypeTicker:
		slf.tickerMessageTotal--
		slf.tickerMessageCost += cost
	default:
		return
	}
	slf.messageTotal--
	slf.messageCost += cost
}

func (slf *monitor) close() {
	slf.ticker.Release()
}

func (slf *monitor) MessageTotal() int64 {
	return slf.messageTotal
}

func (slf *monitor) PacketMessageTotal() int64 {
	return slf.packetMessageTotal
}

func (slf *monitor) ErrorMessageTotal() int64 {
	return slf.errorMessageTotal
}

func (slf *monitor) CrossMessageTotal() int64 {
	return slf.crossMessageTotal
}

func (slf *monitor) TickerMessageTotal() int64 {
	return slf.tickerMessageTotal
}

func (slf *monitor) MessageSecond() int64 {
	return slf.messageSecond
}

func (slf *monitor) PacketMessageSecond() int64 {
	return slf.packetMessageSecond
}

func (slf *monitor) ErrorMessageSecond() int64 {
	return slf.errorMessageSecond
}

func (slf *monitor) CrossMessageSecond() int64 {
	return slf.crossMessageSecond
}

func (slf *monitor) TickerMessageSecond() int64 {
	return slf.tickerMessageSecond
}

func (slf *monitor) MessageCost() time.Duration {
	return slf.messageCost
}

func (slf *monitor) PacketMessageCost() time.Duration {
	return slf.packetMessageCost
}

func (slf *monitor) ErrorMessageCost() time.Duration {
	return slf.errorMessageCost
}

func (slf *monitor) CrossMessageCost() time.Duration {
	return slf.crossMessageCost
}

func (slf *monitor) TickerMessageCost() time.Duration {
	return slf.tickerMessageCost
}

func (slf *monitor) MessageDoneAvg() time.Duration {
	return slf.messageDoneAvg
}

func (slf *monitor) PacketMessageDoneAvg() time.Duration {
	return slf.packetMessageDoneAvg
}

func (slf *monitor) ErrorMessageDoneAvg() time.Duration {
	return slf.errorMessageDoneAvg
}

func (slf *monitor) CrossMessageDoneAvg() time.Duration {
	return slf.crossMessageDoneAvg
}

func (slf *monitor) TickerMessageDoneAvg() time.Duration {
	return slf.tickerMessageDoneAvg
}

func (slf *monitor) MessageQPS() int64 {
	return slf.messageQPS
}

func (slf *monitor) PacketMessageQPS() int64 {
	return slf.packetMessageQPS
}

func (slf *monitor) ErrorMessageQPS() int64 {
	return slf.errorMessageQPS
}

func (slf *monitor) CrossMessageQPS() int64 {
	return slf.crossMessageQPS
}

func (slf *monitor) TickerMessageQPS() int64 {
	return slf.tickerMessageQPS
}

func (slf *monitor) MessageTopQPS() int64 {
	return slf.messageTopQPS
}

func (slf *monitor) PacketMessageTopQPS() int64 {
	return slf.packetMessageTopQPS
}

func (slf *monitor) ErrorMessageTopQPS() int64 {
	return slf.errorMessageTopQPS
}

func (slf *monitor) CrossMessageTopQPS() int64 {
	return slf.crossMessageTopQPS
}

func (slf *monitor) TickerMessageTopQPS() int64 {
	return slf.tickerMessageTopQPS
}
