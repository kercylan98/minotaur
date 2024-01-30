package login

import (
	"fmt"
	"github.com/kercylan98/minotaur/modular/example/internal/service/expose"
)

type Service struct {
	Attack expose.Attack
	name   string
}

func (l *Service) OnInit() {
	l.name = "login"
}

func (l *Service) OnPreload() {
	fmt.Println(l.name, "call", l.Attack.Name())
}

func (l *Service) OnMount() {

}

func (l *Service) Name() string {
	return l.name
}
