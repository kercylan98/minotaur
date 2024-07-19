package prc

import "sync/atomic"

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
	closed  atomic.Bool
	address PhysicalAddress
}

func (c *sharedStreamProcess) Initialize(rc *ResourceController, id *ProcessId) {
	// 该进程不注册，不会触发
}

func (c *sharedStreamProcess) DeliveryUserMessage(receiver, sender, forward *ProcessRef, message Message) {
	c.packMessage(receiver, sender, forward, message, false)
}

func (c *sharedStreamProcess) DeliverySystemMessage(receiver, sender, forward *ProcessRef, message Message) {
	c.packMessage(receiver, sender, forward, message, true)
}

func (c *sharedStreamProcess) packMessage(receiver, sender, forward *ProcessRef, message Message, system bool) {
	name, data, err := c.shared.codec.Encode(message)
	if err != nil {
		panic(err)
	}
	dm := &DeliveryMessage{
		MessageType: name,
		MessageData: data,
		System:      system,
	}
	if sender != nil {
		dm.Sender = sender.id
	}
	if receiver != nil {
		dm.Receiver = receiver.id
	}
	if err = c.stream.Send(&SharedMessage{
		MessageType: &SharedMessage_DeliveryMessage{
			DeliveryMessage: dm,
		},
	}); err != nil {
		panic(err)
	}
}

func (c *sharedStreamProcess) IsTerminated() bool {
	return c.closed.Load()
}

func (c *sharedStreamProcess) Terminate(source *ProcessRef) {
	// 该进程不注册，不会由资源控制器触发
}
