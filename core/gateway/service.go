package gateway

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisor"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/router"
	"google.golang.org/protobuf/proto"
	"time"
)

func NewService(gatewayAddr string, info *ServiceInfo) *Service {
	return &Service{
		gatewayAddr: gatewayAddr,
		info:        info,
	}
}

type Service struct {
	support     *vivid.ModuleSupport
	gatewayAddr string
	info        *ServiceInfo
	router      *router.Multistage[func(ctx vivid.ActorContext, response *ResponseWrapper, packet transport.Packet)]
}

func (s *Service) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	s.support = support
	s.router = router.NewMultistage[func(ctx vivid.ActorContext, response *ResponseWrapper, packet transport.Packet)]()

	s.router.Route(MessageId_MI_ServiceRegister, s.onServiceRegisterResponse)
	s.router.Route(MessageId_MI_Heartbeat, s.onHeartbeatResponse)

	s.support.System().ActorOf(func() vivid.Actor {
		return transport.NewStreamClient(&transport.StreamClientTCPCore{Addr: s.gatewayAddr}, transport.StreamClientConfig{
			ConnectionOpenedHandler: s.onConnectionOpened,
			ConnectionPacketHandler: s.onConnectionPacket,
			ConnectionClosedHandler: s.onConnectionClosed,
		})
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix("gateway")
		options.WithSupervisorStrategy(supervisor.OneForOne(-1, time.Millisecond*100, time.Second, func(reason, message supervisor.Message) supervisor.Directive {
			return supervisor.DirectiveRestart
		}))
	})
}

func (s *Service) onConnectionOpened(ctx vivid.ActorContext) {
	data, err := wrapRequest(MessageId_MI_ServiceRegister, &ServiceRegisterRequest{
		ServiceInfo: s.info,
	})
	if err != nil {
		ctx.Tell(ctx.Ref(), err)
		return
	}

	ctx.Tell(ctx.Ref(), transport.NewPacket(data))
}

func (s *Service) onConnectionPacket(ctx vivid.ActorContext, packet transport.Packet) {
	var w ResponseWrapper
	// err : 数据包能够正常解析但是会包含一个额外的错误，暂时不清楚具体原因
	_ = proto.Unmarshal(packet.GetBytes(), &w)
	//if err != nil {
	//	ctx.Tell(ctx.Ref(), fmt.Errorf("unmarshal packet error: %w", err))
	//	return
	//}

	handler := s.router.Match(w.Id)
	if handler == nil {
		ctx.Tell(ctx.Ref(), fmt.Errorf("no handler for packet message id: %s", w.Id.String()))
		return
	}

	handler(ctx, &w, packet)
}

func (s *Service) onConnectionClosed(ctx vivid.ActorContext, err error) {
	if err != nil {
		s.support.Logger().Error("Gateway", log.String("type", "service disconnected"), log.Any("service", s.info), log.Err(err))
	} else {
		s.support.Logger().Error("Gateway", log.String("type", "service disconnected"), log.Any("service", s.info))
	}
}

func (s *Service) onServiceRegisterResponse(ctx vivid.ActorContext, response *ResponseWrapper, packet transport.Packet) {
	latencyOverall := time.Duration(response.ServerSendAt-response.ClientSendAt) * time.Millisecond
	s.support.Logger().Info("Gateway", log.String("type", "service registered"), log.Duration("delay", latencyOverall), log.Any("service", s.info))

	// 开启心跳
	ctx.RepeatedTask("heartbeat", chrono.SchedulerInstantly, time.Second*5, chrono.SchedulerForever, s.onHeartbeat)
}

func (s *Service) onHeartbeat(ctx vivid.ActorContext) {
	data, err := wrapRequest(MessageId_MI_Heartbeat, nil)
	if err != nil {
		ctx.Tell(ctx.Ref(), err)
		return
	}

	ctx.Tell(ctx.Ref(), transport.NewPacket(data))
}

func (s *Service) onHeartbeatResponse(ctx vivid.ActorContext, response *ResponseWrapper, packet transport.Packet) {
	// nothing to do
}
