package discovery

import (
	"context"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/internal/utils"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
)

func NewNats(conn *nats.Conn, js nats.JetStreamContext, bucketName, bucketDesc, keyPrefix string, ttl time.Duration) (n *Nats, err error) {
	n = &Nats{
		conn:         conn,
		js:           js,
		services:     make(map[string]map[string]rpc.Service),
		registerCh:   make(chan rpc.CallableService, 128),
		unregisterCh: make(chan rpc.CallableService, 128),
		keyPrefix:    keyPrefix,
		watchKey:     keyPrefix + ".*",
	}
	n.ctx, n.cancel = context.WithCancel(context.Background())
	n.kv, err = utils.InitNatsBucket(js, bucketName, bucketDesc, ttl)
	starter, err := n.watch()
	if err != nil {
		return nil, err
	}
	go starter()
	return
}

type Nats struct {
	conn         *nats.Conn
	js           nats.JetStreamContext
	ctx          context.Context
	cancel       context.CancelFunc
	kv           nats.KeyValue
	servicesRW   sync.RWMutex
	services     map[string]map[string]rpc.Service // serviceName -> instanceId -> service
	registerCh   chan rpc.CallableService
	unregisterCh chan rpc.CallableService
	keyPrefix    string
	watchKey     string
}

func (n *Nats) WatchRegister() <-chan rpc.CallableService {
	return n.registerCh
}

func (n *Nats) WatchUnregister() <-chan rpc.CallableService {
	return n.unregisterCh
}

func (n *Nats) Close() error {
	n.cancel()
	return nil
}

func (n *Nats) watch() (func(), error) {
	watcher, err := n.kv.Watch(n.watchKey)
	if err != nil {
		return nil, err
	}

	return func() {
		defer func() {
			_ = watcher.Stop()
			close(n.registerCh)
			close(n.unregisterCh)
		}()

		for {
			select {
			case <-n.ctx.Done():
				return
			case entry := <-watcher.Updates():
				if entry == nil {
					continue
				}

				var service rpc.Service
				toolkit.UnmarshalJSON(entry.Value(), &service)
				var cs = new(NatsDiscoveryService).init(service, n.conn, n.keyPrefix+".call."+service.InstanceId)
				switch entry.Operation() {
				case nats.KeyValuePut:
					n.registerCh <- cs
				case nats.KeyValueDelete:
					n.unregisterCh <- cs
				default:
				}
			}
		}
	}, nil
}
