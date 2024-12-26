package controller

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

type OAuthController struct {
	modules struct {
		actorSystem module.ActorSystemModule
	}
}

func (o *OAuthController) OnInitialize(ctx *application.Context, loader *application.ServiceLoader) (err error) {
	o.modules.actorSystem = application.LoadModule[module.ActorSystemModule](ctx)
	o.modules.actorSystem.ActorSystem().ActorOfF(func() vivid.Actor {
		return o.OnReceive()
	})
	return
}

func (o *OAuthController) OnReceive() vivid.FunctionalActor {
	// state
	return func(ctx vivid.ActorContext) {
		switch ctx.Message().(type) {
		case *vivid.Message:

		}
	}
}
