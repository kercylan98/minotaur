package cross

import (
	"encoding/json"
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

func NewNats(url string, options ...NatsOption) *Nats {
	n := &Nats{
		url:     url,
		subject: "MINOTAUR_CROSS",
		messagePool: synchronization.NewPool[*Message](1024*100, func() *Message {
			return new(Message)
		}, func(data *Message) {
			data.ServerId = 0
			data.Packet = nil
		}),
	}
	for _, option := range options {
		option(n)
	}
	return n
}

type Nats struct {
	conn        *nats.Conn
	url         string
	subject     string
	options     []nats.Option
	messagePool *synchronization.Pool[*Message]
}

func (slf *Nats) Init(server *server.Server, packetHandle func(serverId int64, packet []byte)) (err error) {
	if slf.conn == nil {
		if len(slf.options) == 0 {
			slf.options = append(slf.options,
				nats.ReconnectWait(time.Second*5),
				nats.MaxReconnects(-1),
				nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
					log.Error("Cross.Nats", zap.String("info", "disconnect"), zap.Error(err))
				}),
				nats.ReconnectHandler(func(conn *nats.Conn) {
					log.Info("Cross.Nats", zap.String("info", "reconnect"))
				}),
			)
		}
		slf.conn, err = nats.Connect(slf.url, slf.options...)
		if err != nil {
			return err
		}
	}
	_, err = slf.conn.Subscribe(fmt.Sprintf("%s_%d", slf.subject, server.GetID()), func(msg *nats.Msg) {
		message := slf.messagePool.Get()
		defer slf.messagePool.Release(message)
		if err := json.Unmarshal(msg.Data, &message); err != nil {
			log.Error("Cross.Nats", zap.Error(err))
			return
		}
		packetHandle(message.ServerId, message.Packet)
	})
	return err
}

func (slf *Nats) PushMessage(serverId int64, packet []byte) error {
	message := slf.messagePool.Get()
	defer slf.messagePool.Release(message)
	message.ServerId = serverId
	message.Packet = packet
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return slf.conn.Publish(fmt.Sprintf("%s_%d", slf.subject, serverId), data)
}

func (slf *Nats) Release() {
	slf.conn.Close()
}
