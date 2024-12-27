package controller

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/api/rpcmessage"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

var _ application.Controller = (*OAuthController)(nil)

func NewOAuthController() *OAuthController {
	return &OAuthController{}
}

type OAuthController struct {
	modules struct {
		actorSystem module.ActorSystemModule
		rpc         module.RPCModule
	}
}

func (o *OAuthController) OnInitialize(ctx *application.Context, loader *application.ServiceLoader) (err error) {
	o.modules.actorSystem = application.LoadModule[module.ActorSystemModule](ctx)
	o.modules.rpc = application.LoadModule[module.RPCModule](ctx)

	o.modules.actorSystem.ActorSystem().ActorOfF(func() vivid.Actor {
		return o.onSpawnActor()
	})
	return
}

func (o *OAuthController) onSpawnActor() vivid.FunctionalActor {
	// state
	return func(ctx vivid.ActorContext) {
		switch ctx.Message().(type) {
		case *vivid.OnLaunch:
			o.modules.rpc.RegisterMessage(ctx, new(rpcmessage.RPCOAuthLoginRequest))
			if err := o.modules.rpc.AffirmMessage(); err != nil {
				panic(err)
			}
		case *rpcmessage.RPCOAuthLoginRequest:
			fmt.Println("RPCOAuthLoginRequest")
			ctx.Reply(&rpcmessage.RPCOAuthLoginResponse{
				Token: "token",
			})
		}
	}
}
