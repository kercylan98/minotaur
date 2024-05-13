package client

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/transporter"
	"io"
	netRpc "net/rpc"
)

// NewGoRPC 用于创建一个基于 GoRPC 的客户端
func NewGoRPC(network, address string, codec rpc.Codec) (rpc.Client, error) {
	cli := &goRPC{
		codec: codec,
	}
	var err error
	cli.cli, err = netRpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return cli, err
}

type goRPC struct {
	cli   *netRpc.Client
	codec rpc.Codec
}

func (c *goRPC) OnInit(network, address string) error {
	cli, err := netRpc.Dial(network, address)
	if err != nil {
		return err
	}
	c.cli = cli
	return nil
}

func (c *goRPC) Tell(route rpc.Route, data any) error {
	raw, err := c.encode(route, data)
	if err != nil {
		return err
	}

	return c.cli.Call(transporter.GoRPCRoute, raw, nil)
}

func (c *goRPC) AsyncTell(route rpc.Route, data any, callback ...func(err error)) {
	raw, err := c.encode(route, data)
	if err != nil {
		if len(callback) > 0 {
			callback[0](err)
		}
		return
	}

	call := c.cli.Go(transporter.GoRPCRoute, raw, nil, nil)
	go func(call *netRpc.Call) {
		<-call.Done
		if call.Error != nil {
			if len(callback) > 0 {
				callback[0](call.Error)
			}
		}
	}(call)
}

func (c *goRPC) Ask(route rpc.Route, data any) (rpc.Response, error) {
	raw, err := c.encode(route, data)
	if err != nil {
		return nil, err
	}

	var reply []byte
	if err = c.cli.Call(transporter.GoRPCRoute, raw, &reply); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	resp := rpc.NewResponse(func(dst any) error {
		return c.codec.DecodeData(reply, dst)
	})
	return resp, nil
}

func (c *goRPC) encode(route rpc.Route, data any) ([]byte, error) {
	raw, err := c.codec.EncodeData(data)
	if err != nil {
		return nil, err
	}

	req := rpc.NewRequest(route, raw)
	if raw, err = c.codec.EncodeRequest(req); err != nil {
		return nil, err
	}

	return raw, nil
}
