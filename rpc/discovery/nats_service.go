package discovery

import (
	"context"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/registry"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/nats-io/nats.go"
)

type NatsDiscoveryService struct {
	service rpc.Service
	conn    *nats.Conn
	sub     string
}

func (n *NatsDiscoveryService) init(service rpc.Service, conn *nats.Conn, sub string) *NatsDiscoveryService {
	n.service = service
	n.conn = conn
	n.sub = sub
	return n
}

func (n *NatsDiscoveryService) UnaryCall(route ...rpc.Route) rpc.UnaryCaller {
	return func(ctx context.Context, params any) (rpc.Reader, error) {
		resp, err := n.conn.RequestWithContext(ctx, n.sub, new(registry.NatsCaller).Marshal(route, toolkit.MarshalJSON(params)))
		if err != nil {
			return nil, err
		}
		return func(dst any) {
			toolkit.UnmarshalJSON(resp.Data, dst)
		}, nil
	}
}

func (n *NatsDiscoveryService) UnaryNotifyCall(route ...rpc.Route) rpc.UnaryNotifyCaller {
	return func(ctx context.Context, params any) error {
		_, err := n.conn.RequestWithContext(ctx, n.sub, new(registry.NatsCaller).Marshal(route, toolkit.MarshalJSON(params)))
		return err
	}
}

func (n *NatsDiscoveryService) AsyncUnaryCall(route ...rpc.Route) rpc.AsyncUnaryCaller {
	return func(ctx context.Context, params any, callback func(reader rpc.Reader, err error)) {
		go func() {
			resp, err := n.conn.RequestWithContext(ctx, n.sub, new(registry.NatsCaller).Marshal(route, toolkit.MarshalJSON(params)))
			if err != nil {
				callback(nil, err)
			}

			callback(func(dst any) {
				toolkit.UnmarshalJSON(resp.Data, dst)
			}, nil)
		}()
	}
}

func (n *NatsDiscoveryService) AsyncNotifyCall(route ...rpc.Route) rpc.AsyncNotifyCaller {
	return func(params any) error {
		return n.conn.Publish(n.sub, new(registry.NatsCaller).Marshal(route, toolkit.MarshalJSON(params)))
	}
}

func (n *NatsDiscoveryService) GetServiceInfo() rpc.Service {
	return n.service
}
