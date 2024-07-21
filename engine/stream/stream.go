package stream

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
)

// NewStream 创建一个新的 Stream 实例，它可以处理的消息类型如下：
//
// 处理消息：
//   - Packet: 写入数据到流中
//   - error: 报告异常并触发监管策略（如果有）
//
// 特殊行为表现：
//   - Packet: 收到 Packet 消息即表示对端发送的数据，需要对数据进行处理
//
// 该实例的行为表现由 performance 参数指定
func NewStream(conn Conn, performance behavior.Performance[vivid.ActorContext]) *Stream {
	return &Stream{
		conn:        conn,
		performance: performance,
	}
}

// Stream 是基于 Actor 模型实现的通用的流式数据传输结构
type Stream struct {
	conn        Conn
	performance behavior.Performance[vivid.ActorContext]
	behavior    behavior.Behavior[vivid.ActorContext]
}

func (s *Stream) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		s.onLaunch(ctx)
	case *vivid.OnTerminate:
		s.onTerminate(ctx)
	case Packet:
		s.onPacket(ctx, m)
	case readPacket:
		s.onReadPacket(ctx, (Packet)(m))
	case error:
		s.onError(ctx, m)
	default:
		s.behavior.Perform(ctx)
	}
}

func (s *Stream) onLaunch(ctx vivid.ActorContext) {
	s.behavior = behavior.New[vivid.ActorContext]()
	s.behavior.Become(s.performance)
	s.behavior.Perform(ctx)

	f := ctx.Future()
	f.AwaitForward(ctx.Ref(), func() prc.Message {
		for {
			pkt, err := s.conn.Read()
			if err != nil {
				return err
			}

			ctx.Tell(ctx.Ref(), (readPacket)(pkt))
		}
	})
}

func (s *Stream) onError(ctx vivid.ActorContext, err error) {
	if err == nil {
		return
	}
	ctx.ReportAbnormal(err)
}

func (s *Stream) onTerminate(ctx vivid.ActorContext) {
	s.behavior.Perform(ctx)
	_ = s.conn.Close()
}

func (s *Stream) onPacket(ctx vivid.ActorContext, packet Packet) {
	if err := s.conn.Write(packet); err != nil {
		ctx.Tell(ctx.Ref(), err)
	}
}

func (s *Stream) onReadPacket(ctx vivid.ActorContext, m Packet) {
	ctx.CastMessage(m)
	s.behavior.Perform(ctx)
}
