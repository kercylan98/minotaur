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
	a.name = "attack"
}

func (a *Service) OnPreload() {
	fmt.Println(a.name, "call", a.Login.Name())
}

func (a *Service) OnMount() {

}

func (a *Service) Name() string {
	return a.name
}
