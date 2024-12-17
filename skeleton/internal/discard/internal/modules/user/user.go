package user

import (
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/module"
	"github.com/kercylan98/minotaur/skeleton/internal/modules"
)

func NewUserModule() module.Module[*application.Context] {
	return &userModule{}
}

type userModule struct {
	app      *application.Context
	services modules.UserModule
}

func (u *userModule) OnInitialize(app *application.Context) error {
	u.app = app
	u.services = newUserServices(u)
	return nil
}
