package moduleimpl

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

var _ module.ActorSystemModule = (*ActorSystemModule)(nil)

func NewActorSystemModule() *ActorSystemModule {
	return &ActorSystemModule{}
}

type ActorSystemModule struct {
	actorSystem *vivid.ActorSystem
}

func (a *ActorSystemModule) OnInitialize(ctx *application.Context) (err error) {
	a.actorSystem = vivid.NewActorSystem()

	return
}

func (a *ActorSystemModule) OnDependencyInitialize(ctx *application.Context) (err error) {
	return
}

func (a *ActorSystemModule) OnDependencySetup() (err error) {
	return
}

func (a *ActorSystemModule) OnSetup() (err error) {
	return
}

func (a *ActorSystemModule) ActorSystem() *vivid.ActorSystem {
	return a.actorSystem
}
