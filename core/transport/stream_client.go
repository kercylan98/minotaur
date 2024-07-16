package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
)

type StreamClientCore interface {
	// OnConnect 打开连接
	OnConnect() error

	// OnRead 读取数据包
	OnRead() (Packet, error)

	// OnWrite 写入数据包
	OnWrite(Packet) error

	// OnClose 关闭连接
	OnClose() error
}

type (
	streamClientReactPacket struct{ Packet }
)

func NewStreamClient(core StreamClientCore, config ...StreamClientConfig) *StreamClient {
	sc := &StreamClient{
		core: core,
	}
	if len(config) > 0 {
		sc.config = config[0]
	}
	return sc
}

type StreamClient struct {
	core   StreamClientCore
	err    error
	config StreamClientConfig
}

func (c *StreamClient) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		c.onLaunch(ctx)
	case vivid.FutureForwardMessage:
		ctx.Tell(ctx.Ref(), m.Error)
	case vivid.OnTerminate:
		c.onTerminate(ctx)
	case streamClientReactPacket:
		c.onReactPacket(ctx, m)
	case Packet:
		c.onWritePacket(ctx, m)
	case error:
		c.onError(ctx, m)
	}
}

func (c *StreamClient) onLaunch(ctx vivid.ActorContext) {
	if err := c.core.OnConnect(); err != nil {
		c.err = err
		ctx.ReportAbnormal(c.err)
		return
	}

	if c.config.ConnectionOpenedHandler != nil {
		c.config.ConnectionOpenedHandler(ctx)
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		defer func() {
			if err := recover(); err != nil {
				switch reason := err.(type) {
				case error:
					ctx.Tell(ctx.Ref(), err)
				default:
					ctx.Tell(ctx.Ref(), fmt.Errorf("stream client read panic: %v", reason))
				}
			}
		}()
		for {
			pkt, err := c.core.OnRead()
			if err != nil {
				return err
			}

			ctx.Tell(ctx.Ref(), streamClientReactPacket{pkt})
		}
	})
}

func (c *StreamClient) onWritePacket(ctx vivid.ActorContext, m Packet) {
	if err := c.core.OnWrite(m); err != nil {
		ctx.Tell(ctx.Ref(), err)
	}
}

func (c *StreamClient) onTerminate(ctx vivid.ActorContext) {
	if c.config.ConnectionClosedHandler != nil {
		c.config.ConnectionClosedHandler(ctx, c.err)
	}
	if c.core != nil {
		_ = c.core.OnClose()
	}
}

func (c *StreamClient) onReactPacket(ctx vivid.ActorContext, m streamClientReactPacket) {
	if c.config.ConnectionPacketHandler != nil {
		c.config.ConnectionPacketHandler(ctx, m)
	}
}

func (c *StreamClient) onError(ctx vivid.ActorContext, err error) {
	c.err = err
	ctx.ReportAbnormal(c.err)
}
