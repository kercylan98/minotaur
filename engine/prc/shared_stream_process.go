package prc

import (
	"github.com/kercylan98/minotaur/toolkit/buffer"
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
		batches: buffer.NewRing[*DeliveryMessage](),
	}
}

type sharedStreamProcess struct {
	stream  sharedStream
	shared  *Shared
	address PhysicalAddress
	batches *buffer.Ring[*DeliveryMessage]
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
	c.batches.Write(dm)
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
				empty := c.batches.IsEmpty()
				c.lock.RUnlock()
				if empty {
					break
				} else if !c.state.CompareAndSwap(sharedStreamProcessStateIdle, sharedStreamProcessStateActive) {
					break
				}
			}
			c.send()
		}()
	}
}

func (c *sharedStreamProcess) send() {
	for {
		c.lock.Lock()
		messages := c.batches.ReadAll()
		c.lock.Unlock()
		if len(messages) == 0 {
			return
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
			c.batches.Reset()
			return
		}
	}
}
