package registry

import (
	"context"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/internal/utils"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/nats-io/nats.go"
	"time"
)

func NewNats(conn *nats.Conn, js nats.JetStreamContext, bucketName, bucketDesc, keyPrefix string, ttl, keepAliveInterval time.Duration) (n *Nats, err error) {
	n = &Nats{
		conn:      conn,
		js:        js,
		key:       keyPrefix,
		keepAlive: keepAliveInterval,
		caller:    make(chan rpc.Caller, 1024),
	}
	n.ctx, n.cancel = context.WithCancel(context.Background())
	n.kv, err = utils.InitNatsBucket(js, bucketName, bucketDesc, ttl)
	return
}

// Nats 基于 Nats 实现的注册器
type Nats struct {
	conn         *nats.Conn
	ctx          context.Context
	cancel       context.CancelFunc
	js           nats.JetStreamContext
	kv           nats.KeyValue
	caller       chan rpc.Caller
	service      rpc.Service
	serviceCache []byte
	keepAlive    time.Duration
	key          string
	subject      string
}

func (n *Nats) OnRegister(service rpc.Service) error {
	n.service = service
	n.key = n.key + "." + service.InstanceId
	n.subject = n.key + ".call." + service.InstanceId
	n.serviceCache = toolkit.MarshalJSON(service)

	// 创建订阅
	sub, err := n.conn.QueueSubscribe(n.subject, n.service.Name, n.onSubscribe)
	if err != nil {
		return err
	}
	sub.SetClosedHandler(func(subject string) {
		close(n.caller)
	})

	// 订阅成功后，开始维持心跳
	if err = n.onHeartbeat(); err != nil {
		return err
	}
	n.run(sub)
	return nil
}

func (n *Nats) OnUnregister() error {
	n.cancel()
	return n.kv.Delete(n.key)
}

func (n *Nats) WatchCall() <-chan rpc.Caller {
	return n.caller
}

func (n *Nats) onHeartbeat() error {
	_, err := n.kv.Put(n.key, n.serviceCache)
	return err
}

func (n *Nats) onSubscribe(data *nats.Msg) {
	var caller = new(NatsCaller).Unmarshal(data)
	if err := toolkit.UnmarshalJSONE(data.Data, &caller); err != nil {
		return
	}
	n.caller <- caller
}

func (n *Nats) run(sub *nats.Subscription) {
	ticker := time.NewTicker(n.keepAlive)

	defer func() {
		_ = sub.Unsubscribe()
		ticker.Stop()
	}()

	for {
		select {
		case <-n.ctx.Done():
			return
		case <-ticker.C:
			if err := n.onHeartbeat(); err != nil {
				return
			}
		default:
		}
	}
}
