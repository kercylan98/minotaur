package rpccore

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/codec"
	"github.com/nats-io/nats.go"
	"time"
)

const (
	bucketName = "minotaur-rpc-services"
	bucketDesc = "Minotaur RPC Services Management Bucket, used for storing service information."
)

// NewNats 基于 Nats 的 RPC 核心，该核心被要求 Nats 必须支持 JetStream
func NewNats(addr string, options ...nats.Option) (rpc.Core, error) {
	var core = new(Nats)
	var err error
	core.services = make(map[string]map[string]natsService)
	core.targets = make(map[string]natsService)
	core.ctx, core.cancel = context.WithCancel(context.Background())

	// 连接 Nats
	core.conn, err = nats.Connect(addr, options...)
	if err != nil {
		panic(err)
	}
	core.js, err = core.conn.JetStream()
	if err != nil {
		panic(err)
	}

	return core, nil
}

type Nats struct {
	conn         *nats.Conn
	js           nats.JetStreamContext
	kv           nats.KeyValue
	ctx          context.Context
	cancel       context.CancelFunc
	info         rpc.ServiceInfo
	matcher      rpc.RouteMatcher
	subscription *nats.Subscription
	watcher      nats.KeyWatcher
	base64       codec.Base64

	curr     natsService
	services map[string]map[string]natsService
	targets  map[string]natsService
}

func (n *Nats) OnInit(info rpc.ServiceInfo, matcher rpc.RouteMatcher, routes [][]rpc.Route) error {
	n.info = info
	n.matcher = matcher
	n.curr = natsService{
		ServerInfo: info,
		Routes:     routes,
	}
	var err []error
	err = append(err, n.initBucket())
	err = append(err, n.initSubscription())
	err = append(err, n.initWatcher())

	n.update(n.curr)
	go n.loop()
	return nil
}

func (n *Nats) OnCall(routes ...rpc.Route) func(request any) error {
	var packet = natsPacket{
		IsRequest: true,
		Request: natsPacketRequest{
			Routes: routes,
		},
	}

	return func(request any) error {
		packet.Request.Data = toolkit.MarshalJSON(request)

		jsonData := toolkit.MarshalJSON(routes)
		b64, err := n.base64.Encode(jsonData)
		if err != nil {
			return err
		}
		target := n.targets[string(b64)]
		return n.conn.Publish("services.rpc."+target.ServerInfo.UniqueId, toolkit.MarshalJSON(packet))
	}
}

func (n *Nats) Close() {
	n.cancel()
}

func (n *Nats) refresh() error {
	_, err := n.kv.Put("services."+n.info.UniqueId, toolkit.MarshalJSON(n.info))
	return err
}

func (n *Nats) update(info natsService) {
	nodes, exist := n.services[info.ServerInfo.Name]
	if !exist {
		nodes = make(map[string]natsService)
		n.services[info.ServerInfo.Name] = nodes
	}

	nodes[info.ServerInfo.UniqueId] = info

	// 为路由生成对应服务的映射
	for _, parts := range info.Routes {
		jsonData := toolkit.MarshalJSON(parts)
		b64, err := n.base64.Encode(jsonData)
		if err != nil {
			continue
		}

		n.targets[string(b64)] = info
	}
}

func (n *Nats) loop() {
	heartbeat := time.NewTicker(time.Second * 3)
	defer func() {
		heartbeat.Stop()
		_ = n.kv.Delete("services." + n.info.UniqueId)
		_ = n.watcher.Stop()
		_ = n.subscription.Unsubscribe()
	}()

	for {
		select {
		case <-n.ctx.Done():
			return
		case <-heartbeat.C:
			if err := n.refresh(); err != nil {
				// TODO: log
			}
		case entry := <-n.watcher.Updates():
			if entry != nil {
				switch entry.Operation() {
				case nats.KeyValuePut:
					var info natsService
					toolkit.UnmarshalJSON(entry.Value(), &info)
					n.update(info)
				case nats.KeyValueDelete:
					var info natsService
					toolkit.UnmarshalJSON(entry.Value(), &info)
					delete(n.services[info.ServerInfo.Name], info.ServerInfo.UniqueId)
					if len(n.services[info.ServerInfo.Name]) == 0 {
						delete(n.services, info.ServerInfo.Name)
					}
				default:
				}

			}

		}
	}
}

func (n *Nats) onCall(msg *nats.Msg) {
	var packet natsPacket
	toolkit.UnmarshalJSON(msg.Data, &packet)
	if packet.IsRequest {
		handler := n.matcher.Match(packet.Request.Routes...)
		if handler == nil {
			return
		}

		handler(func(dst any) {
			toolkit.UnmarshalJSON(packet.Request.Data, dst)
		})
	}

}

func (n *Nats) initBucket() error {
	// 创建 KeyValue 存储桶
	kv, err := n.js.KeyValue(bucketName)
	if err != nil {
		if !errors.Is(err, nats.ErrBucketNotFound) {
			return err
		}
	}

	if kv == nil {
		kv, err = n.js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:      bucketName,
			Description: bucketDesc,
			History:     9,
			TTL:         time.Second * 10,
			Storage:     nats.FileStorage,
		})
		if err != nil {
			return err
		}
	}
	n.kv = kv
	return nil
}

func (n *Nats) initSubscription() error {
	// 启动订阅
	var err error
	if n.subscription, err = n.conn.QueueSubscribe("services.rpc."+n.info.UniqueId, n.info.UniqueId, n.onCall); err != nil {
		return err
	}
	return err
}

func (n *Nats) initWatcher() error {
	// 注册监听
	w, err := n.kv.Watch("services.*")
	if err != nil {
		return err
	}
	if err = n.refresh(); err != nil {
		_ = w.Stop()
		return err
	}
	n.watcher = w
	return nil
}
