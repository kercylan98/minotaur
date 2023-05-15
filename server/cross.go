package server

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"go.uber.org/zap"
)

// cross 跨服功能 TODO: 跨服逻辑存在问题
type cross struct {
	server         *Server
	messageChannel chan *crossMessage
	messagePool    *synchronization.Pool[*crossMessage]
	queues         map[CrossQueueName]CrossQueue
}

func (slf *cross) Run(server *Server, queues ...CrossQueue) error {
	slf.server = server
	slf.queues = map[CrossQueueName]CrossQueue{}
	slf.messagePool = synchronization.NewPool[*crossMessage](100,
		func() *crossMessage {
			return &crossMessage{}
		}, func(data *crossMessage) {
			data.toServerId = 0
			data.ServerId = 0
			data.Queue = ""
			data.Packet = nil
		},
	)
	slf.messageChannel = make(chan *crossMessage, 4096*100)
	for i := 0; i < len(slf.queues); i++ {
		queue := queues[i]
		if _, exist := slf.queues[queue.GetName()]; exist {
			return ErrCrossDuplicateQueue
		}
		if err := queue.Init(); err != nil {
			return err
		}
		slf.queues[queue.GetName()] = queue
		queue.Subscribe(slf.server.GetID(), func(bytes []byte) {
			message := slf.messagePool.Get()
			if err := json.Unmarshal(bytes, message); err != nil {
				log.Error("Cross", zap.String("Queue.Receive", string(queue.GetName())), zap.String("Packet", string(bytes)), zap.Error(err))
				return
			}
			slf.server.PushMessage(MessageTypeCross, message.ServerId, message.Queue, message.Packet)
			slf.messagePool.Release(message)
		})
	}
	go func() {
		for message := range slf.messageChannel {
			queue := slf.queues[message.Queue]
			data, err := json.Marshal(message)
			if err != nil {
				log.Error("Cross", zap.String("Queue.Push", string(queue.GetName())), zap.String("Packet", string(message.Packet)), zap.Error(err))
			} else if err = queue.Publish(message.toServerId, data); err != nil {
				log.Error("Cross", zap.String("Queue.Push", string(queue.GetName())), zap.Error(err))
			}
			slf.messagePool.Release(message)
		}
	}()
	return nil
}

func (slf *cross) PushCrossMessage(queue CrossQueueName, serverId int64, packet []byte) {
	message := slf.messagePool.Get()
	message.toServerId = serverId
	message.ServerId = slf.server.GetID()
	message.Queue = queue
	message.Packet = packet
	slf.messageChannel <- message
}

func (slf *cross) shutdownCross() {
	close(slf.messageChannel)
	slf.messagePool.Close()
	slf.messagePool = nil
}
