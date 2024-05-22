package vivid

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// DeadLetterStream 是一个用于管理死信事件的流
type DeadLetterStream interface {
	// DeadLetter 用于向 DeadLetterStream 中发送一条死信
	DeadLetter(deadLetter DeadLetterEvent)

	// Watch 用于监听指定类型的死信事件
	Watch(eventType DeadLetterEventType) DeadLetterEvent
}

type _DeadLetterStream struct {
	seq    atomic.Uint64
	events map[DeadLetterEventType]*deadLetterEvents
	rw     sync.RWMutex
}

type deadLetterEvents struct {
	buf  *buffer.Ring[DeadLetterEvent]
	cond *sync.Cond
}

func (s *_DeadLetterStream) getEvents(typ DeadLetterEventType) *deadLetterEvents {
	_, exist := deadLetterEventTypeStrings[typ]
	if !exist {
		panic(errors.New("unknown dead letter event type"))
	}

	s.rw.Lock()
	defer s.rw.Unlock()
	if s.events == nil {
		s.events = make(map[DeadLetterEventType]*deadLetterEvents)
	}
	events, exist := s.events[typ]
	if !exist {
		events = &deadLetterEvents{
			buf:  buffer.NewRing[DeadLetterEvent](100),
			cond: sync.NewCond(new(sync.Mutex)),
		}
		s.events[typ] = events
	}
	return events

}

func (s *_DeadLetterStream) DeadLetter(deadLetter DeadLetterEvent) {
	events := s.getEvents(deadLetter.Type)

	deadLetter.Seq = s.seq.Add(1)
	deadLetter.Stack = debug.Stack()

	events.cond.L.Lock()
	events.buf.Write(deadLetter)
	events.cond.L.Unlock()
	events.cond.Broadcast()
}

func (s *_DeadLetterStream) Watch(eventType DeadLetterEventType) DeadLetterEvent {
	events := s.getEvents(eventType)

	events.cond.L.Lock()
	event, err := events.buf.Read()
	for err != nil {
		if !errors.Is(err, buffer.ErrBufferIsEmpty) {
			panic(err)
		}
		events.cond.Wait()
		event, err = events.buf.Read()
	}
	events.cond.L.Unlock()

	return event
}
