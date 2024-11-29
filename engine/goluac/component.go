package goluac

import (
	"errors"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func Bind(ctx vivid.ActorContext, code string) {
	if ctx.HasValue(goluacKey) {
		panic(errors.New("the goluac component has been bound"))
	}
	ctx.Tell(ctx.Ref(), newGoluac(ctx, code))
}

func NewComponent() Component {
	return &component{}
}

var goluacKey Component = (*component)(nil)

type Component interface {
	vivid.Component
	vivid.ActorReceiveMessageCaptureComponent
}

type component struct {
}

func (c *component) OnInitialize(actorSystem *vivid.ActorSystem) error {
	return nil
}

func (c *component) OnActorReceiveMessageCapture(ctx vivid.ActorContext) (abort bool) {
	switch m := ctx.Message().(type) {
	case *goluac:
		ctx.SetValue(goluacKey, m)
	case *vivid.OnTerminated:
		c, exist := ctx.GetValue(goluacKey).(*goluac)
		if exist {
			c.lib.Close()
		}
	}

	return
}
