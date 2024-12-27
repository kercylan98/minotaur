package module

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"google.golang.org/protobuf/proto"
	"time"
)

type RPCModule interface {
	application.Module

	RegisterMessage(ctx vivid.ActorContext, message proto.Message)
	AffirmMessage() error
	Request(message proto.Message, timeout ...time.Duration) future.Future[vivid.Message]
	Tell(message proto.Message)
}
