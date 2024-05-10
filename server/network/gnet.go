package network

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/panjf2000/gnet/v2"
	"time"
)

var (
	schemaWebSocket = "ws"
	schemaTcp       = "tcp"
	schemaTcp4      = "tcp4"
	schemaTcp6      = "tcp6"
	schemaUdp       = "udp"
	schemaUdp4      = "udp4"
	schemaUdp6      = "udp6"
	schemaUnix      = "unix"
)

type gNetHandler interface {
	OnInit(core *gNetCore)
	gnet.EventHandler

	GetEngine() *gnet.Engine
}

func newGNetCore(handler gNetHandler, schema, addr string, pattern ...string) server.Network {
	ws := &gNetCore{
		handler: handler,
		addr:    addr,
		schema:  schema,
		pattern: collection.FindFirstOrDefaultInSlice(pattern, "/"),
	}
	return ws
}

type gNetCore struct {
	ctx        context.Context
	controller server.Controller
	handler    gNetHandler
	addr       string
	schema     string
	pattern    string
}

func (w *gNetCore) OnSetup(ctx context.Context, controller server.Controller) (err error) {
	w.ctx = ctx
	w.controller = controller
	w.handler.OnInit(w)
	return
}

func (w *gNetCore) OnRun() (err error) {
	var addr string
	switch w.schema {
	case schemaTcp, schemaWebSocket:
		addr = fmt.Sprintf("tcp://%s", w.addr)
	case schemaTcp4:
		addr = fmt.Sprintf("tcp4://%s", w.addr)
	case schemaTcp6:
		addr = fmt.Sprintf("tcp6://%s", w.addr)
	case schemaUdp:
		addr = fmt.Sprintf("udp://%s", w.addr)
	case schemaUdp4:
		addr = fmt.Sprintf("udp4://%s", w.addr)
	case schemaUdp6:
		addr = fmt.Sprintf("udp6://%s", w.addr)
	case schemaUnix:
		addr = fmt.Sprintf("unix://%s", w.addr)
	default:
		return fmt.Errorf("unsupported schema: %s", w.schema)
	}

	err = gnet.Run(w.handler, addr, gnet.WithLogger(&gNetLogger{w.controller}))
	return
}

func (w *gNetCore) OnShutdown() error {
	if w.handler.GetEngine() != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return w.handler.GetEngine().Stop(ctx)
	}
	return nil
}

func (w *gNetCore) Schema() string {
	return w.schema
}

func (w *gNetCore) Address() string {
	if w.pattern == "/" {
		return w.addr
	}
	return fmt.Sprintf("%s:%s", w.addr, w.pattern)
}
