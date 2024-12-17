package user

import "github.com/kercylan98/minotaur/skeleton/internal/modules"

func newUserServices(module *userModule) modules.UserModule {
	return &userServices{module: module}
}

type userServices struct {
	module *userModule
}
