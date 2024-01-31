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
	expose.LoginExpose = l
	l.name = "login"
}

func (l *Service) OnPreload() {
	l.Attack = expose.AttackExpose
}

func (l *Service) OnMount() {
	fmt.Println("attack service mounted, call", l.Attack.Name(), "service")
}

func (l *Service) Name() string {
	return l.name
}
