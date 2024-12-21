package prc

import (
	"sync"
	"sync/atomic"
)

const (
	sharedStreamProcessStateIdle uint32 = iota
	sharedStreamProcessStateActive

	sharedStreamBatchLimit = 1024
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
	batches [][]byte
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
		message = &sharedErrorMessage{Message: err.Error()}
	}

	var dm *deliveryMessage
	switch wrapper := message.(type) {
	case *MessageWrapper:
		name, data, err := c.shared.config.codec.Encode(wrapper.Message)
		if err != nil {
			panic(err)
		}
		dm = wrapDeliveryMessage(name, data, system, wrapper.Sender, wrapper.Receiver)
	default:
		name, data, err := c.shared.config.codec.Encode(message)
		if err != nil {
			panic(err)
		}
		dm = wrapDeliveryMessage(name, data, system, sender, receiver)
	}

	// 持久化网络消息，避免消息丢失
	_, data, err := c.shared.config.codec.Encode(dm)
	if err != nil {
		panic(err)
	}

	// 消息入列
	c.lock.Lock()
	c.batches = append(c.batches, data)
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
				stop := c.send()
				c.state.Store(sharedStreamProcessStateIdle)
				if stop {
					break
				}
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

func (c *sharedStreamProcess) send() (stop bool) {
	for {
		c.lock.Lock()
		n := len(c.batches)
		var messages [][]byte
		if n < sharedStreamBatchLimit {
			messages = c.batches
			c.batches = nil
		} else {
			messages = c.batches[:sharedStreamBatchLimit]
			c.batches = c.batches[sharedStreamBatchLimit:]
		}
		c.lock.Unlock()
		if len(messages) == 0 {
			break
		}
		var sm *sharedMessage
		if len(messages) == 1 {
			sm = &sharedMessage{
				MessageType: &sharedMessageDeliveryMessage{
					DeliveryMessage: messages[0],
				},
			}
		} else {
			sm = &sharedMessage{
				MessageType: &sharedMessageBatchDeliveryMessage{
					BatchDeliveryMessage: &batchDeliveryMessage{Messages: messages},
				},
			}
		}

		if err := c.stream.Send(sm); err != nil {
			c.shared.detachStream(c.address)
			if c.shared.config.transportErrorHandler != nil {
				c.shared.config.transportErrorHandler(c.address, err)
			}
			c.lock.Lock()
			c.batches = append(messages, c.batches...)
			c.lock.Unlock()
			stop = true
			break
		}
	}
	return
}
