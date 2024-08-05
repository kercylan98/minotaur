package prc

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
	"sync/atomic"
)

const (
	sharedStreamProcessStateIdle uint32 = iota
	sharedStreamProcessStateActive
)

func newSharedStreamProcess(address PhysicalAddress, stream sharedStream, shared *Shared) *sharedStreamProcess {
	return &sharedStreamProcess{
		stream:  stream,
		shared:  shared,
		address: address,
	}
}

type sharedStreamProcess struct {
	stream  sharedStream
	shared  *Shared
	address PhysicalAddress
	batches []*DeliveryMessage
	lock    sync.RWMutex
	state   atomic.Uint32
}

func (c *sharedStreamProcess) Initialize(rc *ResourceController, id *ProcessId) {
	// 该进程不注册，不会触发
}

func (c *sharedStreamProcess) DeliveryUserMessage(receiver, sender, forward *ProcessId, message Message) {
	c.packMessage(receiver, sender, forward, message, false)
}

func (c *sharedStreamProcess) DeliverySystemMessage(receiver, sender, forward *ProcessId, message Message) {
	c.packMessage(receiver, sender, forward, message, true)
}

func (c *sharedStreamProcess) packMessage(receiver, sender, forward *ProcessId, message Message, system bool) {
	if err, ok := message.(error); ok {
		message = &SharedErrorMessage{Message: err.Error()}
	}

	name, data, err := c.shared.config.codec.Encode(message)
	if err != nil {
		panic(err)
	}
	dm := &DeliveryMessage{
		MessageType: name,
		MessageData: data,
		System:      system,
		Sender:      sender,
		Receiver:    receiver,
	}

	// 入列
	c.lock.Lock()
	c.batches = append(c.batches, dm)
	c.lock.Unlock()

	c.activation()
}

func (c *sharedStreamProcess) IsTerminated() bool {
	return false
}

func (c *sharedStreamProcess) Terminate(source *ProcessId) {
	// 该进程不注册，不会由资源控制器触发
}

func (c *sharedStreamProcess) activation() {
	if c.state.CompareAndSwap(sharedStreamProcessStateIdle, sharedStreamProcessStateActive) {
		go func() {
			for {
				c.send()
				c.state.Store(sharedStreamProcessStateIdle)
				c.lock.RLock()
				empty := len(c.batches) == 0
				c.lock.RUnlock()
				if empty {
					break
				} else if !c.state.CompareAndSwap(sharedStreamProcessStateIdle, sharedStreamProcessStateActive) {
					break
				}
			}
		}()
	}
}

func (c *sharedStreamProcess) send() {
	for {
		c.lock.Lock()
		n := len(c.batches)
		var messages []*DeliveryMessage
		if n < 1024 {
			messages = c.batches
			c.batches = nil
		} else {
			messages = c.batches[:1024]
			c.batches = c.batches[1024:]
		}
		c.lock.Unlock()
		if len(messages) == 0 {
			break
		}
		var sm *SharedMessage
		if len(messages) == 1 {
			sm = &SharedMessage{
				MessageType: &SharedMessage_DeliveryMessage{
					DeliveryMessage: messages[0],
				},
			}
		} else {
			sm = &SharedMessage{
				MessageType: &SharedMessage_BatchDeliveryMessage{
					BatchDeliveryMessage: &BatchDeliveryMessage{Messages: messages},
				},
			}
		}

		if err := c.stream.Send(sm); err != nil {
			c.shared.detachStream(c.address)
			c.shared.rc.logger().Error("ResourceController", log.Err(err))
			c.lock.Lock()
			c.batches = nil
			c.lock.Unlock()
			break
		}
	}
}
