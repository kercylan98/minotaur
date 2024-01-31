package attack

import (
	"fmt"
	"github.com/kercylan98/minotaur/modular/example/internal/service/expose"
)

type Service struct {
	Login expose.Login
	name  string
}

func (a *Service) OnInit() {
	expose.AttackExpose = a
	a.name = "attack"
}

func (a *Service) OnPreload() {
	a.Login = expose.LoginExpose
}

func (a *Service) OnMount() {
	fmt.Println("attack service mounted, call", a.Login.Name(), "service")
}

func (a *Service) Name() string {
	return a.name
}
