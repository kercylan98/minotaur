package user

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/common/mappers"
	modelsv1 "github.com/kercylan98/minotaur/skeleton/common/models/v1"
)

func newActor(component *userComponent) *actor {
	return &actor{component: component}
}

type actor struct {
	component *userComponent
}

func (a *actor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *modelsv1.CreateUserRequest:
		a.onCreateUser(ctx, m)
	}
}

func (a *actor) onCreateUser(ctx vivid.ActorContext, m *modelsv1.CreateUserRequest) {
	switch request := m.UserType.(type) {
	case *modelsv1.CreateUserRequest_Guest:
		a.onCreateGuestUser(ctx, request)
	}
}

func (a *actor) onCreateGuestUser(ctx vivid.ActorContext, request *modelsv1.CreateUserRequest_Guest) {
	if user, err := mappers.GetUserMapper().CreateGuestUser(a.component.database.GetMDB()); err != nil {
		ctx.Reply(err)
	} else {
		ctx.Reply(user)
	}
}
