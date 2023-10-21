package cross

import (
	"encoding/json"
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/nats-io/nats.go"
	"time"
)

const (
	nasMark = "Cross.Nats"
)

// NewNats 创建一个基于 Nats 实现的跨服消息功能组件
func NewNats(url string, options ...NatsOption) *Nats {
	n := &Nats{
		url:     url,
		subject: "MINOTAUR_CROSS",
		messagePool: concurrent.NewPool[*Message](1024*100, func() *Message {
			return new(Message)
		}, func(data *Message) {
			data.ServerId = ""
			data.Packet = nil
		}),
	}
	for _, option := range options {
		option(n)
	}
	return n
}

// Nats 基于 Nats 实现的跨服消息功能组件
type Nats struct {
	conn        *nats.Conn
	url         string
	subject     string
	options     []nats.Option
	messagePool *concurrent.Pool[*Message]
}

func (slf *Nats) Init(server *server.Server, packetHandle func(serverId string, packet []byte)) (err error) {
	if slf.conn == nil {
		if len(slf.options) == 0 {
			slf.options = append(slf.options,
				nats.ReconnectWait(time.Second*5),
				nats.MaxReconnects(-1),
				nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
					log.Error(nasMark, log.String("info", "disconnect"), log.Err(err))
				}),
				nats.ReconnectHandler(func(conn *nats.Conn) {
					log.Info(nasMark, log.String("info", "reconnect"))
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
			log.Error(nasMark, log.Err(err))
			return
		}
		packetHandle(message.ServerId, message.Packet)
	})
	return err
}

func (slf *Nats) PushMessage(serverId string, packet []byte) error {
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
